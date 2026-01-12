# Agent Guide for Workout Tracker

Welcome to the Workout Tracker repository. This guide provides essential information for AI agents to work effectively on this codebase.

## 🚀 Essential Commands

### Backend (Go)
- **Build Server**: `make build-server` (Outputs to `tmp/workout-tracker`)
- **Run Tests**: `make test-go` (Includes package, template, and command tests)
- **Lint**: `golangci-lint run` (or `make test-go` which includes it)
- **Development Mode**: `make dev-backend` (Uses `air` for live reload and `templ generate --watch`)
- **Generate Templates**: `make build-templates` (Runs `templ generate`)
- **Format Templates**: `make format-templates` (Runs `templ fmt`)

### Frontend (Lit & Tailwind)
- **Install Deps**: `make install-deps` (Runs `npm install` in `frontend/`)
- **Build Assets**: `make build-frontend` (Builds JS/CSS)
- **Watch Assets**: `make watch/tailwind` or `make watch/esbuild`
- **Sync Translations**: `cd frontend && npm run translations`

### Full Development
- **Docker Compose**: `make dev-docker` (Starts PostgreSQL and the app)
- **Cleanup**: `make dev-docker-clean`

### Visuals
- **Track Animation**: `make track-gif` (Generates `docs/track.gif` using k6 and ImageMagick)

---

## 🏗 Code Organization

- `cmd/`: Entry points for the server (`workout-tracker`), debug CLI (`wt-debug`), and Fitbit sync.
- `pkg/`: Core application logic.
- `pkg/app/`: HTTP handlers, routes, and application state.
- `pkg/database/`: GORM models and database interactions.
- `pkg/converters/`: GPX/FIT/TCX parsing logic.
- `views/`: UI components using **a-h/templ**.
- `frontend/`: Frontend source code (Lit components, TypeScript, Tailwind CSS).
- `assets/`: Generated static assets (JS/CSS).

### Tools
- **Nix**: Use `nix develop` or `direnv allow` for a pre-configured environment.
- **Swagger**: Run `make swagger` to update documentation.

---

## 🛠 Tech Stack & Patterns

### Backend
- **Framework**: [Echo](https://echo.labstack.com/) for routing and middleware.
- **ORM**: [GORM](https://gorm.io/) with support for SQLite, MySQL, and PostgreSQL.
- **Templates**: [templ](https://templ.guide/) for type-safe HTML components.
- **Database Migrations**: Handled automatically via GORM's `AutoMigrate` in `pkg/database/gorm.go`.

### Frontend
- **HTMX**: Used for dynamic UI updates without full page reloads.
- **Lit**: Web Components for complex interactive elements (like maps).
- **Tailwind CSS**: For styling.
- **esbuild**: For bundling frontend assets.

### I18n
- Backend uses `github.com/invopop/ctxi18n`.
- Translations are in `translations/*.yaml`.
- Frontend has its own localization in `frontend/src/generated/locales/`.

---

## 📝 Conventions & Style

- **Error Handling**: Use `a.redirectWithError(c, url, err)` in web handlers or `a.renderAPIError(c, resp, err)` in API handlers.
- **Templates**: Always run `make build-templates` after modifying `.templ` files.
- **Testing**:
  - Prefer `pkg/app/api_handlers_test.go` style for integration tests.
  - Use `test-short` flag for quick iterations.
- **Mocking**: The project uses `configuredApp(t)` in tests to provide a ready-to-use application instance with an in-memory database.
- **Git Commits**: Use conventional commit format. Use `git commit --verbose --signoff --gpg-sign`.

---

## ⚠️ Gotchas & Non-Obvious Patterns

- **Generated Files**: `*_templ.go` files are generated. Do not edit them directly; edit the corresponding `.templ` file and run `make build-templates`.
- **HTMX Redirects**: For HTMX requests, use `Hx-Redirect` header instead of standard 302 redirects if you want the whole page to navigate.
- **Workout Processing**: Workouts are often marked as "Dirty" and processed in the background (see `BackgroundWorker` in `pkg/app/background.go`).
- **GPX Data**: Stored as raw bytes in the `gpx_data` table, while processed track points are in `map_data_details`.
- **CSRF/Security**: The app uses `echo-jwt` for authentication and `scs` for session management.
- **API Keys**: Users can have API keys (found in `Profile.APIKey`). The API expects `Authorization: Bearer <key>` or `?api-key=<key>`.
- **Offline Mode**: The app has an offline mode (`Config.Offline`) that disables external geocoding services.
