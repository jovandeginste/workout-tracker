![golangish-htmxish-logo.png](golangish-htmxish-logo.png)

# HTMX library for Go
[![GoDoc](https://godoc.org/github.com/stackus/hxgo?status.svg)](https://godoc.org/github.com/stackus/hxgo)

This comprehensive library offers an array of functions and types specifically designed to streamline the handling of [HTMX](https://htmx.org/) requests and the construction of responses in the Go applications.

> README.md logo image courtesy of ChatGPT.

## Features
- Request and Response header helpers
- Easy APIs to build complex HTMX responses for Locations, Reswaps, and Triggers

```go
import (
    "net/http"
    
    "github.com/stackus/hxgo"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
    if hx.IsHtmx(r) {
        // do something
		
        // load up on HTMX headers and set the status code to send back to the client
        err := hx.Response(w,
            hx.Location("/new-location",
                hx.Target("#my-target"),
                hx.Swap(hx.SwapInnerHtml.IgnoreTitle()),
                hx.Values(map[string]string{"key": "value"}),
            ),
            hx.StatusStopPolling,
            hx.Trigger(
                hx.Event("my-event"),
                hx.Event("my-other-event", "my-other-event-value"),
                hx.Event("my-complex-event", map[string]any{
                    "foo": "bar",
                    "baz": 123,
                }
            ),
        )
        if err != nil {
            // handle error
        }
    }
}
```

## Installation
The minimum version of Go required is **1.18**. Generics have been used to make some types and options easier to work with.

Install using `go get`:
```bash
go get github.com/stackus/hxgo
```

Then import the package into your project:
```go
import "github.com/stackus/hxgo"
```

You'll then use `hx.*` to access the functions and types.

## Working with Requests
To determine if a request is an HTMX request, use the `IsHtmx` function:

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    if hx.IsHtmx(r) {
        // do something
    }
}
```

Helpers exist for each of the [HTMX request headers](https://htmx.org/reference/#request_headers):

- `HX-Boosted`: Use the `IsBoosted` function to determine if the request is a boosted request
- `HX-Current-URL`: Use the `GetCurrentUrl` function to get the current URL of the request
- `HX-History-Restore-Request`: Use the `IsHistoryRestoreRequest` function to determine if the request is a history restore request
- `HX-Prompt`: Use the `GetPrompt` function to get the prompt value of the request
- `HX-Request`: Use the `IsRequest` or `IsHTMX` functions to determine if the request is an HTMX request
- `HX-Target`: Use the `GetTarget` function to get the target value of the request
- `HX-Trigger-Name`: Use the `GetTriggerName` function to get the trigger name of the request
- `HX-Trigger`: Use the `GetTrigger` function to get the trigger value of the request

`Is*` functions return a boolean while `Get*` functions return a string. The absence of the corresponding HTMX header will return false or an empty string respectively.

## Working with Responses
Use the `Response` function to modify the `http.ResponseWriter` to return an HTMX response:

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    err := hx.Response(w, hx.Retarget("/new-location"))
	if err != nil {
        // handle error
    }
}
```

Each of the [HTMX response headers](https://htmx.org/reference/#response_headers) has a corresponding option to set the header:

- `HX-Location`: Use the `Location` option with a variable number of properties to set the location header. See the [Location](#location) section for more details.
- `HX-Push-Url`: Use the `PushURL` option to push a new URL into the browser history
- `HX-Redirect`: Use the `Redirect` option to redirect the browser to a new URL
- `HX-Refresh`: Use the `Refresh` option to refresh the browser
- `HX-Replace-Url`: Use the `ReplaceUrl` option to replace the current URL in the browser history
- `HX-Reswap`: Use the `Reswap` option or one of the `Swap*` constants to specify how the response will be swapped. See the [Reswap](#reswap) section for more details.
- `HX-Retarget`: Use the `Retarget` option with a CSS selector to redirect the response to a new element
- `HX-Reselect`: Use the `Reselect` option with a CSS selector to designate a different element in the response to be used
- `HX-Trigger`: Use the `Trigger` option to trigger client-side events. See the [Trigger](#trigger) section for more details.
- `HX-Trigger-After-Settle`: Use the `TriggerAfterSettle` option to trigger client-side events after the response has settled. See the [Trigger](#trigger) section for more details.
- `HX-Trigger-After-Swap`: Use the `TriggerAfterSwap` option to trigger client-side events after the response has been swapped. See the [Trigger](#trigger) section for more details.

### Location
The `Location` option is used to set the [HX-Location Response Header](https://htmx.org/headers/hx-location/). It takes a path string and then an optional number of properties. The following properties are supported:

- `Source`: The `Source` property is used to set the source element of the location header.
- `Event`: The `EventName` property is used to set the name of the event of the location header.  
  > Note: This property is called `EventName` so that it does not conflict with the `Event` property used by the `Trigger` option.
- `Handler`: The `Handler` property is used to set the handler of the location header.
- `Target`: The `Target` property is used to set the target of the location header.
- `Swap`: The `Swap` property is used to set the swap of the location header. The value may be a string or any of the `Swap*` constants.
- `Values`: The `Values` property is used to set the values of the location header. The value may be anything, but it is recommended to use a `map[string]any` or struct with JSON tags.
- `Headers`: The `Headers` property is used to set the headers of the location header. The value needs to be a `map[string]string`.
- `Select`: The `Select` property is used to set the select of the location header.

Setting just the path:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Location("/new-location"))
	  // Hx-Location: /new-location
}
```
Setting multiple properties:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Location("/new-location",
        hx.Target("#my-target"),
        hx.Swap(hx.SwapInnerHtml.IgnoreTitle()),
        hx.Values(map[string]string{"key": "value"}),
    ))
    // Hx-Location: {"path":"/new-location","target":"#my-target","swap":"innerHTML ignoreTitle:true","values":{"key":"value"}}
}
```

### Reswap
The `Reswap` option is used to set the HX-Reswap response header. Using the `Reswap` option directly is possible, but it is recommended to use one of the `Swap*` constants instead. The following constants are supported:

- `SwapInnerHtml`: Sets the HX-Reswap response header to `innerHTML`
- `SwapOuterHtml`: Sets the HX-Reswap response header to `outerHTML`
- `SwapBeforeBegin`: Sets the HX-Reswap response header to `beforebegin`
- `SwapAfterBegin`: Sets the HX-Reswap response header to `afterbegin`
- `SwapBeforeEnd`: Sets the HX-Reswap response header to `beforeend`
- `SwapAfterEnd`: Sets the HX-Reswap response header to `afterend`
- `SwapDelete`: Sets the HX-Reswap response header to `delete`
- `SwapNone`: Sets the HX-Reswap response header to `none`

The result from `Reswap` and each constant can be chained with modifiers to configure the header even further. The following modifiers are supported:

- `Transition`: Adds `transition:true` to enable the use of the View Transition API
- `Swap`: Used with a `time.Duration` to set the swap delay
- `Settle`: Used with a `time.Duration` to set the settle delay
- `IgnoreTitle`: Adds `ignoreTitle:true` to ignore the title of the response
- `Scroll`: Used with a CSS selector to scroll to the element after swapping
- `Show`: Used with a CSS selector to show the element after swapping
- `FocusScroll`: Used with a boolean to set the focus scroll behavior

Setting just the reswap header two ways:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Reswap("innerHTML"))
    // Hx-Reswap: innerHTML
    hx.Response(w, hx.SwapInnerHtml)
    // Hx-Reswap: innerHTML
}
```

Setting the reswap header with modifiers:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.SwapInnerHtml.IgnoreTitle().Transition())
    // Hx-Reswap: innerHTML ignoreTitle:true transition:true
}
```

### Trigger
The `Trigger` option is used to set the [HX-Trigger Response Header](https://htmx.org/headers/hx-trigger/). It takes a variable number of events to trigger on the client.

Events are created using `hx.Event` and can be either simple names or complex objects. The supported events include:

Setting a simple event:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Trigger(hx.Event("my-event")))
    // Hx-Trigger: {"my-event":null}
}
```

Setting a complex event:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    myEvent := map[string]any{
        "foo": "bar",
        "baz": 123,
    }

    hx.Response(w, hx.Trigger(hx.Event("my-event", myEvent)))
	  // Hx-Trigger: {"my-event":{"foo":"bar","baz":123}}
}
```

Setting multiple events:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Trigger(
        hx.Event("my-event"),
        hx.Event("my-other-event", "my-other-event-value"),
    ))
    // Hx-Trigger: {"my-event":null,"my-other-event":"my-other-event-value"}
}
```

The `data`, which is the second parameter of the `Event`, is variadic. If more than one data value is passed, the event is set to an array of those values. The following events demonstrate this equivalence:

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Trigger(
        hx.Event("my-event-1", "foo", "bar"),
        hx.Event("my-event-2", []string{"foo", "bar"}),
    ))
    // Hx-Trigger: {"my-event-1":["foo","bar"], "my-event-2":["foo","bar"]}
}
```

Both `TriggerAfterSettle` and `TriggerAfterSwap` are available to trigger events after the response has settled or been swapped respectively. They take the same event arguments as `Trigger`.

### Status
The `Status` option is used to set the HTTP status code of the response. There is only one status constant available:

- `StatusStopPolling`: Sets the HTTP status code to 286 which is used by HTMX to halt polling requests

Setting the status code:
```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.StatusStopPolling)
    // HTTP/1.1 286
}
```

The `Status` option can be used to set any HTTP status code and is not limited to the constants provided by this library.

```go
func MyHandler(w http.ResponseWriter, r *http.Request) {
    hx.Response(w, hx.Status(http.StatusGone))
    // HTTP/1.1 410
}
```

## Usage with different HTTP frameworks
With the standard library, and other frameworks that adhere to its `http.ResponseWriter` interface, the `Response` function can be used directly to modify the response.

### Standard Library (and some like Chi)
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stackus/hxgo"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	// Add HTMX headers and a status code to the response
	err := hx.Response(w,
		hx.Location("/foo"),
		hx.StatusStopPolling,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Write the response body
	_, _ = fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	fmt.Println("Server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
```

### Echo, Fiber, and other frameworks
For frameworks that do not use the standard library's `http.ResponseWriter` interface, there are request and response helpers available to make it easier to work with HTMX.

For example, with [Echo](https://echo.labstack.com/):

```go
package main

import (
	"github.com/labstack/echo/v4"

	"github.com/stackus/hxgo"
	"github.com/stackus/hxgo/hxecho"
)

func main() {
	// Create a new instance of Echo
	e := echo.New()

	// Define a route for "/"
	e.GET("/", func(c echo.Context) error {
		// use hxecho.IsHtmx to determine if the request is an HTMX request
		if hxecho.IsHtmx(c) {
			// do something
			// Adds HTMX headers but does not set the Status Code
			r, err := hxecho.Response(c,
				// Continue to use the base htmx types and options
				hx.Location("/foo"),
				hx.StatusStopPolling,
			)
			if err != nil {
				return err
			}

			// Set the HTMX status code here and response body
			return c.String(r.StatusCode(), "Hello Echo")
		}
	})

	// Start the server on port 8080
	e.Logger.Fatal(e.Start(":8080"))
}
```

You will find request and response helpers for the following frameworks:
- Echo: [hxecho](./hxecho)
- Fiber: [hxfiber](./hxfiber)
- Gin: [hxgin](./hxgin)

The `Response` function for each library will return a default status of 200 if no status is set.
If you need to set a status code, you can use the `Status` option.

### Contributions
Contributions are welcome! Please open an issue or submit a pull request. If at all possible, please provide an example with your bug reports and tests with your pull requests.

#### Reporting Bugs
- If you find a bug, please open an issue.
- Include a clear description of the bug, steps to reproduce it, and any relevant logs or screenshots.
- Before creating a new issue, please check if it has already been reported to avoid duplicates.

#### Suggesting Enhancements
- We're always looking to improve our library. If you have ideas for new features or enhancements, feel free to open an issue to discuss it.
- Clearly explain your suggestion and its potential benefits.

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
