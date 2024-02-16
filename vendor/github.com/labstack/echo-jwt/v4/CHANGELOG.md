# Changelog

## v4.2.0 - 2023-01-26

**Breaking change:** [JWT](github.com/golang-jwt/jwt) has been upgraded to `v5`. Check/test all your code involved with JWT tokens/claims. Search for `github.com/golang-jwt/jwt/v4` 
and replace it with `github.com/golang-jwt/jwt/v5`

**Enhancements**

* Upgrade `golang-jwt/jwt` library to `v5` [#9](https://github.com/labstack/echo-jwt/pull/9)


## v4.1.0 - 2023-01-26

**Enhancements**

* Add TokenExtractionError and TokenParsingError types to help distinguishing error source in ErrorHandler [#6](https://github.com/labstack/echo-jwt/pull/6)


## v4.0.1 - 2023-01-24

**Fixes**

* Fix data race in error path [#4](https://github.com/labstack/echo-jwt/pull/4)


**Enhancements**

* add TokenError as error returned when parsing fails [#3](https://github.com/labstack/echo-jwt/pull/3)


## v4.0.0 - 2022-12-27

* First release
