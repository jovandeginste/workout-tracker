[![Sourcegraph](https://sourcegraph.com/github.com/labstack/echo-jwt/-/badge.svg?style=flat-square)](https://sourcegraph.com/github.com/labstack/echo-jwt?badge)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/labstack/echo-jwt/v4)
[![Go Report Card](https://goreportcard.com/badge/github.com/labstack/echo-jwt?style=flat-square)](https://goreportcard.com/report/github.com/labstack/echo-jwt)
[![Codecov](https://img.shields.io/codecov/c/github/labstack/echo-jwt.svg?style=flat-square)](https://codecov.io/gh/labstack/echo-jwt)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/labstack/echo-jwt/master/LICENSE)

# Echo JWT middleware

JWT middleware for [Echo](https://github.com/labstack/echo) framework. This middleware uses by default [golang-jwt/jwt/v5](https://github.com/golang-jwt/jwt) 
as JWT implementation.

## Versioning

This repository does not use semantic versioning. MAJOR version tracks which Echo version should be used. MINOR version
tracks API changes (possibly backwards incompatible) and PATCH version is incremented for fixes.

For Echo `v4` use `v4.x.y` releases.
Minimal needed Echo versions:
* `v4.0.0` needs Echo `v4.7.0+`

`main` branch is compatible with the latest Echo version.

## Usage

Add JWT middleware dependency with go modules
```bash
go get github.com/labstack/echo-jwt/v4
```

Use as import statement
```go
import "github.com/labstack/echo-jwt/v4"
```

Add middleware in simplified form, by providing only the secret key
```go
e.Use(echojwt.JWT([]byte("secret")))
```

Add middleware with configuration options
```go
e.Use(echojwt.WithConfig(echojwt.Config{
  // ...
  SigningKey:             []byte("secret"),
  // ...
}))
```

Extract token in handler
```go
import "github.com/golang-jwt/jwt/v5"

// ...

e.GET("/", func(c echo.Context) error {
  token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
  if !ok {
    return errors.New("JWT token missing or invalid")
  }
  claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
  if !ok {
    return errors.New("failed to cast claims as jwt.MapClaims")
  }
  return c.JSON(http.StatusOK, claims)
})
```

## Full example

```go
package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte("secret"),
	}))

	e.GET("/", func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
```

Test with
```bash
curl -v -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ" http://localhost:8080
```

Output should be
```bash
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET / HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.81.0
> Accept: */*
> Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Sun, 27 Nov 2022 21:34:17 GMT
< Content-Length: 52
< 
{"admin":true,"name":"John Doe","sub":"1234567890"}
```
