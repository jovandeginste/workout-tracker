package hx

import (
	"fmt"
	"time"
)

// Reswap specifies how the response will be swapped in an HTMX request.
//
// Use the Reswap() directly to set this header in the response, or choose
// from the constants below for common swap styles.
//
// The constants correspond to the different swap styles:
//   - SwapInnerHtml: Replaces the inner HTML of the target element.
//   - SwapOuterHtml: Replaces the entire target element with the response.
//   - SwapBeforeBegin: Inserts the response before the target element.
//   - SwapAfterBegin: Inserts the response before the first child of the target element.
//   - SwapBeforeEnd: Inserts the response after the last child of the target element.
//   - SwapAfterEnd: Inserts the response after the target element.
//   - SwapDelete: Deletes the target element, regardless of the response.
//   - SwapNone: Does not append content from the response (out-of-band items will still be processed).
//
// For more details, see: https://htmx.org/attributes/hx-swap
//
// Simple usage:
//
//	hx.Response(w, hx.Reswap("innerHTML"))
//	// Sets HX-Reswap header to "innerHTML"
//
// Constant usage example (same as above):
//
//	hx.Response(w, hx.SwapInnerHtml)
//	// Also sets HX-Reswap header to "innerHTML"
//
// Constant usage example (with modifiers):
//
//	hx.Response(w, hx.SwapInnerHtml.Swap(1*time.Second).Settle(2*time.Second))
//	// Sets HX-Reswap header to "innerHTML swap:1s settle:2s"
type Reswap string

func (s Reswap) apply(o *HtmxResponse) { o.headers[HxReswap] = string(s) }

// Reswap constants
const (
	// SwapInnerHtml replace the inner HTML of the target element
	SwapInnerHtml Reswap = "innerHTML"
	// SwapOuterHtml replace the entire target element with the response
	SwapOuterHtml Reswap = "outerHTML"
	// SwapAfterBegin insert the response before the target element
	SwapBeforeBegin Reswap = "beforebegin"
	// SwapAfterBegin insert the response before the first child of the target element
	SwapAfterBegin Reswap = "afterbegin"
	// SwapBeforeEnd insert the response after the last child of the target element
	SwapBeforeEnd Reswap = "beforeend"
	// SwapAfterEnd insert the response after the target element
	SwapAfterEnd Reswap = "afterend"
	// SwapDelete deletes the target element regardless of the response
	SwapDelete Reswap = "delete"
	// SwapNone does not append content from response (out of band items will still be processed)
	SwapNone Reswap = "none"
)

// Transition (reswap header modifier) allows you to specify the use of the View Transition API when a swap occurs
//
// More details: https://htmx.org/attributes/hx-swap/#transition-transition
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.Transition())
//	// Sets HX-Reswap header to "innerHTML transition:true"
func (s Reswap) Transition() Reswap {
	return Reswap(string(s) + " transition:true")
}

// Swap (reswap header modifier) is used to set a time wait after receiving a response before swapping the content
//
// More details: https://htmx.org/attributes/hx-swap/#timing-swap-settle
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.Swap(1*time.Second))
//	// Sets HX-Reswap header to "innerHTML swap:1s"
func (s Reswap) Swap(dur time.Duration) Reswap {
	return Reswap(fmt.Sprintf("%s swap:%s", s, dur))
}

// Settle (reswap header modifier) is used to set a time to wait after swapping before triggering the settle step
//
// More details: https://htmx.org/attributes/hx-swap/#timing-swap-settle
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.Settle(1*time.Second))
//	// Sets HX-Reswap header to "innerHTML settle:1s"
func (s Reswap) Settle(dur time.Duration) Reswap {
	return Reswap(fmt.Sprintf("%s settle:%s", s, dur))
}

// IgnoreTitle (reswap header modifier) is used to ignore any <title> tags in the response
//
// More details: https://htmx.org/attributes/hx-swap/#ignore-title
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.IgnoreTitle())
//	// Sets HX-Reswap header to "innerHTML ignoreTitle:true"
func (s Reswap) IgnoreTitle() Reswap {
	return Reswap(string(s) + " ignoreTitle:true")
}

// Scroll (reswap header modifier) is used to scroll to the "top", "bottom", a specific element after swapping
//
// More details: https://htmx.org/attributes/hx-swap/#scrolling-scroll-show
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.Scroll("top"))
//	// Sets HX-Reswap header to "innerHTML scroll:top"
//	hx.Response(w, hx.SwapInnerHtml.Scroll("#another-div:top"))
//	// Sets HX-Reswap header to "innerHTML scroll:#another-div:top"
func (s Reswap) Scroll(target string) Reswap {
	return Reswap(fmt.Sprintf("%s scroll:%s", s, target))
}

// Show (reswap header modifier) is used to show the "top", "bottom", or a specific element after swapping
//
// More details: https://htmx.org/attributes/hx-swap/#scrolling-scroll-show
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.Show("top"))
//	// Sets HX-Reswap header to "innerHTML show:top"
//	hx.Response(w, hx.SwapInnerHtml.Show("#another-div:top"))
//	// Sets HX-Reswap header to "innerHTML show:#another-div:top"
func (s Reswap) Show(target string) Reswap {
	return Reswap(fmt.Sprintf("%s show:%s", s, target))
}

// FocusScroll (reswap header modifier) can be used to enable to disable scrolling to the element after swapping
//
// More details: https://htmx.org/attributes/hx-swap/#focus-scroll
//
// Example usage:
//
//	hx.Response(w, hx.SwapInnerHtml.FocusScroll(true))
//	// Sets HX-Reswap header to "innerHTML focus-scroll:true"
func (s Reswap) FocusScroll(focus bool) Reswap {
	return Reswap(fmt.Sprintf("%s focus-scroll:%t", s, focus))
}
