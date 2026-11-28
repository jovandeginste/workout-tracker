# LDAP authentication (#205) — implementation design

Status: implemented on branch `feat/ldap-auth`.

## Context

#205 asks for central authentication against an LDAP directory. The issue
comment says there are no plans yet, but a PR is welcome, and that `REMOTE_USER`
behind an authenticating reverse proxy is the intended alternative.

Decision: direct LDAP in the app. It keeps the single-binary deployment and does
not trust a proxy header. `REMOTE_USER` is not part of this work.

This supersedes the current `ISSUES.md` triage note that points at
`REMOTE_USER`.

Goals:

- Users sign in with their directory credentials.
- Local accounts keep working, so there is a break-glass admin.
- No directory password is ever stored.
- Existing installations behave exactly as before until `ldap_enabled` is set.
- Works with OpenLDAP and Active Directory through configuration, not code.

Out of scope for the first PR:

- `REMOTE_USER` / header authentication.
- OAuth, SAML, Kerberos.
- Background synchronisation of the whole directory.
- Nested AD group expansion.
- Password-change flows.
- Rate limiting on the login endpoint. There is none today, and this work does
  not add one.

## 1. User model: an explicit authentication source

Add an `AuthSource` to `UserData`:

```go
const (
	AuthSourceLocal = "local"
	AuthSourceLDAP  = "ldap"
)

type UserData struct {
	// ...
	AuthSource string `gorm:"type:varchar(16);not null;default:local" json:"authSource"`
}
```

No `form` tag. Registration and admin handlers set it explicitly; never bind it
straight from a request. The registration handler must still overwrite
`u.AuthSource = database.AuthSourceLocal` after `c.Bind(u)`, because Echo's
binder can match an `AuthSource` form field by the struct field name even
without a form tag.

Rules:

| Stored user | LDAP enabled         | Behaviour                                              |
| ----------- | -------------------- | ------------------------------------------------------ |
| `local`     | no                   | local password, unchanged                              |
| `local`     | yes                  | local password only; LDAP never takes over the account |
| `ldap`      | yes                  | LDAP bind only; local password is ignored              |
| missing     | yes, auto-create on  | create an `ldap` user after a successful bind          |
| missing     | yes, auto-create off | login fails                                            |

The important one is rule two. If a local admin `jo` exists and the directory
also has a `jo`, a successful LDAP bind must not log in the local admin. That
would be an account takeover. Linking an existing user to LDAP is an explicit
admin action, never automatic.

Auto-create is off by default. An LDAP user must therefore already exist in the
app, created by an admin or converted from a local account. That keeps the
directory from silently adding app users.

Existing invariants need a small change. Treat an empty source as local, so a
row that predates the migration keeps the old behaviour:

```go
func (u *User) authSource() string {
	if u.AuthSource == "" {
		return AuthSourceLocal
	}

	return u.AuthSource
}

func (u *User) BeforeSave(_ *gorm.DB) error {
	u.AuthSource = u.authSource()
	u.GenerateAPIKey(false)
	u.GenerateSalt()

	return u.IsValid()
}

func (u *User) IsValid() error {
	switch u.authSource() {
	case AuthSourceLocal:
		if u.Password == "" {
			return ErrPasswordInvalidLength
		}
	case AuthSourceLDAP:
	default:
		return ErrInvalidAuthSource
	}

	// username validation stays as it is
}

func (u *User) IsActive() bool {
	if u == nil || !u.Active || u.Username == "" {
		return false
	}

	// Keep the old rule for local accounts: no password, no login.
	if u.authSource() == AuthSourceLocal && u.Password == "" {
		return false
	}

	return true
}

func (u *User) ValidLogin(password string) bool {
	if !u.IsActive() || u.authSource() != AuthSourceLocal || u.Password == "" {
		return false
	}

	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(u.AddSalt(password))) == nil
}
```

