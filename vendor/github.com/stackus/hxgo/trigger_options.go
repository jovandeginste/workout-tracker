package hx

// Request & Response Headers
const (
	// HxTrigger
	// (in request) the id of the trigger element if it exists
	// (in response) allows you to trigger events on the client
	//
	// More details: https://htmx.org/reference/#request_headers and https://htmx.org/reference/#response_headers
	//
	//  - Use the GetTrigger() function to fetch this header from the request
	//  - Use the Trigger(...events) option to set this header on the response
	HxTrigger = "HX-Trigger"
)

// Trigger allows you to trigger events on the client
//
// More details: https://htmx.org/reference/#response_headers
//
// Use Event to create events to pass to this option.
//
// Simple example:
//
//	hx.Response(w, hx.Trigger(hx.Event("myEvent")))
//	// Sets HX-Trigger header to {"myEvent":null}
//
// Example with data:
//
//	hx.Response(w, hx.Trigger(hx.Event("myEvent", "myData")))
//	// Sets HX-Trigger header to {"myEvent":"myData"}
//
// Example with multiple events:
//
//	hx.Response(w, hx.Trigger(
//		hx.Event("myEvent", "myData"),
//		hx.Event("myOtherEvent", "myOtherData"),
//	))
//	// Sets HX-Trigger header to {"myEvent":"myData","myOtherEvent":"myOtherData"}
//
// See also: TriggerAfterSettle and TriggerAfterSwap
func Trigger(events ...event) responseOptionFunc {
	return func(o *HtmxResponse) {
		data := triggeredEvents(events)
		o.headers[HxTrigger] = string(data)
	}
}

// TriggerAfterSettle triggers client-side events after the settle step.
//
// More details: https://htmx.org/reference/#response_headers
//
// For more details, see: Trigger
func TriggerAfterSettle(events ...event) responseOptionFunc {
	return func(o *HtmxResponse) {
		data := triggeredEvents(events)
		o.headers[HxTriggerAfterSettle] = string(data)
	}
}

// TriggerAfterSwap triggers client-side events after the swap step.
//
// More details: https://htmx.org/reference/#response_headers
//
// For more details, see: Trigger
func TriggerAfterSwap(events ...event) responseOptionFunc {
	return func(o *HtmxResponse) {
		data := triggeredEvents(events)
		o.headers[HxTriggerAfterSwap] = string(data)
	}
}

// Event creates an event to pass to one of the Trigger options.
//
// Simple named event:
//
//	hx.Event("myEvent")
//	// Returns {"myEvent":null}
//
// Named event with string data:
//
//	hx.Event("myEvent", "myData")
//	// Returns {"myEvent":"myData"}
//
// Named event with multiple data items (treated like an array):
//
//	hx.Event("myEvent", "myData1", "myData2")
//	// Returns {"myEvent":["myData1","myData2"]}
//
// Named event with a struct or map (treated like an object):
//
//	hx.Event("myEvent", map[string]string{"myKey": "myValue"})
//	// Returns {"myEvent":{"myKey":"myValue"}}
func Event(name string, data ...any) event {
	return func() map[string]any {
		switch len(data) {
		case 0:
			return map[string]any{name: nil}
		case 1:
			return map[string]any{name: data[0]}
		default:
			return map[string]any{name: data}
		}
	}
}
