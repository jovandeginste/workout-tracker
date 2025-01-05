# timeago
Basically it's golang port of https://github.com/hustcc/timeago.js

timeago is a go library used to format date with `*** time ago` statement.
- i18n supported.
- Time ago and time in supported.

such as 

```plain
just now
12 seconds ago
2 hours ago
3 days ago
3 weeks ago
2 years ago

in 12 seconds
in 3 minutes
in 24 days
in 6 months
```


Supports all i18n of timeago.js
 
- Install

```bash
go get github.com/sersh88/timeago
```
- Usage

```go
t := time.Date(2021, 1, 1, 20, 34, 58, 651387237, time.UTC)
log.Println(timeago.New(t).Format())
log.Println(timeago.New(t).WithLocale("ru").Format())
log.Println(timeago.New(t).WithLocale("ru").WithRelativeTime(t.Add(-time.Hour*24)).Format())
```

The default locale is `en`. If passed locale not exists, default will be used.
You can also define your own locale with `RegisterLocale()` function
```go
  // Example:
  func myOwnDeLocale(_ float64, index int) (ago string, in string) {
    var res = [][]string{
      {"gerade eben", "vor einer Weile"},
      {"vor %d Sekunden", "in %d Sekunden"},
      {"vor 1 Minute", "in 1 Minute"},
      {"vor %d Minuten", "in %d Minuten"},
      {"vor 1 Stunde", "in 1 Stunde"},
      {"vor %d Stunden", "in %d Stunden"},
      {"vor 1 Tag", "in 1 Tag"},
      {"vor %d Tagen", "in %d Tagen"},
      {"vor 1 Woche", "in 1 Woche"},
      {"vor %d Wochen", "in %d Wochen"},
      {"vor 1 Monat", "in 1 Monat"},
      {"vor %d Monaten", "in %d Monaten"},
      {"vor 1 Jahr", "in 1 Jahr"},
      {"vor %d Jahren", "in %d Jahren"},
    }[index]
    return res[0], res[1]
  }
  timeago.RegisterLocale("de", myOwnDeLocale)
```
passed arguments to locale function have the same meaning as in [timeago.js](https://timeago.org/)
