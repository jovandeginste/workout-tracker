package hx

import (
	"net/http"
)

// Request Headers
const (
	// HxBoosted indicates that the request is a boosted request.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the IsBoosted() function to check this header in the request.
	HxBoosted = "Hx-Boosted"

	// HxCurrentUrl represents the current URL of the browser.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the GetCurrentUrl() function to fetch this header from the request.
	HxCurrentUrl = "Hx-Current-Url"

	// HxHistoryRestoreRequest is "true" if the request is for history restoration after a miss in the local history cache.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the IsHistoryRestoreRequest() function to check this header in the request.
	HxHistoryRestoreRequest = "Hx-History-Restore-Request"

	// HxPrompt captures the user's response to an HX-Prompt.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the GetPrompt() function to fetch this header from the request.
	HxPrompt = "Hx-Prompt"

	// HxRequest is always "true" if the request is an HTMX request.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the IsRequest() function to check this header in the request.
	HxRequest = "Hx-Request"

	// HxTarget identifies the ID of the target element, if it exists.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the GetTarget() function to fetch this header from the request.
	HxTarget = "Hx-Target"

	// HxTriggerName denotes the name of the triggered element, if it exists.
	//
	// See https://htmx.org/reference/#request_headers for more details.
	//
	// Use the GetTriggerName() function to fetch this header from the request.
	HxTriggerName = "Hx-Trigger-Name"
)

// IsBoosted checks the HX-Boosted header
//
// Returns true if the request is a boosted request
func IsBoosted(r *http.Request) bool {
	return r.Header.Get(HxBoosted) != ""
}

// GetCurrentUrl extracts the HX-Current-URL header from an HTTP request.
//
// It returns the current URL of the browser if the header exists.
// If the header is not present, it returns an empty string.
func GetCurrentUrl(r *http.Request) string {
	return r.Header.Get(HxCurrentUrl)
}

// IsHistoryRestoreRequest determines if an HTTP request is a history restore request.
//
// It checks the presence of the HX-History-Restore-Request header in the request.
// Returns true if the header is present, otherwise returns false.
func IsHistoryRestoreRequest(r *http.Request) bool {
	return r.Header.Get(HxHistoryRestoreRequest) != ""
}

// GetPrompt extracts the HX-Prompt header from an HTTP request.
//
// It returns the user response to an hx-prompt if the header exists.
// If the header is not present, it returns an empty string.
func GetPrompt(r *http.Request) string {
	return r.Header.Get(HxPrompt)
}

// IsRequest determines if an HTTP request is an HTMX request.
//
// It checks the presence of the HX-Request header in the request.
// Returns true if the header is present, otherwise returns false.
func IsRequest(r *http.Request) bool {
	return r.Header.Get(HxRequest) != ""
}

// IsHtmx determines if an HTTP request is an HTMX request.
//
// Does the same thing as IsRequest, only with a more user-friendly name.
func IsHtmx(r *http.Request) bool {
	return IsRequest(r)
}

// GetTarget extracts the HX-Target header from an HTTP request.
//
// It returns the ID of the target element if the header exists.
// If the header is not present, it returns an empty string.
func GetTarget(r *http.Request) string {
	return r.Header.Get(HxTarget)
}

// GetTriggerName extracts the HX-Trigger-Name header from an HTTP request.
//
// It returns the name of the triggered element if the header exists.
// If the header is not present, it returns an empty string.
func GetTriggerName(r *http.Request) string {
	return r.Header.Get(HxTriggerName)
}

// GetTrigger extracts the HX-Trigger header from an HTTP request.
//
// It returns the ID of the trigger element if the header exists.
// If the header is not present, it returns an empty string.
func GetTrigger(r *http.Request) string {
	return r.Header.Get(HxTrigger)
}
