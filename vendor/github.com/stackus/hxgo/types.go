package hx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// internal types related to Location

type property interface {
	apply(*location)
}

type propertyFunc func(*location)

func (f propertyFunc) apply(o *location) { f(o) }

type location struct {
	Path    string            `json:"path"`
	Source  string            `json:"source,omitempty"`
	Event   string            `json:"event,omitempty"`
	Handler string            `json:"handler,omitempty"`
	Target  string            `json:"target,omitempty"`
	Swap    string            `json:"swap,omitempty"`
	Values  any               `json:"values,omitempty"`
	Headers map[string]string `json:"headers,omitempty"`
	Select  string            `json:"select,omitempty"`
}

// TODO Can you return a complete event in the HX-Location event property? If so uncomment this
// func (e event) apply(o *location) {
// 	data, err := json.Marshal(e)
// 	if err != nil {
// 		panic(fmt.Errorf("unable to marshal event: %w", err))
// 	}
// 	o.Event = string(data)
// }

// internal types related to the response

// HtmxResponse is a struct that contains the headers and status code to be returned to the client
//
// This is helpful for using HTMX with a framework that doesn't implement the stdlib http.ResponseWriter
type HtmxResponse struct {
	headers map[string]string
	status  int
}

func (r HtmxResponse) Headers() map[string]string { return r.headers }
func (r HtmxResponse) StatusCode() int {
	if r.status == 0 {
		return http.StatusOK
	}
	return r.status
}

// ResponseOption is an interface that can be used to set the headers and status code of the response
type ResponseOption interface {
	apply(*HtmxResponse)
}

type responseOptionFunc func(*HtmxResponse)

func (f responseOptionFunc) apply(o *HtmxResponse) { f(o) }

// internal types related to triggering events

type event func() map[string]any

func triggeredEvents(events []event) []byte {
	m := make(map[string]any)
	for _, event := range events {
		for k, v := range event() {
			m[k] = v
		}
	}
	data, err := json.Marshal(m)
	if err != nil {
		panic(fmt.Errorf("unable to marshal all events: %w", err))
	}

	return data
}