`SetPassword` refuses to set a local password on an LDAP account:

```go
func (u *User) SetPassword(password string) error {
	if u.authSource() == AuthSourceLDAP {
		return ErrPasswordManagedExternally
	}
	// ...
}
```

An LDAP user can have an empty `password` column; `not null` still holds.

Existing rows migrate to `auth_source = 'local'` through the GORM default. Add a
post-migration update for rows with an empty value as a safety net.

## 2. Configuration

LDAP settings live in the config file or environment only. Not in the admin UI:
the bind password is a secret, and a bad runtime change can lock everyone out.
This matches how DSN and JWT keys are handled.

Add these fields to `EnvConfig` in `pkg/database/config.go`, all with `gorm:"-"`
so the bind password and connection details never end up in the database:

| Key                         | Default | Meaning                                                                             |
| --------------------------- | ------- | ----------------------------------------------------------------------------------- |
| `ldap_enabled`              | `false` | Enable LDAP login                                                                   |
| `ldap_url`                  |         | `ldap://host:389` or `ldaps://host:636`                                             |
| `ldap_start_tls`            | `false` | Use StartTLS on an `ldap://` URL                                                    |
| `ldap_insecure_skip_verify` | `false` | Skip certificate verification; logs a warning                                       |
| `ldap_ca_file`              |         | PEM file with a private CA                                                          |
| `ldap_bind_dn`              |         | Service account DN for search mode                                                  |
| `ldap_bind_password`        |         | Service account password                                                            |
| `ldap_bind_password_file`   |         | File with the service account password                                              |
| `ldap_user_dn`              |         | DN template for direct-bind mode, e.g. `uid={username},ou=people,dc=example,dc=com` |
| `ldap_base_dn`              |         | Search base for search mode                                                         |
| `ldap_user_filter`          |         | Search filter with one `{username}` placeholder, e.g. `(uid={username})`            |
| `ldap_username_attribute`   | `uid`   | Attribute used as the app username                                                  |
| `ldap_name_attribute`       | `cn`    | Attribute used as the display name                                                  |
| `ldap_required_group`       |         | Full DN; user must be a member to log in                                            |
| `ldap_admin_group`          |         | Full DN; members become app admins                                                  |
| `ldap_auto_create`          | `false` | Create unknown LDAP users on first login                                            |
| `ldap_default_active`       | `true`  | `Active` value for newly created LDAP users                                         |
| `ldap_timeout_seconds`      | `10`    | Dial and operation timeout                                                          |

Two modes, exactly one per configuration:

- Search mode: `ldap_bind_dn` + `ldap_bind_password(_file)` + `ldap_base_dn` +
  `ldap_user_filter`.
- Direct-bind mode: `ldap_user_dn` only. The user's entry is read with a base
  search after the bind.

Validation at startup, before the server listens:

- URL has scheme `ldap` or `ldaps` and a host.
- StartTLS is only used with `ldap://`.
- Exactly one of `ldap_bind_dn` and `ldap_user_dn` is set.
- `ldap_user_filter` and `ldap_user_dn` each contain exactly one `{username}`.
- Group values parse as DNs.
- The password file is readable when configured.

Read the password file at startup and strip only a trailing newline, not
arbitrary whitespace: a password may legitimately end with a space. When both
`ldap_bind_password` and `ldap_bind_password_file` are set, the direct value
wins, matching the existing `dsn` / `dsn_file` behaviour.

Invalid config is a startup error, not a login-time surprise. LDAP settings are
read at startup; changing them needs a restart.

With `ldap_auto_create` off, the admin needs a way to create an LDAP user. This
design adds a small create-user form in the admin section, next to the existing
user list. No temporary registration or direct database edit is needed.

OpenLDAP example:

