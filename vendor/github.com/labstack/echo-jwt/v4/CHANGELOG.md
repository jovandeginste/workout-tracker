# Changelog

## v4.3.0 - 2024-12-04

**Enhancements**

* Update Echo dependency to v4.13.0 by @aldas in [#28](https://github.com/labstack/echo-jwt/pull/28)


## v4.2.1 - 2024-12-04

**Enhancements**

* Return HTTP status 400 if missing JWT by @kitloong in [#13](https://github.com/labstack/echo-jwt/pull/13)
* Update dependencies and CI flow by @aldas in [#21](https://github.com/labstack/echo-jwt/pull/21), [#24](https://github.com/labstack/echo-jwt/pull/24), [#27](https://github.com/labstack/echo-jwt/pull/27)
* Improve readme formatting by @aldas in [#25](https://github.com/labstack/echo-jwt/pull/25)


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
