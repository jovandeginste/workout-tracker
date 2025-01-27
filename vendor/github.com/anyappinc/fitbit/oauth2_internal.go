package fitbit

// NOTE: Most lines of this code were adapted from
// https://go.googlesource.com/oauth2/+/refs/heads/master/internal
// which is distributed under BSD-3-Clause,
// to customize behavior on token refresh.
// The original license is available at https://go.googlesource.com/oauth2/+/refs/heads/master/LICENSE

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/context/ctxhttp"
	"golang.org/x/oauth2"
)

// ContextKey is just an empty struct. It exists so HTTPClient can be
// an immutable public variable with a unique type. It's immutable
// because nobody else can create a ContextKey, being unexported.
type ContextKey struct{}

// HTTPClient is the context key to use with golang.org/x/net/context's
// WithValue function to associate an *http.Client value with a context.
var HTTPClient ContextKey

func newTokenRequest(tokenURL, clientID, clientSecret string, v url.Values, appType ApplicationType) (*http.Request, error) {
	if appType != ServerApplication {
		v.Set("client_id", clientID)
	}
	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if appType == ServerApplication {
		req.SetBasicAuth(url.QueryEscape(clientID), url.QueryEscape(clientSecret))
	}
	return req, nil
}

func contextClient(ctx context.Context) *http.Client {
	if ctx != nil {
		if hc, ok := ctx.Value(HTTPClient).(*http.Client); ok {
			return hc
		}
	}
	return http.DefaultClient
}

type tokenJSON struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

func (e *tokenJSON) expiry() (t time.Time) {
	if v := e.ExpiresIn; v != 0 {
		return time.Now().Add(time.Duration(v) * time.Second)
	}
	return
}

func doTokenRoundTrip(ctx context.Context, req *http.Request) (*Token, error) {
	r, err := ctxhttp.Do(ctx, contextClient(ctx), req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1<<20))
	r.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("fitbit(oauth2): cannot fetch token: %v", err)
	}
	if code := r.StatusCode; code < 200 || code > 299 {
		return nil, &oauth2.RetrieveError{
			Response: r,
			Body:     body,
		}
	}

	// It is expected that the response's 'content-Type' is `application/json`
	var tj tokenJSON
	if err = json.Unmarshal(body, &tj); err != nil {
		return nil, err
	}
	return &Token{
		AccessToken:  tj.AccessToken,
		TokenType:    tj.TokenType,
		RefreshToken: tj.RefreshToken,
		Expiry:       tj.expiry(),
	}, nil
}

func retrieveToken(ctx context.Context, clientID, clientSecret, tokenURL string, v url.Values, appType ApplicationType) (*Token, error) {
	req, err := newTokenRequest(tokenURL, clientID, clientSecret, v, appType)
	if err != nil {
		return nil, err
	}
	return doTokenRoundTrip(ctx, req)
}