```yaml
ldap_enabled: true
ldap_url: ldaps://ldap.example.com:636
ldap_bind_dn: cn=workout-tracker,ou=services,dc=example,dc=com
ldap_bind_password_file: /run/secrets/ldap-bind-password
ldap_base_dn: ou=people,dc=example,dc=com
ldap_user_filter: "(uid={username})"
ldap_username_attribute: uid
ldap_name_attribute: cn
ldap_required_group: cn=workout-tracker-users,ou=groups,dc=example,dc=com
ldap_admin_group: cn=workout-tracker-admins,ou=groups,dc=example,dc=com
ldap_auto_create: false
```

Active Directory example:

```yaml
ldap_enabled: true
ldap_url: ldaps://dc01.example.com:636
ldap_bind_dn: CN=svc-workout,OU=Service Accounts,DC=example,DC=com
ldap_bind_password_file: /run/secrets/ldap-bind-password
ldap_base_dn: OU=Users,DC=example,DC=com
ldap_user_filter: "(sAMAccountName={username})"
ldap_username_attribute: sAMAccountName
ldap_name_attribute: displayName
ldap_admin_group: CN=Workout Admins,OU=Groups,DC=example,DC=com
ldap_auto_create: false
```

Direct-bind example:

```yaml
ldap_enabled: true
ldap_url: ldaps://ldap.example.com:636
ldap_user_dn: "uid={username},ou=people,dc=example,dc=com"
ldap_username_attribute: uid
ldap_name_attribute: cn
```

Group checks use the `memberOf` attribute only. There is no group search and no
nested-group expansion. In direct-bind mode that only works if the user can read
their own `memberOf`; otherwise required-group and admin mapping need search
mode. Document that.

Environment variables follow the existing `WT_` prefix, e.g. `WT_LDAP_URL`,
`WT_LDAP_BIND_PASSWORD_FILE`.

## 3. LDAP client

New package `pkg/auth`, using `github.com/go-ldap/ldap/v3` (pure Go, no cgo).

```go
type Identity struct {
	Username string
	Name     string
	Admin    bool
}

type Authenticator interface {
	Authenticate(ctx context.Context, username, password string) (*Identity, error)
}
```

The LDAP implementation exposes `NewLDAP(Config) (*LDAP, error)` and keeps the
connection behind a small interface so the tests can use a fake:

```go
type conn interface {
	Bind(username, password string) error
	Search(*ldap.SearchRequest) (*ldap.SearchResult, error)
	StartTLS(*tls.Config) error
	Close() error
}
```

Login flow inside the client:

1. Reject an empty username or password before dialling. An empty password can
   be an unauthenticated bind on some servers.
2. Dial with a timeout. Use `ldaps://`, or `StartTLS` with the configured CA.
3. Search mode:
   - Bind as the service account.
   - Search `ldap_base_dn` with `ldap_user_filter`, replacing `{username}` with
     `ldap.EscapeFilter(username)`, and a size limit of 2.
   - Require exactly one result. Zero or more than one means invalid
     credentials; log the details, return the generic error.
   - Bind as the entry DN with the supplied password.
4. Direct-bind mode:
   - Build the DN from `ldap_user_dn`, replacing `{username}` with an escaped DN
     value.
   - Bind as that DN with the supplied password.
   - Read the entry with a base-scope search for the configured attributes.
5. Read `ldap_username_attribute`, `ldap_name_attribute`, and `memberOf`.
   - Missing username attribute: use the login name.
   - Missing name attribute: use the username.
   - Truncate the name to the `Name` column limit.
   - If the username does not fit the app's username rules, fail the login and
     log a clear configuration error. Never truncate a username.
6. If `ldap_required_group` is set and the user is not a direct member, return
   invalid credentials. If `ldap_admin_group` is set, set `Admin` from
   membership. Compare DNs with `ldap.ParseDN` and `DN.Equal`, not strings.
7. Return the identity. No password is stored or logged.

Error mapping:

- `LDAPResultInvalidCredentials`, no result, multiple results:
  `ErrInvalidCredentials`.
