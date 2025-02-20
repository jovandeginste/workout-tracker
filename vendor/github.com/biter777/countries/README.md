# countries

Countries - ISO 639, ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2, Alpha-3) countries codes with subdivisions and names (on eng and rus), ISO4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, IANA ccTLD countries domains, FIPS, IOC/NOC and FIFA codes, **VERY VERY FAST**, NO maps[], NO slices[], NO init() funcs, NO external links/files/data, NO interface{}, NO specific dependencies, compatible with Databases/JSON/BSON/GOB/XML/CSV, Emoji countries flags and currencies support, UN M.49, FIFA codes, full support ISO 639-1, ISO 3166-1, ISO 3166-2, ISO 4217, ITU-T E.164, Unicode CLDR and IANA ccTLD standarts.

**Supported standarts:**
   - ISO 639-1
   - ISO 3166-1
   - ISO 3166-2
   - ISO 4217
   - ICANN
   - IANA ccTLD
   - ITU-T E.164
   - IOC
   - NOC
   - FIFA
   - FIPS
   - UN M.49
   - Unicode CLDR 
   - Unicode Emoticons Flags
   - Unicode Emoticons Currencies 
   - UN World Countries Capitals

[![GoDev](https://img.shields.io/badge/godev-reference-5b77b3)](https://pkg.go.dev/github.com/biter777/countries?tab=doc)
[![GoAwesome](https://img.shields.io/badge/awesome%20go-reference-5b77b3)](https://awesome-go.com/utilities/)
[![Coder](https://img.shields.io/badge/coder-reference-5b77b3)](https://coder.social/biter777/countries)
[![DOI](https://zenodo.org/badge/182808313.svg)](https://zenodo.org/badge/latestdoi/182808313)
[![codeclimate](https://codeclimate.com/github/biter777/countries/badges/gpa.svg)](https://codeclimate.com/github/biter777/countries)
[![GolangCI](https://golangci.com/badges/github.com/biter777/countries.svg?style=flat)](https://golangci.com/r/github.com/biter777/countries)
[![GoReport](https://goreportcard.com/badge/github.com/biter777/countries)](https://goreportcard.com/report/github.com/biter777/countries)
[![Codiga](https://img.shields.io/badge/codiga%20quality-A+-brightgreen)](https://app.codiga.io/project/3255/dashboard)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/08eb1d2ff62e465091b3a288ae078a96)](https://www.codacy.com/manual/biter777/countries?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=biter777/countries&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/biter777/countries/branch/master/graph/badge.svg)](https://codecov.io/gh/biter777/countries)
[![Coverage Status](https://coveralls.io/repos/github/biter777/countries/badge.svg?branch=master)](https://coveralls.io/github/biter777/countries?branch=master)
[![Coverage](https://img.shields.io/badge/coverage-gocover.io-brightgreen)](https://gocover.io/github.com/biter777/countries)
[![ISO](https://img.shields.io/badge/powered%20by-ISO-brightgreen)](https://www.iso.org/)
[![ITU](https://img.shields.io/badge/powered%20by-ITU-brightgreen)](https://www.itu.int/)
[![IANA](https://img.shields.io/badge/powered%20by-IANA-brightgreen)](http://www.iana.org/)
[![ICANN](https://img.shields.io/badge/powered%20by-ICANN-brightgreen)](https://www.icann.org/)
[![M49](https://img.shields.io/badge/powered%20by-UN%20M49-brightgreen)](https://unstats.un.org/unsd/methodology/m49/)
[![CLDR](https://img.shields.io/badge/powered%20by-CLDR-brightgreen)](https://cldr.unicode.org/)
[![License](https://img.shields.io/badge/License-BSD%202--Clause-brightgreen.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Build status](https://ci.appveyor.com/api/projects/status/t9lpor9o8tpacpmr/branch/master?svg=true)](https://ci.appveyor.com/project/biter777/countries/branch/master)
[![Build Status](https://github.com/biter777/countries/actions/workflows/go.yml/badge.svg)](https://github.com/biter777/countries/actions/workflows/go.yml)
[![CLDR](https://img.shields.io/badge/deepsource-passing-brightgreen)]([https://cldr.unicode.org/](https://deepsource.io/gh/biter777/countries))
<a href="//www.dmca.com/Protection/Status.aspx?ID=7a019cc5-ec73-464b-9707-4b33726f348f" title="DMCA.com Protection Status" class="dmca-badge"> <img src ="https://img.shields.io/badge/DMCA-protected-brightgreen" alt="DMCA.com Protection Status" /></a>
[![Dependencies Free](https://img.shields.io/badge/dependencies-free-brightgreen)](https://pkg.go.dev/github.com/biter777/countries?tab=imports)
[![Gluten Free](https://img.shields.io/badge/gluten-free-brightgreen)](https://www.scsglobalservices.com/services/gluten-free-certification)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen)](https://github.com/biter777/countries/pulls)
[![DepShield Badge](https://depshield.sonatype.org/badges/biter777/countries/depshield.svg)](https://depshield.github.io)
[![Stars](https://img.shields.io/github/stars/biter777/countries?label=Please%20like%20us&style=social)](https://github.com/biter777/countries/stargazers)
<br/>

## installation

```shell
go get github.com/biter777/countries
```

## usage

```go
countryJapan := countries.Japan
fmt.Printf("Country name in english: %v\n", countryJapan)                   // Japan
fmt.Printf("Country name in russian: %v\n", countryJapan.StringRus())       // Япония
fmt.Printf("Country ISO-3166 digit code: %d\n", countryJapan)               // 392
fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", countryJapan.Alpha2())    // JP
fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", countryJapan.Alpha3())    // JPN
fmt.Printf("Country IOC/NOC code: %v\n", countryJapan.IOC())                // JPN
fmt.Printf("Country FIFA code: %v\n", countryJapan.FIFA())                  // JPN
fmt.Printf("Country FIPS code: %v\n", countryJapan.FIPS())                  // JA
fmt.Printf("Country Capital: %v\n", countryJapan.Capital())                 // Tokyo
fmt.Printf("Country ITU-T E.164 call code: %v\n", countryJapan.CallCodes()) // +81
fmt.Printf("Country ccTLD domain: %v\n", countryJapan.Domain())             // .jp
fmt.Printf("Country UN M.49 region name: %v\n", countryJapan.Region())      // Asia
fmt.Printf("Country UN M.49 region code: %d\n", countryJapan.Region())      // 142
fmt.Printf("Country emoji/flag: %v\n", countryJapan.Emoji())                // 🇯🇵
fmt.Printf("Country Subdivisions: %v\n", countryJapan.Subdivisions())       // Hokkaido Aomori Iwate Miyagi Akita Yamagata Fukushima Ibaraki Tochigi Gunma Saitama Chiba Tokyo Kanagawa Niigata Toyama Ishikawa Fukui Yamanashi Nagano Gifu Shizuoka Aichi Mie Shiga Kyoto Osaka Hyogo Nara Wakayama Tottori Shimane Okayama Hiroshima Yamaguchi Tokushima Kagawa Ehime Kochi Fukuoka Saga Nagasaki Kumamoto Oita Miyazaki Kagoshima Okinawa

currencyJapan := countryJapan.Currency()
fmt.Printf("Country ISO-4217 Currency name in english: %v\n", currencyJapan)           // Yen
fmt.Printf("Country ISO-4217 Currency digit code: %d\n", currencyJapan)                // 392
fmt.Printf("Country ISO-4217 Currency Alpha code: %v\n", currencyJapan.Alpha())        // JPY
fmt.Printf("Country Currency emoji: %v\n", currencyJapan.Emoji())                      // 💴
fmt.Printf("Country of Currency %v: %v\n\n", currencyJapan, currencyJapan.Countries()) // Japan

// OR you can alternative use:
japanInfo := countries.Japan.Info()
fmt.Printf("Country name in english: %v\n", japanInfo.Name)                          // Japan
fmt.Printf("Country ISO-3166 digit code: %d\n", japanInfo.Code)                      // 392
fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", japanInfo.Alpha2)                  // JP
fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", japanInfo.Alpha3)                  // JPN
fmt.Printf("Country IOC/NOC code: %v\n", japanInfo.IOC)                              // JPN
fmt.Printf("Country FIFA code: %v\n", japanInfo.FIFA)                                // JPN
fmt.Printf("Country FIPS code: %v\n", japanInfo.FIPS)                                // JA
fmt.Printf("Country Capital: %v\n", japanInfo.Capital)                               // Tokyo
fmt.Printf("Country ITU-T E.164 call code: %v\n", japanInfo.CallCodes)               // +81
fmt.Printf("Country ccTLD domain: %v\n", japanInfo.Domain)                           // .jp
fmt.Printf("Country UN M.49 region name: %v\n", japanInfo.Region)                    // Asia
fmt.Printf("Country UN M.49 region code: %d\n", japanInfo.Region)                    // 142
fmt.Printf("Country emoji/flag: %v\n", japanInfo.Emoji)                              // 🇯🇵
fmt.Printf("Country ISO-4217 Currency name in english: %v\n", japanInfo.Currency)    // Yen
fmt.Printf("Country ISO-4217 Currency digit code: %d\n", japanInfo.Currency)         // 392
fmt.Printf("Country ISO-4217 Currency Alpha code: %v\n", japanInfo.Currency.Alpha()) // JPY
fmt.Printf("Country Subdivisions: %v\n", japanInfo.Subdivisions)                     // Hokkaido Aomori Iwate Miyagi Akita Yamagata Fukushima Ibaraki Tochigi Gunma Saitama Chiba Tokyo Kanagawa Niigata Toyama Ishikawa Fukui Yamanashi Nagano Gifu Shizuoka Aichi Mie Shiga Kyoto Osaka Hyogo Nara Wakayama Tottori Shimane Okayama Hiroshima Yamaguchi Tokushima Kagawa Ehime Kochi Fukuoka Saga Nagasaki Kumamoto Oita Miyazaki Kagoshima Okinawa

// Detection/Lookup usage
// Detect/Lookup by country name
country := countries.ByName("angola")
fmt.Printf("Country name in english: %v\n", country)                // Angola
fmt.Printf("Country ISO-3166 digit code: %d\n", country)            // 24
fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", country.Alpha2()) // AO
fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", country.Alpha3()) // AGO
// Detect/Lookup by country code
country = countries.ByName("AO")
fmt.Printf("Country name in english: %v\n", country.String())       // Angola
fmt.Printf("Country ISO-3166 digit code: %d\n", country)            // 24
fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", country.Alpha2()) // AO
fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", country.Alpha3()) // AGO
// Detect/Lookup by code/numeric
country = countries.ByNumeric(24)
fmt.Printf("Country name in english: %v\n", country)                // Angola
fmt.Printf("Country ISO-3166 digit code: %d\n", country)            // 24
fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", country.Alpha2()) // AO
fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", country.Alpha3()) // AGO

// Comparing usage
// Compare by code/numeric
if countries.ByName("angola") == countries.AGO {
	fmt.Println("Yes! It's Angola!") // Yes! It's Angola!
}
// Compare by name
if strings.EqualFold("angola", countries.AGO.String()) {
	fmt.Println("Yes! It's Angola!") // Yes! It's Angola!
}

// Database usage
type User struct {
	gorm.Model
	Name     string
	Country  countries.CountryCode
	Currency countries.CurrencyCode
}
user := &User{Name: "Helen", Country: countries.Slovenia, Currency: countries.CurrencyEUR}
db, err := gorm.Open("postgres", 500, "host=127.0.0.2 port=5432 user=usr password=1234567 dbname=db")
if err != nil {
	panic(err)
}
defer db.Close()
db.Create(user)
```

## Options

For Emoji use Emoji(). Enjoy!

```go
import "github.com/biter777/countries"
```

For more complex options, consult the [documentation](http://godoc.org/github.com/biter777/countries).

## Contributing

1. **Welcome pull requests, bug fixes and issue reports**

	[Contributors list](https://github.com/biter777/countries/graphs/contributors)
	
2. **Donate** - a donation isn't necessary, but it's welcome.

	<noscript><a href="https://liberapay.com/biter777/donate"><img alt="Donate using Liberapay" src="https://liberapay.com/assets/widgets/donate.svg"></a></noscript>
	[![ko-fi](https://www.ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/I2I61D1XZ) <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://usa.visa.com/dam/VCOM/regional/lac/ENG/Default/Partner%20With%20Us/Payment%20Technology/visapos/full-color-800x450.jpg"></a> <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://brand.mastercard.com/content/dam/mccom/brandcenter/thumbnails/mastercard_debit_sym_decal_web_105px.png"></a> <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://developer.apple.com/assets/elements/icons/apple-pay/apple-pay.svg"></a> <a href="https://pay.cloudtips.ru/p/94fc4268" target="_blank"><img height="30" src="https://developers.google.com/pay/api/images/brand-guidelines/google-pay-mark.png"></a> <br/>

3. **Star us** - give us a star, please, if it's not against your religion :)


	[![Stars](https://img.shields.io/github/stars/biter777/countries?label=Please%20like%20us&style=social)](https://github.com/biter777/countries/stargazers)

## Updating ISO 3166

Making use of changes to the [iso-codes](https://salsa.debian.org/iso-codes-team/iso-codes) project.

**TODO** create go generate capability to automatically generate from the json files
in the [data](https://salsa.debian.org/iso-codes-team/iso-codes/-/tree/main/data/)
directory of the iso-codes project.
