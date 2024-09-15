GIT_REF ?= $(shell git symbolic-ref HEAD)
GIT_REF_NAME ?= $(shell git branch --show-current)
GIT_REF_TYPE ?= branch
GIT_COMMIT ?= $(shell git rev-parse HEAD)
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
OUTPUT_FILE ?= tmp/main

THEME_SCREENSHOT_WIDTH ?= 1200
THEME_SCREENSHOT_HEIGHT ?= 900

.PHONY: all clean test build screenshots meta translations install-dev-deps install-deps

all: clean install-deps test build

install-dev-deps:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/air-verse/air@latest

install-deps:
	npm install

clean:
	rm -fv ./assets/output.css ./workout-tracker
	rm -rf ./tmp/ ./node_modules/ ./assets/dist/

dev:
	air

build: build-dist build-tw build-server build-docker build-translations screenshots
meta: swagger screenshots changelog

build-server:
	go build \
		-ldflags "-X 'main.buildTime=$(BUILD_TIME)' -X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.gitRef=$(GIT_REF)' -X 'main.gitRefName=$(GIT_REF_NAME)' -X 'main.gitRefType=$(GIT_REF_TYPE)'" \
		-o $(OUTPUT_FILE) ./

build-docker:
	docker build -t workout-tracker --pull \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
		--build-arg GIT_REF="$(GIT_REF)" \
		--build-arg GIT_REF_NAME="$(GIT_REF_NAME)" \
		--build-arg GIT_REF_TYPE="$(GIT_REF_TYPE)" \
		.

swagger:
	swag init \
		--parseDependency \
		--dir ./pkg/app/,./,./vendor/gorm.io/gorm/,./vendor/github.com/codingsince1985/geo-golang/ \
		--generalInfo api_handlers.go

build-tw:
	npx tailwindcss -i ./main.css -o ./assets/output.css

clean-dist:
	rm -rf ./assets/dist/

build-dist: clean-dist
	mkdir -p ./assets/dist/images
	cp -v ./node_modules/fullcalendar/index.global.min.js ./assets/dist/fullcalendar.min.js
	cp -v ./node_modules/leaflet/dist/leaflet.css ./assets/dist/
	cp -v ./node_modules/leaflet/dist/images/* ./assets/dist/images/
	cp -v ./node_modules/leaflet/dist/leaflet.js ./assets/dist/
	cp -v ./node_modules/shareon/dist/shareon.iife.js  ./assets/dist/
	cp -v ./node_modules/shareon/dist/shareon.min.css ./assets/dist/
	cp -v ./node_modules/apexcharts/dist/apexcharts.min.js ./assets/dist/
	cp -v ./node_modules/apexcharts/dist/apexcharts.css ./assets/dist/
	cp -v ./node_modules/htmx.org/dist/htmx.min.js ./assets/dist/

watch-tw:
	npx tailwindcss -i ./main.css -o ./assets/output.css --watch

build-translations: translations

translations:
	xspreak -o translations/en.json -f json --template-keyword "i18n" -t "views/**/*.html"
	prettier --write translations/*.json

serve:
	$(OUTPUT_FILE)

test: test-go test-assets

test-assets:
	prettier --check .

test-go:
	go test -short -count 1 -mod vendor -covermode=atomic ./...
	golangci-lint run --allow-parallel-runners

screenshots: generate-screenshots screenshots-theme screenshots-responsive

generate-screenshots:
	K6_BROWSER_ARGS="force-dark-mode" k6 run screenshots.js

screenshots-theme:
	mkdir -p tmp/
	convert docs/single_workout-dark.png \
		-resize $(THEME_SCREENSHOT_WIDTH)x$(THEME_SCREENSHOT_HEIGHT)\! \
		tmp/dark_resized.jpg
	convert docs/single_workout-light.png \
		-resize $(THEME_SCREENSHOT_WIDTH)x$(THEME_SCREENSHOT_HEIGHT)\! \
		tmp/light_resized.jpg
	convert -size $(THEME_SCREENSHOT_WIDTH)x$(THEME_SCREENSHOT_HEIGHT) \
		xc:white -draw "polygon 0,0 $(THEME_SCREENSHOT_WIDTH),0 $(THEME_SCREENSHOT_WIDTH),$(THEME_SCREENSHOT_HEIGHT)" \
		tmp/mask.png
	convert tmp/dark_resized.jpg tmp/light_resized.jpg tmp/mask.png \
		-composite docs/single_workout-theme.jpg
	rm -f tmp/dark_resized.jpg tmp/light_resized.jpg tmp/mask.png

screenshots-responsive:
	montage -font Liberation-Sans -density 300 -tile 3x0 -geometry +5+5 -background none docs/dashboard-responsive.png docs/single_workout-responsive.png docs/statistics-responsive.png docs/responsive.png

go-cover:
	go test -short -count 1 -mod vendor -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm -vf coverage.out

update-deps:
	npm update
	go get -u -d -t ./...
	go mod tidy
	go mod vendor

changelog:
	git cliff -o CHANGELOG.md
	prettier --write CHANGELOG.md
	git commit CHANGELOG.md -m "Update changelog" -m "changelog: ignore"