- Network, TLS, service-bind, and other server errors: `ErrLDAPUnavailable`.
- The HTTP handler always shows the existing generic `ErrLoginFailed`; the real
  error goes to the log with the username, never the password.

Filter and DN interpolation must always go through `ldap.EscapeFilter` / escaped
DN values. A test with `*)(uid=*` is mandatory.

## 4. App wiring and login flow

`App` gets one field:

```go
type App struct {
	// ...
	ldapAuth auth.Authenticator
}
```

`Configure()` calls `a.ConfigureAuthentication()` after the database is ready.
It reads the password file, builds `auth.LDAPConfig`, validates it, and sets
`a.ldapAuth`. When `ldap_enabled` is false, the field stays nil.

The sign-in handler stops binding directly into `database.User` and works with
form values:

```go
username := c.FormValue("username")
password := c.FormValue("password")

stored, err := database.GetUserFold(a.db, username)
if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	a.logger.Error("login: lookup user", "error", err)
	return a.redirectWithError(c, a.Reverse("user-login"), ErrLoginFailed)
}

switch {
case stored == nil:
	u, err = a.loginWithLDAP(c, username, password)
case stored.AuthSource == database.AuthSourceLDAP:
	u, err = a.loginWithLDAP(c, stored.Username, password)
default:
	if !stored.ValidLogin(password) {
		err = ErrLoginFailed
	}
	u = stored
}
```

`loginWithLDAP`:

1. Fail if `a.ldapAuth == nil`.
2. Call `Authenticate`.
3. Look up the returned canonical username case-insensitively.
4. If the user exists:
   - It must have `AuthSource == ldap`, otherwise fail. No silent linking.
   - It must be active, otherwise fail.
   - Sync `Name` and, when an admin group is configured, `Admin`.
   - Save only when something changed. A save failure fails the login, so a
     failed admin demotion cannot be ignored.
5. If the user does not exist and `ldap_auto_create` is true:
   - Create the user with `AuthSource = ldap`, `Active = ldap_default_active`,
     `Admin` from the group mapping, and the normal profile defaults.
   - Extract the profile defaults from `userRegisterHandler` into a shared
     helper so registration and JIT provisioning stay in sync.
   - On a unique-constraint race, reload the user and continue as "exists".
6. Otherwise fail.

`ldap_auto_create` defaults to false. In the normal flow the admin has already
created the user with `AuthSource = ldap`, so step 4 is the path that runs. A
successful directory bind for an unknown user does not create an app account.

Factor the last part of both logins into a `finishLogin` helper: put the
username in the session, create the JWT cookie, redirect to the dashboard. The
JWT and session format do not change, so API keys and existing sessions keep
working.

LDAP users are case-insensitive in the directory. Use a case-insensitive lookup
for LDAP login only, and never auto-link a case variant of an existing local
user. A database-level case-insensitive unique index is a possible follow-up; it
is not needed for the first version.

A renamed directory user is treated as a different person: the old app user
remains, and auto-create may make a new one. Storing the directory's immutable
ID (`entryUUID` / `objectGUID`) is the follow-up that fixes this.

## 5. Admin UI

### Source selection

Use the existing `partials.RadioPillGroup` helper for the source selection. It
already renders radio inputs as the same pills used by the statistics filters:

```templ
@partials.RadioPillGroup("admin.auth-source", "auth_source", "pill-compact", []partials.RadioPillOption{
	{Value: database.AuthSourceLocal, LabelKey: "admin.auth-source-local", Checked: u.AuthSource != database.AuthSourceLDAP},
	{Value: database.AuthSourceLDAP, LabelKey: "admin.auth-source-ldap", Checked: u.AuthSource == database.AuthSourceLDAP},
})
```

No new CSS is needed.

### Edit user

`views/admin/edit_user.templ` gets the source pills and behaves like this:

