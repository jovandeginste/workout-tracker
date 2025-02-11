package hx

// Status is used to set the HTTP status code of the HTMX response.
//
// Example usage:
//
//	hx.Response(w, hx.Status(http.StatusGone))
//	// Sets the HTTP status code to 410.
type Status int

func (s Status) apply(o *HtmxResponse) { o.status = int(s) }

// HTMX status codes.
const (
	// StatusStopPolling sends HTTP status code 286 to the client to stop polling.
	//
	// Example usage:
	//  hx.Response(w, hx.StatusStopPolling)
	//  // Sets the HTTP status code to 286.
	StatusStopPolling Status = 286
)

// PushUrl sets the HX-Push-Url header.
//
// It pushes a new URL into the history stack.
//
// Example usage:
//
//	hx.Response(w, hx.PushUrl("/new-url-location"))
//	// Sets the HX-Push-Url header to "/new-url-location".
type PushUrl string

func (p PushUrl) apply(o *HtmxResponse) { o.headers[HxPushUrl] = string(p) }

// Redirect sets the HX-Redirect header.
//
// It is used for client-side redirects that require a full page reload.
//
// Example usage:
//
//	hx.Response(w, hx.Redirect("/new-url-location"))
//	// Sets the HX-Redirect header to "/new-url-location".
type Redirect string

func (r Redirect) apply(o *HtmxResponse) { o.headers[HxRedirect] = string(r) }

// Refresh sets the HX-Refresh header.
//
// When set to "true", it triggers a full refresh of the client-side page.
// Note: This function always sets it to "true".
//
// Example usage:
//
//	hx.Response(w, hx.Refresh())
//	// Sets the HX-Refresh header to "true".
func Refresh() responseOptionFunc {
	return func(o *HtmxResponse) {
		o.headers[HxRefresh] = "true"
	}
}

// ReplaceUrl sets the HX-Replace-Url header.
//
// It replaces the current URL in the location bar.
//
// Example usage:
//
//	hx.Response(w, hx.ReplaceUrl("/new-url-location"))
//	// Sets the HX-Replace-Url header to "/new-url-location".
type ReplaceUrl string

func (r ReplaceUrl) apply(o *HtmxResponse) { o.headers[HxReplaceUrl] = string(r) }

// Retarget sets the HX-Retarget header.
//
// This option specifies a new CSS selector to redirect the content update to a different element on the page.
//
// Example usage:
//
//	hx.Response(w, hx.Retarget("#new-target"))
//	// Sets the HX-Retarget header to "#new-target".
type Retarget string

func (t Retarget) apply(o *HtmxResponse) { o.headers[HxRetarget] = string(t) }

// Reselect sets the HX-Reselect header.
//
// This option designates a CSS selector to determine which part of the response should be used for swapping in, effectively overriding any existing hx-select on the triggering element.
//
// Example usage:
//
//	hx.Response(w, hx.Reselect("#new-target"))
//	// Sets the HX-Reselect header to "#new-target".
type Reselect string

func (s Reselect) apply(o *HtmxResponse) { o.headers[HxReselect] = string(s) }
