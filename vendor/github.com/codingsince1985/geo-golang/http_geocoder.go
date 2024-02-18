package geo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// DefaultTimeout for the request execution
const DefaultTimeout = time.Second * 8

// ErrTimeout occurs when no response returned within timeoutInSeconds
var ErrTimeout = errors.New("TIMEOUT")

// EndpointBuilder defines functions that build urls for geocode/reverse geocode
type EndpointBuilder interface {
	GeocodeURL(string) string
	ReverseGeocodeURL(Location) string
}

// ResponseParserFactory creates a new ResponseParser
type ResponseParserFactory func() ResponseParser

// ResponseParser defines functions that parse response of geocode/reverse geocode
type ResponseParser interface {
	Location() (*Location, error)
	Address() (*Address, error)
}

// HTTPGeocoder has EndpointBuilder and ResponseParser
type HTTPGeocoder struct {
	EndpointBuilder
	ResponseParserFactory
}

// Geocode returns location for address
func (g HTTPGeocoder) Geocode(address string) (*Location, error) {
	responseParser := g.ResponseParserFactory()

	ctx, cancel := context.WithTimeout(context.TODO(), DefaultTimeout)
	defer cancel()

	type geoResp struct {
		l *Location
		e error
	}
	ch := make(chan geoResp, 1)

	go func(ch chan geoResp) {
		if err := response(ctx, g.GeocodeURL(url.QueryEscape(address)), responseParser); err != nil {
			ch <- geoResp{
				l: nil,
				e: err,
			}
		}

		loc, err := responseParser.Location()
		ch <- geoResp{
			l: loc,
			e: err,
		}
	}(ch)

	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case res := <-ch:
		return res.l, res.e
	}
}

// ReverseGeocode returns address for location
func (g HTTPGeocoder) ReverseGeocode(lat, lng float64) (*Address, error) {
	responseParser := g.ResponseParserFactory()

	ctx, cancel := context.WithTimeout(context.TODO(), DefaultTimeout)
	defer cancel()

	type revResp struct {
		a *Address
		e error
	}
	ch := make(chan revResp, 1)

	go func(ch chan revResp) {
		if err := response(ctx, g.ReverseGeocodeURL(Location{lat, lng}), responseParser); err != nil {
			ch <- revResp{
				a: nil,
				e: err,
			}
		}

		addr, err := responseParser.Address()
		ch <- revResp{
			a: addr,
			e: err,
		}
	}(ch)

	select {
	case <-ctx.Done():
		return nil, ErrTimeout
	case res := <-ch:
		return res.a, res.e
	}
}

// Response gets response from url
func response(ctx context.Context, url string, obj ResponseParser) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	body := strings.Trim(string(data), " []")
	DebugLogger.Printf("Received response: %s\n", body)
	if body == "" {
		return nil
	}
	if err := json.Unmarshal([]byte(body), obj); err != nil {
		ErrLogger.Printf("Error unmarshalling response: %s\n", err.Error())
		return err
	}

	return nil
}

// ParseFloat is a helper to parse a string to a float
func ParseFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}
