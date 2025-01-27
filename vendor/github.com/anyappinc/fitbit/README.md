fitbit - Fitbit API client written in go
=======

fitbit package provides the client to communicate with [Firbit Web API](https://dev.fitbit.com/build/reference/web-api/).

This implementation contains the structure to interact with Fitbit Web API,
but only some APIs(endpoints) have been implemented right now.

For more details, see [Implemented APIs](#implemented-apis) section.


## Installation

```
go get github.com/anyappinc/fitbit
```


## Features

- Obtaining tokens through a secure OAuth2 authentication process.
  + This package follows Authorization Code Grant Flow with Proof Key for Code Exchange (PKCE) defined by RFC 7636, which is Fitbit's recommended option.
- The configurable client. You can specify the application type(Server/Client/Personal), locale, language, and scopes.
- Auto-refreshing of an access token using a refresh token when needed.
  + And the hook function is configurable so that you can observe a token refreshing.
- Easy access to the rate limit.
  + For more details, see https://dev.fitbit.com/build/reference/web-api/developer-guide/application-design/#Rate-Limits.


### Implemented APIs

- [Authorization](https://dev.fitbit.com/build/reference/web-api/authorization/)
  + [Authorize](https://dev.fitbit.com/build/reference/web-api/authorization/authorize/)
  + [OAuth2 Token](https://dev.fitbit.com/build/reference/web-api/authorization/oauth2-token/)
  + [Introspect](https://dev.fitbit.com/build/reference/web-api/authorization/introspect/)
  + [Refresh Token](https://dev.fitbit.com/build/reference/web-api/authorization/refresh-token/)
  + [Revoke Token](https://dev.fitbit.com/build/reference/web-api/authorization/revoke-token/)
- [Activity](https://dev.fitbit.com/build/reference/web-api/activity/)
  + [Get Daily Activity Summary](https://dev.fitbit.com/build/reference/web-api/activity/get-daily-activity-summary/)
- [Nutrition](https://dev.fitbit.com/build/reference/web-api/nutrition/)
  + [Get Water Log](https://dev.fitbit.com/build/reference/web-api/nutrition/get-water-log/)
- [User](https://dev.fitbit.com/build/reference/web-api/user/)
  + [Get Profile](https://dev.fitbit.com/build/reference/web-api/user/get-profile/)


### Debug Mode

When debug mode is on, the consent dialog appears every time when users try to authorize.

`Client` has functions below to change debug mode.

- `EnableDebugMode()`
- `DisableDebugMode()`


## Example

```go
package main

import (
  "context"
  "encoding/json"
  "io"
  "log"
  "net/http"
  "strconv"

  "github.com/anyappinc/fitbit"
)

const (
  clientID     = "xxxxxx"
  clientSecret = "******"
  redirectURI  = "http://localhost:8080/link"
)

var (
  fitbitClient *fitbit.Client
  state        string
  codeVerifier string
  userID       string
  token        *fitbit.Token
)

func updateTokenFunc(oldToken, newToken *fitbit.Token) error {
  log.Print("Token updated.")
  return nil
}

func init() {
  fitbitClient = fitbit.NewClient(clientID, clientSecret, fitbit.ServerApplication, &fitbit.Scope{
    Location: true,
    Profile:  true,
    Weight:   true,
  })
  fitbitClient.SetLocaleAndLanguage(fitbit.LocaleJapan)
  fitbitClient.SetUpdateTokenFunc(updateTokenFunc)
  fitbitClient.EnableDebugMode()
}

func main() {
  http.HandleFunc("/authorize", func(w http.ResponseWriter, req *http.Request) {
    authCodeURL, _state, _codeVerifier := fitbitClient.AuthCodeURL(redirectURI)
    state = _state
    codeVerifier = _codeVerifier
    http.Redirect(w, req, authCodeURL.String(), http.StatusSeeOther)
  })

  http.HandleFunc("/link", func(w http.ResponseWriter, req *http.Request) {
    requestQuery := req.URL.Query()
    if requestQuery.Has("error") {
      http.Error(w, requestQuery.Get("error"), http.StatusInternalServerError)
      return
    }
    if requestQuery.Get("state") != state {
      http.Error(w, "state mismatched.", http.StatusBadRequest)
      return
    }
    ctx := context.Background()
    linkResp, err := fitbitClient.Link(ctx, requestQuery.Get("code"), codeVerifier, redirectURI)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    userID = linkResp.UserID
    token = linkResp.Token
    io.WriteString(w, "ok")
  })

  http.HandleFunc("/profile", func(w http.ResponseWriter, req *http.Request) {
    ctx := context.Background()
    profile, _, _, err := fitbitClient.GetProfile(ctx, userID, token)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }
    jsonBytes, _ := json.Marshal(profile)
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
  })

  portString := ":" + strconv.Itoa(8080)
  log.Print("Listening ", portString)
  log.Fatal(http.ListenAndServe(portString, nil))
}
```


## Notes

Most lines of `oauth2_internal.go` were adapted from https://go.googlesource.com/oauth2/+/refs/heads/master/internal, which is distributed under BSD-3-Clause, to customize behavior on token refresh.  
The original license is available at https://go.googlesource.com/oauth2/+/refs/heads/master/LICENSE

This chunk is used instead of the corresponding part of `oauth2` package when the hook function on token auto-refreshing is configured.
So this may cause different behavior from the original one. For example, in fact, this does not do any special handling for App Engine.


## License

TBD