- `local`: password field as it is today, with the existing "leave blank to keep
  current password" hint.
- `ldap`: the password field stays visible but is ignored; a note says the
  account is managed by the directory. `ldap_admin_group` overrides the admin
  checkbox on the next login.
- The username stays editable. It must match the directory username exactly (the
  lookup is case-insensitive, but the stored value is the directory value). Show
  that as a hint.
- Switching local -> LDAP clears the local password and salt. No password is
  required for the switch.
- Switching LDAP -> local requires a non-empty password in the same submit. Set
  the source first, then call `SetPassword`.
- Switching local -> local with an empty password keeps the existing hash.
- A warning reminds the admin to keep at least one local admin as break-glass.
- When LDAP is disabled in the server configuration, a note says that an LDAP
  user cannot sign in until it is enabled.

### Create user

Because auto-create is off, add a small create-user flow:

- `GET /admin/users/new` -> `admin.NewUser` form.
- `POST /admin/users` -> `adminUserCreateHandler`.
- A "New user" button on `views/admin/root.templ`.
- The form reuses the edit-user markup or a shared component. Fields: username,
  name, source pills, active, admin, and a password field that is only relevant
  for local users.
- Defaults: active on, admin off, source local.
- For an LDAP user: no password is set, the row gets
  `AuthSource = AuthSourceLDAP`, and the next login goes through the directory.
- For a local user: a password is required.
- Profile defaults come from the shared helper so the new account looks like a
  registered account.

### Configuration

LDAP connection settings are not editable in the admin UI. Optionally show a
read-only "LDAP enabled" status row to help with debugging. Secrets stay in the
config file or environment.

## 6. Migration

- `AutoMigrate` adds `auth_source` with default `local`.
- Post-migration:
  `UPDATE users SET auth_source = 'local' WHERE auth_source = ''`.
- Existing users, tokens, API keys, and passwords are untouched.
- Existing users are converted one by one through the admin edit page: pick the
  LDAP pill, check the username, save. Their local password is cleared.
- New LDAP users are created through the admin create-user form and start with
  the LDAP pill selected.
- Run `make swagger` because `UserData` is part of the API output.
- `workout-tracker.example.yaml` gets a commented LDAP block.
- README gets an "LDAP authentication" section with the examples above, the
  security notes, and troubleshooting.

## 7. Security checklist

- No automatic linking of existing local accounts. This is the main takeover
  risk.
- Escape filter and DN values. Never interpolate raw input.
- Reject empty passwords before a bind.
- TLS by default. `ldap_insecure_skip_verify` is opt-in and warns at startup.
- Never log bind passwords, user passwords, or full config structs.
- Fail closed on LDAP errors.
- Use a read-only service account, with the minimum rights.
- Group DNs are compared as DNs, not strings.
- Keep at least one local admin as break-glass.
- Warn at startup when LDAP is enabled and registration is not disabled.
  Otherwise someone can register a local account with the same username before
  the LDAP user first logs in.
- There is no rate limiting on `/user/signin` today, and this work does not add
  one. Direct binds can trigger AD account lockout policies. This is an accepted
  gap for now and should be handled separately.
- Disabling a user in LDAP stops new logins, but existing JWT cookies remain
  valid until they expire (currently 10 days). Deactivate the app user to revoke
  access immediately. A periodic sync is out of scope.
- Group changes are applied on the next login. An existing session keeps the
  stored `Admin` flag until it expires or the user signs in again. The same
  applies to `ldap_required_group`.
- API keys are not checked against LDAP. Deactivate the app user to revoke an
  API key.

## 8. Tests

Database/model:

- Local user: valid password required, `ValidLogin` works.
- LDAP user: empty password allowed, `ValidLogin` always false, `SetPassword`
  fails.
- `AuthSource` defaults to local on create and on migrated rows.

LDAP client (fake connection):

