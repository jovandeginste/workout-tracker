package geocoder

// Inspired by https://github.com/codingsince1985/geo-golang

import (
	"cmp"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/codingsince1985/geo-golang"
	"github.com/google/go-querystring/query"
)

var (
	c                  *client
	ErrClientNotSet    = errors.New("geocoder: client not set")
	ErrAddressNotFound = errors.New("geocoder: address not found")
)

const requestInterval = time.Second

type client struct {
	url         string
	client      http.Client
	logger      *slog.Logger
	lastRequest time.Time
	userAgent   string
	m           sync.Mutex
}

type Query struct {
	Format string  `url:"format"`
	Lat    float64 `url:"lat"`
	Lon    float64 `url:"lon"`
}

type Result struct {
	Address     Address  `json:"address"`
	Licence     string   `json:"licence"`
	OsmType     string   `json:"osm_type"`
	Lat         string   `json:"lat"`
	Lon         string   `json:"lon"`
	Category    string   `json:"category"`
	Type        string   `json:"type"`
	Addresstype string   `json:"addresstype"`
	DisplayName string   `json:"display_name"`
	Name        string   `json:"name"`
	Boundingbox []string `json:"boundingbox"`
	PlaceID     int      `json:"place_id"`
	OsmID       int      `json:"osm_id"`
	PlaceRank   int      `json:"place_rank"`
	Importance  float64  `json:"importance"`
}

type Address struct {
	HouseNumber   string `json:"house_number"`
	Road          string `json:"road"`
	Pedestrian    string `json:"pedestrian"`
	Footway       string `json:"footway"`
	Cycleway      string `json:"cycleway"`
	Highway       string `json:"highway"`
	Path          string `json:"path"`
	Neighbourhood string `json:"neighbourhood"`
	Allotments    string `json:"allotments"`
	Quarter       string `json:"quarter"`
	CityDistrict  string `json:"city_district"`
	District      string `json:"district"`
	Borough       string `json:"borough"`
	Suburb        string `json:"suburb"`
	Subdivision   string `json:"subdivision"`
	Municipality  string `json:"municipality"`
	City          string `json:"city"`
	Town          string `json:"town"`
	Village       string `json:"village"`
	Hamlet        string `json:"hamlet"`
	County        string `json:"county"`
	Country       string `json:"country"`
	CountryCode   string `json:"country_code"`
	Region        string `json:"region"`
	State         string `json:"state"`
	StateDistrict string `json:"state_district"`
	Continent     string `json:"continent"`
	Postcode      string `json:"postcode"`
}

func (c *client) wait() {
	c.m.Lock()
	defer func() {
		c.lastRequest = time.Now()
		c.m.Unlock()
	}()

	if c.lastRequest.IsZero() {
		return
	}

	d := requestInterval - time.Since(c.lastRequest)
	if d < 0 {
		return
	}

	c.logger.Warn("Rate limited - waiting " + d.String())
	time.Sleep(d)
}

func SetClient(l *slog.Logger, ua string) {
	c = &client{
		url:       "https://nominatim.openstreetmap.org/",
		userAgent: ua,
		client:    http.Client{},
		logger:    l,
	}
}

func search(a string) ([]Result, error) {
	if c == nil {
		return nil, ErrClientNotSet
	}

	c.wait()

	q := struct {
		Q              string `url:"q"`
		Format         string `url:"format"`
		AddressDetails int    `url:"addressdetails"`
	}{
		Q:              a,
		Format:         "json",
		AddressDetails: 1,
	}

	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, c.url+"search?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	r := []Result{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r, nil
}

func SearchLocations(a string) ([]Result, error) {
	return search(a)
}

func Find(a string) (*geo.Address, error) {
	r, err := search(a)
	if err != nil {
		return nil, err
	}

	for _, f := range r {
		if f.DisplayName == a {
			return f.ToAddress(), nil
		}
	}

	return nil, ErrAddressNotFound
}

func Search(a string) ([]string, error) {
	r, err := search(a)
	if err != nil {
		return nil, err
	}

	addresses := []string{}
	for _, e := range r {
		addresses = append(addresses, e.DisplayName)
	}

	return addresses, nil
}

func Reverse(q Query) (*geo.Address, error) {
	if c == nil {
		return nil, ErrClientNotSet
	}

	c.wait()

	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, c.url+"reverse?"+v.Encode(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.userAgent)

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	r := Result{}
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	return r.ToAddress(), nil
}

func (r Result) ToAddress() *geo.Address {
	return &geo.Address{
		FormattedAddress: r.DisplayName,
		HouseNumber:      r.Address.HouseNumber,
		Street:           r.Address.Street(),
		Postcode:         r.Address.Postcode,
		City:             r.Address.Locality(),
		Suburb:           r.Address.Suburb,
		State:            r.Address.State,
		Country:          r.Address.Country,
		CountryCode:      strings.ToUpper(r.Address.CountryCode),
	}
}

func (a Address) Locality() string {
	return cmp.Or(
		a.City, a.Town, a.Village, a.Hamlet,
	)
}

func (a Address) Street() string {
	return cmp.Or(
		a.Road, a.Pedestrian, a.Path, a.Cycleway, a.Footway, a.Highway,
	)
}
