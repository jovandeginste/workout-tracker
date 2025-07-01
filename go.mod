module github.com/jovandeginste/workout-tracker/v2

go 1.24.1

replace github.com/anyappinc/fitbit v0.0.3 => github.com/jovandeginste/fitbit v0.0.4-0.20250213164811-b0b3b27c3a84

require (
	github.com/a-h/templ v0.3.906
	github.com/alexedwards/scs/gormstore v0.0.0-20250417082927-ab20b3feb5e9
	github.com/alexedwards/scs/v2 v2.8.0
	github.com/anyappinc/fitbit v0.0.3
	github.com/aquasecurity/table v1.10.0
	github.com/biter777/countries v1.7.5
	github.com/cat-dealer/go-rand/v2 v2.0.0
	github.com/codingsince1985/geo-golang v1.8.5
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc
	github.com/fsouza/slognil v0.4.2
	github.com/galeone/tcx v1.0.1-0.20230114151622-8168e1e47884
	github.com/glebarez/sqlite v1.11.0
	github.com/go-gorm/caches/v4 v4.0.5
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/gomarkdown/markdown v0.0.0-20250311123330-531bef5e742b
	github.com/google/go-querystring v1.1.0
	github.com/google/uuid v1.6.0
	github.com/invopop/ctxi18n v0.9.0
	github.com/labstack/echo-jwt/v4 v4.3.1
	github.com/labstack/echo/v4 v4.13.4
	github.com/labstack/gommon v0.4.2
	github.com/lmittmann/tint v1.1.2
	github.com/mattn/go-isatty v0.0.20
	github.com/microcosm-cc/bluemonday v1.0.27
	github.com/muktihari/fit v0.24.5
	github.com/orandin/slog-gorm v1.4.0
	github.com/paulmach/orb v0.11.1
	github.com/ringsaturn/tzf v1.0.0
	github.com/samber/slog-echo v1.16.1
	github.com/sersh88/timeago v1.0.0
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spazzymoto/echo-scs-session v1.0.0
	github.com/spf13/cobra v1.9.1
	github.com/spf13/viper v1.20.1
	github.com/stackus/hxgo v0.3.0
	github.com/stretchr/testify v1.10.0
	github.com/swaggo/swag v1.16.4
	github.com/tkrajina/gpxgo v1.4.0
	github.com/westphae/geomag v1.0.2
	golang.org/x/crypto v0.39.0
	golang.org/x/text v0.26.0
	gorm.io/datatypes v1.2.5
	gorm.io/driver/mysql v1.6.0
	gorm.io/driver/postgres v1.6.0
	gorm.io/gorm v1.30.0
	resty.dev/v3 v3.0.0-beta.3
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/fsnotify/fsnotify v1.9.0 // indirect
	github.com/glebarez/go-sqlite v1.22.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/spec v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/go-sql-driver/mysql v1.9.3 // indirect
	github.com/go-viper/mapstructure/v2 v2.3.0 // indirect
	github.com/gorilla/css v1.0.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pelletier/go-toml/v2 v2.2.4 // indirect
	github.com/philhofer/vec v0.0.0-20140421144027-536fc796d369 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/ringsaturn/tzf-rel-lite v0.0.2025-b // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/sagikazarmark/locafero v0.9.0 // indirect
	github.com/samber/lo v1.51.0 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.14.0 // indirect
	github.com/spf13/cast v1.9.2 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/tidwall/geoindex v1.7.0 // indirect
	github.com/tidwall/geojson v1.4.5 // indirect
	github.com/tidwall/rtree v1.10.0 // indirect
	github.com/twpayne/go-polyline v1.1.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	go.mongodb.org/mongo-driver v1.17.4 // indirect
	go.opentelemetry.io/otel v1.37.0 // indirect
	go.opentelemetry.io/otel/trace v1.37.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/oauth2 v0.30.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	golang.org/x/tools v0.34.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	modernc.org/libc v1.66.1 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.11.0 // indirect
	modernc.org/sqlite v1.38.0 // indirect
)

replace github.com/tkrajina/gpxgo v1.4.0 => github.com/jovandeginste/gpxgo v1.4.1-0.20250629150855-db85929d31f6
