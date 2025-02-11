package hx

import (
	"net/http"
)

// Response Headers
const (
	// HxLocation allows for client-side redirects without a full page reload.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the Location() option to set this header in the response.
	HxLocation = "Hx-Location"

	// HxPushUrl pushes a new URL into the history stack.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the PushUrl() option to set this header in the response.
	HxPushUrl = "Hx-Push-Url"

	// HxRedirect can be used for client-side redirects that require a full page reload.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the Redirect() option to set this header in the response.
	HxRedirect = "Hx-Redirect"

	// HxRefresh when set to "true", triggers a full refresh of the client-side page.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the Refresh() option to set this header in the response.
	HxRefresh = "Hx-Refresh"

	// HxReplaceUrl replaces the current URL in the location bar.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the ReplaceUrl() option to set this header in the response.
	HxReplaceUrl = "Hx-Replace-Url"

	// HxReswap specifies how the response will be swapped.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the Reswap() option to set this header in the response.
	HxReswap = "Hx-Reswap"

	// HxRetarget updates the target of the content update to a different element on the page using a CSS selector.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the Retarget() option to set this header in the response.
	HxRetarget = "Hx-Retarget"

	// HxReselect allows selection of a specific part of the response to be swapped in, using a CSS selector. It overrides any existing hx-select on the triggering element.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the Reselect() option to set this header.
	HxReselect = "Hx-Reselect"

	// HxTriggerAfterSettle triggers client-side events after the settle step.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the TriggerAfterSettle() option to set this header in the response.
	HxTriggerAfterSettle = "Hx-Trigger-After-Settle"

	// HxTriggerAfterSwap triggers client-side events after the swap step.
	//
	// See https://htmx.org/reference/#response_headers for more details.
	//
	// Use the TriggerAfterSwap() option to set this header in the response.
	HxTriggerAfterSwap = "Hx-Trigger-After-Reswap"
)

// Response modifies the http.ResponseWriter to add HTMX headers and status codes.
//
// The following options are available:
//   - Status(int) | StatusStopPolling: Sets the HTTP status code of the HTMX response.
//   - Location(path, ...properties): Enables client-side redirection without a full page reload.
//   - PushUrl(string): Pushes a new URL into the history stack.
//   - Redirect(string): Performs a client-side redirect with a full page reload.
//   - Refresh(bool): If set to "true", triggers a full refresh of the client-side page.
//   - ReplaceUrl(string): Replaces the current URL in the location bar.
//   - Reswap(string) | {Swap constants}: Specifies how the response will be swapped.
//   - Retarget(string): A CSS selector to update the target of the content update to a different page element.
//   - Reselect(string): A CSS selector to select a part of the response to be swapped in, overriding existing hx-select on the triggering element.
//   - Trigger(...events): Triggers client-side events.
//   - TriggerAfterSettle(...events): Triggers client-side events after the settle step.
//   - TriggerAfterSwap(...events): Triggers client-side events after the swap step.
func Response(w http.ResponseWriter, options ...ResponseOption) error {
	o, err := BuildResponse(options...)
	if err != nil {
		return err
	}

	if len(o.headers) > 0 {
		for k, v := range o.headers {
			w.Header().Set(k, v)
		}
	}

	// Support setting the stop polling status code.
	if o.status != 0 {
		w.WriteHeader(o.status)
	}

	return nil
}

// BuildResponse creates a new HtmxResponse from the provided options.
//
// It can be used to create a response helper for your own HTTP library.
//
// Several libraries have already been implemented:
//   - Echo: import github.com/stackus/hxgo/hxecho
//   - Fiber: import github.com/stackus/hxgo/hxfiber
//   - Gin: import github.com/stackus/hxgo/hxgin
func BuildResponse(options ...ResponseOption) (response *HtmxResponse, err error) {
	// Recover from any panics that might happen during the processing of the options.
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	o := &HtmxResponse{
		headers: make(map[string]string),
	}
	for _, option := range options {
		option.apply(o)
	}

	return o, nil
}
