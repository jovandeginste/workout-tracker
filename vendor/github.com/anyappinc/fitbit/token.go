package fitbit

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

var (
	// CSRFStateLength represents the length of `state` generating on authorization process.
	CSRFStateLength uint64 = 128

	// CodeVerifierLength represents the length of `code_verifier` generating on authorization process.
	CodeVerifierLength uint64 = 128
)

// Token represents the OAuth 2.0 Token.
type Token struct {
	AccessToken  string
	TokenType    string
	RefreshToken string
	Expiry       time.Time
}

func (t *Token) asOAuth2Token() *oauth2.Token {
	if t == nil {
		return nil
	}
	return &oauth2.Token{
		AccessToken:  t.AccessToken,
		TokenType:    t.TokenType,
		RefreshToken: t.RefreshToken,
		Expiry:       t.Expiry,
	}
}

type tokenRefresher struct {
	ctx       context.Context
	client    *Client
	lastToken *Token
}

// Token implements the the oauth2.TokenSource interface.
func (tkr *tokenRefresher) Token() (*oauth2.Token, error) {
	token, err := retrieveToken(
		tkr.ctx,
		tkr.client.oauth2Config.ClientID,
		tkr.client.oauth2Config.ClientSecret,
		tkr.client.oauth2Config.Endpoint.TokenURL,
		url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {tkr.lastToken.RefreshToken},
		},
		tkr.client.applicationType,
	)
	if err != nil {
		return nil, err
	}
	if tkr.client.updateTokenFunc != nil {
		if err := tkr.client.updateTokenFunc(tkr.lastToken, token); err != nil {
			return nil, err
		}
	}
	tkr.lastToken = token
	return token.asOAuth2Token(), err
}

type (
	// LinkResponse represents a response of the link request.
	LinkResponse struct {
		UserID string
		Scope  *Scope
		Token  *Token
	}

	rawTokenState struct {
		Active    bool   `json:"active"`
		UserID    string `json:"user_id"`
		TokenType string `json:"token_type"`
		Scope     string `json:"scope"`
		Iat       int64  `json:"iat"`
		Exp       int64  `json:"exp"`
		ClientID  string `json:"client_id"`
	}

	// TokenState represents the active state of an access token
	TokenState struct {
		Active         bool
		UserID         string
		TokenType      string
		Scope          *Scope
		ScopeType      ScopeType
		IssuedDate     *time.Time
		ExpirationDate *time.Time
		ClientID       string
	}
)

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TokenState) UnmarshalJSON(b []byte) error {
	var raw rawTokenState
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	scope, scopeType := parseScopeFromTokenState(raw.Scope)

	// NOTE: `time.UnixMilli` had been added in go1.17,
	// but just use the traditional way for wider support
	issuedDate := time.Unix(raw.Iat/1e3, (raw.Iat%1e3)*1e6)
	expirationDate := time.Unix(raw.Exp/1e3, (raw.Exp%1e3)*1e6)

	t.Active = raw.Active
	t.UserID = raw.UserID
	t.TokenType = raw.TokenType
	t.Scope = scope
	t.ScopeType = scopeType
	t.IssuedDate = &issuedDate
	t.ExpirationDate = &expirationDate
	t.ClientID = raw.ClientID
	return nil
}

// AuthCodeURL returns an url to link with user's Fitbit account.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/authorization/authorize/
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/developer-guide/authorization/
func (c *Client) AuthCodeURL(redirectURI string) (*url.URL, string, string) {
	state := string(randomBytes(CSRFStateLength))
	codeVerifier := randomBytes(CodeVerifierLength)
	hashedCodeVerifier := sha256.Sum256(codeVerifier)
	codeChallenge := base64.RawURLEncoding.EncodeToString(hashedCodeVerifier[:])
	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
		oauth2.SetAuthURLParam("code_challenge_method", CodeChallengeMethod),
		oauth2.SetAuthURLParam("redirect_uri", redirectURI),
	}
	if c.debugMode {
		opts = append(opts, oauth2.ApprovalForce)
	}
	urlString := c.oauth2Config.AuthCodeURL(state, opts...)
	authCodeURL, _ := url.Parse(urlString) // error should never happen
	return authCodeURL, state, string(codeVerifier)
}

// Link obtains data for the user to interact with Fitbit APIs.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/authorization/oauth2-token/
func (c *Client) Link(ctx context.Context, code, codeVerifier, reqURIString string) (*LinkResponse, error) {
	opts := []oauth2.AuthCodeOption{
		oauth2.SetAuthURLParam("code_verifier", codeVerifier),
		oauth2.SetAuthURLParam("redirect_uri", reqURIString),
	}
	if c.applicationType == ServerApplication {
		// `client_id` parameter seems unnecessary, but add this just to make sure
		// since this is noted "required" in the official document
		opts = append(opts, oauth2.SetAuthURLParam("client_id", c.oauth2Config.ClientID))
	}
	token, err := c.oauth2Config.Exchange(ctx, code, opts...)
	if err != nil {
		if rErr := (*oauth2.RetrieveError)(nil); errors.As(err, &rErr) {
			if e := parseError(rErr.Response, rErr.Body); e != nil {
				return nil, fmt.Errorf("fitbit(oauth2): cannot fetch token: %w", e)
			}
		}
		return nil, fmt.Errorf("fitbit(oauth2): cannot fetch token: %w", err)
	}
	return &LinkResponse{
		UserID: token.Extra("user_id").(string),
		Scope:  newScope(strings.Split(token.Extra("scope").(string), " ")),
		Token: &Token{
			AccessToken:  token.AccessToken,
			TokenType:    token.TokenType,
			RefreshToken: token.RefreshToken,
			Expiry:       token.Expiry,
		},
	}, nil
}

// IntrospectToken retrieves the active state of an OAuth 2.0 token.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/authorization/introspect/
func (c *Client) IntrospectToken(ctx context.Context, token *Token) (*TokenState, *RateLimit, []byte, error) {
	endpoint := c.getEndpoint("IntrospectToken")
	values := url.Values{}
	values.Set("token", token.AccessToken)
	b, rateLimit, err := c.postRequest(ctx, token, endpoint, values)
	if err != nil {
		return nil, nil, b, err
	}
	var tokenState TokenState
	if err := json.Unmarshal(b, &tokenState); err != nil {
		return nil, rateLimit, b, err
	}
	return &tokenState, rateLimit, b, nil
}

// RevokeAccessToken disables the user's authorizations and all tokens,
// associated with the specified access token.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/authorization/revoke-token/
func (c *Client) RevokeAccessToken(ctx context.Context, token *Token) (*RateLimit, error) {
	return c.revokeToken(ctx, token, token.AccessToken)
}

// RevokeRefreshToken disables the user's authorizations and all tokens,
// associated with the specified refresh token.
//
// Web API Reference: https://dev.fitbit.com/build/reference/web-api/authorization/revoke-token/
func (c *Client) RevokeRefreshToken(ctx context.Context, token *Token) (*RateLimit, error) {
	return c.revokeToken(ctx, token, token.RefreshToken)
}

func (c *Client) revokeToken(ctx context.Context, token *Token, target string) (*RateLimit, error) {
	endpoint := c.getEndpoint("RevokeToken")
	values := url.Values{}
	values.Set("token", target)
	if c.applicationType != ServerApplication {
		values.Set("client_id", c.oauth2Config.ClientID)
	}
	_, rateLimit, err := c.postRequest(ctx, token, endpoint, values)
	if err != nil {
		return nil, err
	}
	return rateLimit, nil
}