- Valid credentials return the identity.
- Empty password never dials.
- No result and multiple results return invalid credentials.
- Invalid user bind returns invalid credentials.
- Network/TLS errors return unavailable.
- Filter and DN escaping.
- Admin group and required group membership.
- Missing `memberOf` fails closed when a group is configured.

App/login:

- LDAP login creates a user with profile defaults when auto-create is on
  (covering the optional setting).
- LDAP login does not take over an existing local user.
- LDAP login fails for an unknown user when auto-create is off (the default).
- Inactive LDAP user cannot log in.
- Local login still works with LDAP enabled.
- LDAP errors do not leak details and do not create users.
- Successful login sets the token cookie and redirects.

Admin:

- Create a local user: password required, profile defaults applied.
- Create an LDAP user: no password, source stored as `ldap`.
- Duplicate username is rejected.
- Convert local -> LDAP clears the password; local login stops working.
- Convert LDAP -> local requires a new password; local login starts working.
- The pill selection renders the right checked state.

Config:

- Env/file parsing.
- Invalid URL, both modes set, missing placeholder, unreadable password file.

Manual:

- OpenLDAP in Docker with a test LDIF, login with a normal user and an admin.
- AD or a similar directory over LDAPS with a private CA.
- StartTLS and `ldap_insecure_skip_verify` only on a test setup.

## 9. File-by-file

- `pkg/database/user.go`: `AuthSource`, constants, invariant changes,
  case-insensitive lookup for LDAP.
- `pkg/database/config.go`: LDAP fields in `EnvConfig`.
- `pkg/database/gorm.go`: post-migration default.
- `pkg/auth/auth.go`: `Identity` and the `Authenticator` interface.
- `pkg/auth/ldap.go`: client, config validation, DN/group handling.
- `pkg/auth/ldap_test.go`: fake-connection tests.
- `pkg/app/app.go`: `ldapAuth` field.
- `pkg/app/config.go`: defaults, env bindings, validation hook.
- `pkg/app/auth_external.go`: LDAP configuration, login, provisioning,
  `finishLogin`.
- `pkg/app/users_handlers.go`: sign-in and registration changes.
- `pkg/app/admin_handlers.go`: create user, auth source conversion, password
  rules.
- `pkg/app/routes.go` (via `adminRoutes`): `GET /admin/users/new`,
  `POST /admin/users`.
- `pkg/app/*_test.go`: handler tests with a fake authenticator.
- `views/admin/edit_user.templ`: auth source pills and notes.
- `views/admin/new_user.templ` (or a shared form component): create form.
- `views/admin/root.templ`: "New user" button.
- `views/user/login.templ`: optional directory hint.
- `translations/en.yaml`, `translations/nl.yaml`: new admin/login strings.
- `README.md`, `workout-tracker.example.yaml`: configuration and security docs.
- `docs/swagger.json`, `docs/swagger.yaml`: regenerated.

## 10. Rollout

1. Model and config foundation, with local behaviour unchanged.
2. LDAP client and unit tests.
3. Login integration and handler tests.
4. Admin create-user flow, source pills, conversion, and translations; run
   `make build-templates`.
5. Docs, example config, swagger.
6. Optional follow-ups: rate limiting, `REMOTE_USER` provider, directory sync,
   immutable directory IDs.

Add the dependency with `go get github.com/go-ldap/ldap/v3`, then run
`make go-vendor`. Use `make test-go` before calling it done.

`ldap_enabled` defaults to false, so the first four steps can merge without
changing any existing deployment.

## 11. Locked decisions

1. Direct LDAP in the app. `REMOTE_USER` is not part of this work.
2. `ldap_auto_create` defaults to false. Admins create LDAP users explicitly or
   convert existing local users.
3. Group checks use `memberOf` only. No group search, no nested groups.
4. Existing users are convertible in the admin UI through a source
   pill-selection.
5. No rate limiting in this work. The existing login endpoint stays as it is.
