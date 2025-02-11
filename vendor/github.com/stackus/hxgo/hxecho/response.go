package hxecho

import (
	"github.com/labstack/echo/v4"

	"github.com/stackus/hxgo"
)

// Response modifies the echo.Context to add HTMX headers and status codes.
//
// This will set the HTMX headers but will not set the Status Code. Use the
// returned response to set the Status Code later with `response.StatusCode()`.
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
func Response(ctx echo.Context, options ...hx.ResponseOption) (*hx.HtmxResponse, error) {
	r, err := hx.BuildResponse(options...)
	if err != nil {
		return nil, err
	}

	for k, v := range r.Headers() {
		ctx.Response().Header().Set(k, v)
	}

	// Skip setting the Status Code for Echo to avoid superfluous write errors

	return r, nil
}
