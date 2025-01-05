# ctxi18n

[![Lint](https://github.com/invopop/ctxi18n/actions/workflows/lint.yaml/badge.svg)](https://github.com/invopop/ctxi18n/actions/workflows/lint.yaml)
[![Test Go](https://github.com/invopop/ctxi18n/actions/workflows/test.yaml/badge.svg)](https://github.com/invopop/ctxi18n/actions/workflows/test.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/invopop/ctxi18n)](https://goreportcard.com/report/github.com/invopop/ctxi18n)
[![codecov](https://codecov.io/gh/invopop/ctxi18n/graph/badge.svg?token=17H4KY3CNZ)](https://codecov.io/gh/invopop/ctxi18n)
[![GoDoc](https://godoc.org/github.com/invopop/ctxi18n?status.svg)](https://godoc.org/github.com/invopop/ctxi18n)
![Latest Tag](https://img.shields.io/github/v/tag/invopop/ctxi18n)

Go Context Internationalization - translating apps easily.

## Introduction

`ctxi18n` is heavily influenced by [internationalization in Ruby on Rails](https://guides.rubyonrails.org/i18n.html) and aims to make it just as straightforward in Go applications.

As the name suggests, `ctxi18n` focusses on making i18n data available inside an application's context instances, but is sufficiently flexible to used directly if needed.

Key Features:

- Loads locale files written in YAML or JSON with a similar structure those in Ruby i18n.
- Makes it easy to add a locale object to the context.
- Supports `fs.FS` to load data.
- Short method names like `i18n.T()` or `i18n.N()`.
- Support for simple interpolation using keys, e.g. `Some %{key} text`
- Support for pluralization rules.
- Default values when translations are missing.

## Usage

Import the library with:

```go
import "github.com/invopop/ctxi18n"
```

First you'll need to load YAML or JSON translation definitions. Files may be named and structured however you like, but the contents must always follow the same pattern of language and properties, for example:

```yaml
en:
  welcome:
    title: "Welcome to our application!"
    login: "Log in"
    signup: "Sign up"
    forgot-password: "Forgot Password?"
es:
  welcome:
    title: "¡Bienvenido a nuestra aplicación!"
    login: "Ingresarse"
    signup: "Registrarse"
    forgot-password: "¿Olvidaste tu contraseña?
```

The first level of properties of the object **must** always define the locale that the rest of sub-object's contents will provide translations for.

Files will all be deep-merged on top of each other so you can safely extend dictionaries from multiple sources.

To load the dictionary run something like the following where the `asset.Content` is a package containing [embedded files](https://pkg.go.dev/embed):

```go
if err := ctxi18n.Load(assets.Content); err != nil {
    panic(err)
}
```

If you'd like to set a default base language to try to use for any missing translations, load the assets with a default:

```go
if err := ctxi18n.LoadWithDefault(assets.Content, "en"); err != nil {
    panic(err)
}
```

You'll now have a global set of locales prepared in memory and ready to use. Assuming your application uses some kind of context such as from an HTTP or gRPC request, you'll want to add a single locale to it:

```go
ctx = ctxi18n.WithLocale(ctx, "en")
```

Locale selection is performed according to [RFC9110](https://www.rfc-editor.org/rfc/rfc9110.html) and the `Accept-Language` header, so you can pass in a code string and an attempt will be made to find the best match:

```go
ctx = ctxi18n.WithLocale(ctx, "en-US,en;q=0.9,es;q=0.8")
```

In this example, the first locale to matched will be `en-US`, followed by just `en`, then `es`:

Getting translations is straightforward, you have two options:

1.  call methods defined in the package with the context, or,
2.  extract the locale from the context and use.

To translate without extracting the locale, you'll need to load the `i18n` package which contains all the structures and methods used by the main `ctxi18n` without any globals:

```go
import "github.com/invopop/ctxi18n/i18n"
```

Then use it with the context:

```go
fmt.Println(i18n.T(ctx, "welcome.title"))
```

Notice in the example that `title` was previously defined inside the `welcome` object in the source YAML, and we're accessing it here by defining the path `welcome.title`.

To use the `Locale` object directly, extract it from the context and call the methods:

```go
l := ctxi18n.Locale(ctx)
fmt.Println(l.T("welcome.title"))
```

There is no preferred way on how to use this library, so please use whatever best first your application and coding style. Sometimes it makes sense to pass in the context in every call, other times the code can be shorter and more concise by extracting it.

### Defaults

If a translation is missing from the locale a "missing" text will be produced, for example:

```go
fmt.Println(l.T("welcome.no.text"))
```

Will return a text that follows the `fmt.Sprintf` missing convention:

```
!(MISSING welcome.no.text)
```

This can be useful for translators to figure out which texts are missing, but sometimes a default value is more appropriate:

```go
fmt.Println(i18n.T(ctx, "welcome.question", i18n.Default("Just ask!")))
// output: "Just ask!"
code := "EUR"
fmt.Println(i18n.T(ctx, "currencies."+code, i18n.Default(code)))
// output: "EUR"
```

An alternative to using defaults is to check if the key exists using the `Has` method:

```go
if !i18n.Has(ctx, "welcome.question") {
  fmt.Println("Just ask!")
}
```

### Interpolation

Go's default approach for interpolation using the `fmt.Sprintf` and related methods is good for simple use-cases. For example, given the following translation:

```yaml
en:
  welcome:
    title: "Hi %s, welcome to our App!"
```

You can get the translated text and interpolate with:

```go
i18n.T(ctx, "welcome.title", "Sam")
```

This however is an _anti-pattern_ when it comes to translating applications as translators may need to change the order of replaced words. To get around this, `ctxi18n` supports simple named interpolation as follows:

```yaml
en:
  welcome:
    title: "Hi %{name}, welcome to our App!"
```

```go
i18n.T(ctx, "welcome.title", i18n.M{"name":"Sam"})
```

The `i18n.M` map is used to perform a simple find and replace on the matching translation. The `fmt.Sprint` method is used to convert values into strings, so you don't need to worry about simple serialization like for numbers.

Interpolation can also be used alongside default values:

```go
i18n.T(ctx, "welcome.title", i18n.Default("Hi %{name}"), i18n.M{"name":"Sam"})
```

## Pluralization

When texts include references to numbers we need internationalization libraries like `ctxi18n` that help define multiple possible translations according to a number. Pluralized translations are defined like this:

```yaml
en:
  inbox:
    emails:
      zero: "You have no emails."
      one: "You have %{count} email."
      other: "You have %{count} emails.
```

The `inbox.emails` tag has a sub-object that defines all the translations we need according to the pluralization rules of the language. In the case of English which uses the default rule set, `zero` is an optional definition that will be used if provided and fallback on `other` if not.

To use these translations, call the `i18n.N` method:

```go
count := 2
fmt.Println(i18n.N(ctx, "inbox.emails", count, i18n.M{"count": count}))
```

The output from this will be: "You have 2 emails."

In the current implementation of `ctxi18n` there are very few pluralization rules defined, please submit PRs if your language is not covered!

## Scopes

As your application gets more complex, it can get repetitive having to use the same base keys. To get around this, use the `WithScope` helper method inside a context:

```go
ctx := i18n.WithScope(ctx, "welcome")
i18n.T(ctx, ".title", i18n.M{"name":"Sam"})
```

Anything with the `.` at the beginning will append the scope. You can continue to use any other key in the locale by not using the `.` at the front.

## Templ

[Templ](https://templ.guide/) is a templating library that helps you create components that render fragments of HTML and compose them to create screens, pages, documents or apps.

The following "Hello World" example is taken from the [Templ Guide](https://templ.guide) and shows how you can quickly add translations the leverage the built-in `ctx` variable provided by Templ.

```yaml
en:
  welcome:
    hello: "Hello, %{name}"
```

```go
package main

import "github.com/invopop/ctxi18n/i18n"

templ Hello(name string) {
  <span class="hello">
    { i18n.T(ctx, "welcome.hello", i18n.M{"name": name}) }
  </span>
}

templ Greeting(person Person) {
  <div class="greeting">
    @Hello(person.Name)
  </div>
}
```

To save even more typing, it might be worth defining your own templ wrappers around those defined in the `i18n` package. Check out the [gobl.html `t` package](https://github.com/invopop/gobl.html/tree/main/components/t) for an example.

# Examples

The following is a list of Open Source projects using this library from which you can see working examples for your own solutions. Please send in a PR if you'd like to add your project!

- [GOBL HTML](https://github.com/invopop/gobl.html) - generate HTML files like invoices from [GOBL](https://gobl.org) documents.
