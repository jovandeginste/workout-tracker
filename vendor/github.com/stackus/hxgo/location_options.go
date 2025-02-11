package hx

import (
	"encoding/json"
	"fmt"
)

// Location sets the HX-Location header.
//
// This header is used for client-side redirection without a full page reload.
// Additional properties like Target can be specified for more complex behaviors.
//
// For more details, see: https://htmx.org/headers/hx-location
//
// Simple redirection example:
//
//	hx.Response(w, hx.Location("/test"))
//	// Sets HX-Location header to "/test"
//
// Redirection with additional target example:
//
//	hx.Response(w, hx.Location("/test",
//	  hx.Target("#testdiv"),
//	))
//	// Sets HX-Location header to a JSON object: {"path":"/test","target":"#testdiv"}
func Location(path string, properties ...property) responseOptionFunc {
	return func(o *HtmxResponse) {
		loc := location{
			Path: path,
		}

		if len(properties) == 0 {
			o.headers[HxLocation] = loc.Path
			return
		}

		for _, property := range properties {
			property.apply(&loc)
		}

		value, err := json.Marshal(loc)
		if err != nil {
			panic(fmt.Errorf("unable to marshal HX-Location header: %w", err))
		}

		o.headers[HxLocation] = string(value)
	}
}

// Source sets the 'source' property of the HX-Location header.
//
// More details: https://htmx.org/headers/hx-location
type Source string

func (l Source) apply(o *location) { o.Source = string(l) }

// EventName sets the 'event' property of the HX-Location header.
//
// More details: https://htmx.org/headers/hx-location
type EventName string

func (l EventName) apply(o *location) { o.Event = string(l) }

// Handler sets the 'handler' property of the HX-Location header.
//
// More details: https://htmx.org/headers/hx-location
type Handler string

func (l Handler) apply(o *location) { o.Handler = string(l) }

// Target sets the 'target' property of the HX-Location header.
//
// More details: https://htmx.org/headers/hx-location
type Target string

func (l Target) apply(o *location) { o.Target = string(l) }

// Swap sets the 'swap' property of the HX-Location header.
//
// Either a string or a Reswap constant can be used.
//
// More details: https://htmx.org/headers/hx-location
func Swap[T string | Reswap](swap T) propertyFunc {
	return func(o *location) { o.Swap = string(swap) }
}

// Values sets the 'values' property of the HX-Location header.
//
// Accepts any type, but a map[string]any or a struct with JSON tags is recommended.
//
// More details: https://htmx.org/headers/hx-location
func Values(values any) propertyFunc {
	return func(o *location) { o.Values = values }
}

// Headers sets the 'headers' property of the HX-Location header.
//
// Accepts a map[string]string.
//
// More details: https://htmx.org/headers/hx-location
type Headers map[string]string

func (l Headers) apply(o *location) { o.Headers = l }

// Select sets the 'select' property of the HX-Location header.
//
// More details: https://htmx.org/headers/hx-location
type Select string

func (l Select) apply(o *location) { o.Select = string(l) }
