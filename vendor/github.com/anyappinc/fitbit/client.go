package fitbit

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/anyappinc/fitbit/logger"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
)

// Client is a client to interact with Fitbit APIs.
type Client struct {
	oauth2Config    *oauth2.Config
	locale          Locale
	language        Locale
	applicationType ApplicationType
	updateTokenFunc func(*Token, *Token) error
	debugMode       bool
}

// NewClient initializes Fitbit API Client.
//
// When `applicationType` is ServerApplication, `clientSecret` is required.
//
// When `applicationType` is ClientApplication or PersonalApplication,
// `clientSecret` should be empty.
func NewClient(clientID, clientSecret string, applicationType ApplicationType, scope *Scope) *Client {
	fitbitEndpoint := endpoints.Fitbit
	if applicationType == ServerApplication {
		if clientSecret == "" {
			logger.Warn.Print("Client Secret is not set. ServerApplication requires Client Secret.")
		}
		fitbitEndpoint.AuthStyle = oauth2.AuthStyleInHeader
	} else {
		if clientSecret != "" {
			logger.Warn.Print("Client Secret is set. Client Secret should not be set, except for ServerApplication.")
		}
		fitbitEndpoint.AuthStyle = oauth2.AuthStyleInParams
	}
	return &Client{
		oauth2Config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			Scopes:       scope.convert(),
			Endpoint:     fitbitEndpoint,
		},
		locale:          LocaleUnitedStates, // default setting. See https://dev.fitbit.com/build/reference/web-api/developer-guide/application-design/#Language
		applicationType: applicationType,
	}
}

// SetLocale sets locale.
// This value is used to set Accept-Locale header.
//
// See more details https://dev.fitbit.com/build/reference/web-api/developer-guide/application-design/#Language
func (c *Client) SetLocale(locale Locale) {
	c.locale = locale
}

// SetLanguage sets language.
// This value is used to set Accept-Language header.
//
// See more details https://dev.fitbit.com/build/reference/web-api/developer-guide/application-design/#Unit-Systems
func (c *Client) SetLanguage(locale Locale) {
	c.language = locale
}

// SetLocaleAndLanguage just calls both `SetLocale` and `SetLanguage`.
func (c *Client) SetLocaleAndLanguage(locale Locale) {
	c.SetLocale(locale)
	c.SetLanguage(locale)
}

// GetUnit returns Unit that corresponds the current language setting.
func (c *Client) GetUnit() *Unit {
	return getCorrespondingUnit(c.language)
}

// SetUpdateTokenFunc sets the function to be invoked when a token is updated.
func (c *Client) SetUpdateTokenFunc(f func(*Token, *Token) error) {
	c.updateTokenFunc = f
}

// EnableDebugMode enables debug mode
func (c *Client) EnableDebugMode() {
	c.debugMode = true
}

// DisableDebugMode disables debug mode
func (c *Client) DisableDebugMode() {
	c.debugMode = false
}

func (c *Client) getRequest(ctx context.Context, token *Token, url string) ([]byte, *RateLimit, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	b, rateLimit, err := c.request(ctx, token, req)
	return b, rateLimit, wrapAsRequestError("Get", url, err)
}

func (c *Client) postRequest(ctx context.Context, token *Token, url string, data url.Values) ([]byte, *RateLimit, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	b, rateLimit, err := c.request(ctx, token, req)
	return b, rateLimit, wrapAsRequestError("Post", url, err)
}

func (c *Client) getEndpoint(label string, params ...interface{}) string {
	return fmt.Sprintf(apiBaseURL+apiEndpoints[label], params...)
}

func (c *Client) newHTTPClient(ctx context.Context, token *Token) *http.Client {
	if c.updateTokenFunc != nil {
		return oauth2.NewClient(ctx, c.tokenSource(ctx, token))
	}
	return c.oauth2Config.Client(ctx, token.asOAuth2Token())
}

func (c *Client) tokenSource(ctx context.Context, token *Token) oauth2.TokenSource {
	t, tkr := token.asOAuth2Token(), &tokenRefresher{
		ctx:       ctx,
		client:    c,
		lastToken: token,
	}
	return oauth2.ReuseTokenSource(t, tkr)
}

func (c *Client) request(ctx context.Context, token *Token, req *http.Request) ([]byte, *RateLimit, error) {
	httpClient := c.newHTTPClient(ctx, token)
	req.Header.Set("Accept-Locale", c.locale.asString())
	req.Header.Set("Accept-Language", c.language.asString())
	resp, err := httpClient.Do(req)
	if err != nil {
		if uErr := (*url.Error)(nil); errors.As(err, &uErr) {
			if rErr := (*oauth2.RetrieveError)(nil); errors.As(uErr, &rErr) {
				if e := parseError(rErr.Response, rErr.Body); e != nil {
					return nil, nil, fmt.Errorf("fitbit(oauth2): cannot fetch token: %w", e)
				}
				return nil, nil, errors.New("fitbit(oauth2): cannot fetch token")
			}
			return nil, nil, uErr.Unwrap()
		}
		return nil, nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	rateLimit := extractRateLimit(&resp.Header)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return b, rateLimit, parseError(resp, b)
	}
	return b, rateLimit, nil
}
