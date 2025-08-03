export GOBIN := $(shell pwd)/tmp/bin
export PATH := $(GOBIN):$(PATH)

GIT_REF ?= $(shell git symbolic-ref HEAD)
GIT_REF_NAME ?= $(shell git branch --show-current)
GIT_REF_TYPE ?= branch
GIT_COMMIT ?= $(shell git rev-parse HEAD)
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
WT_OUTPUT_FILE ?= tmp/workout-tracker
WT_DEBUG_OUTPUT_FILE ?= tmp/wt-debug

THEME_SCREENSHOT_WIDTH ?= 1200
THEME_SCREENSHOT_HEIGHT ?= 900
TEMPL_PROXY_PORT=8090
TEMPL_APP_PORT=8080
TEMPL_VERSION ?= $(shell grep "github.com/a-h/templ" go.mod | awk '{print $$2}')
TEMPL_COMMAND ?= go run github.com/a-h/templ/cmd/templ@$(TEMPL_VERSION)

GO_TEST=go test -short -count 1 -mod vendor -covermode=atomic

BRANCH_NAME_DEPS ?= update-deps

.PHONY: all clean test build screenshots meta install-deps

all: clean install-deps test build

release-patch release-minor release-major:
	$(MAKE) release VERSION=$(shell go run github.com/mdomke/git-semver/v6@latest -target $(subst release-,,$@))

release:
	git tag -s -a $(VERSION) -m "Release $(VERSION)"
	@echo "Now run 'git push --tags' and create a new release"
	@echo "New release: https://github.com/jovandeginste/workout-tracker/releases/new"

install-deps:
	cd frontend && npm install

clean:
	rm -fv ./assets/output.css ./workout-tracker
	rm -rf ./tmp/ ./node_modules/ ./assets/dist/


watch/templ:
	$(TEMPL_COMMAND) generate --watch \
			--open-browser=false \
			--proxy="http://localhost:$(TEMPL_APP_PORT)" \
			--proxyport="$(TEMPL_PROXY_PORT)" \
			--proxybind="0.0.0.0"

watch/server:
	go run github.com/air-verse/air@latest \
			--build.full_bin           "APP_ENV=development $(WT_OUTPUT_FILE)" \
			--build.cmd                "make build-server notify-proxy" \
			--build.delay              1000 \
			--build.exclude_dir        "assets,docs,testdata,tmp,vendor" \
			--build.exclude_regex      "_test.go" \
			--build.exclude_unchanged  false \
			--build.include_ext        "go,html,json,yaml" \
			--build.stop_on_error      true \
			--screen.clear_on_rebuild  false 

watch/tailwind:
	npx tailwindcss \
			-i ./main.css -o ./assets/output.css --minify --watch=always

notify-proxy:
	$(TEMPL_COMMAND) generate \
			--notify-proxy --proxyport=$(TEMPL_PROXY_PORT)

dev-backend:
	$(MAKE) watch/templ &
	$(MAKE) watch/server

dev: 
	$(MAKE) watch/templ &
	$(MAKE) watch/server &
	$(MAKE) watch/tailwind &
	sleep infinity

dev-docker:
	docker compose -f docker-compose.dev.yaml up --build

build: build-server build-docker screenshots
meta: swagger screenshots changelog

build-cli: build-frontend build-templates
	go build \
			-ldflags "-X 'main.buildTime=$(BUILD_TIME)' -X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.gitRef=$(GIT_REF)' -X 'main.gitRefName=$(GIT_REF_NAME)' -X 'main.gitRefType=$(GIT_REF_TYPE)'" \
			-o $(WT_DEBUG_OUTPUT_FILE) ./cmd/wt-debug/

build-server:
	go build \
			-ldflags "-X 'main.buildTime=$(BUILD_TIME)' -X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.gitRef=$(GIT_REF)' -X 'main.gitRefName=$(GIT_REF_NAME)' -X 'main.gitRefType=$(GIT_REF_TYPE)'" \
			-o $(WT_OUTPUT_FILE) ./cmd/workout-tracker/

build-docker:
	docker build \
			-t workout-tracker --pull \
			--build-arg BUILD_TIME="$(BUILD_TIME)" \
			--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
			--build-arg GIT_REF="$(GIT_REF)" \
			--build-arg GIT_REF_NAME="$(GIT_REF_NAME)" \
			--build-arg GIT_REF_TYPE="$(GIT_REF_TYPE)" \
			.

swagger:
	go run github.com/swaggo/swag/cmd/swag@latest init \
			--parseDependency \
			--dir ./pkg/app/,./pkg/database/,./vendor/gorm.io/gorm/,./vendor/github.com/codingsince1985/geo-golang/ \
			--generalInfo api_handlers.go
	git commit docs/ -m "Update swagger" -m "changelog: ignore" || echo "No changes to commit"

build-frontend:
	cd frontend && npm run build

build-templates:
	$(TEMPL_COMMAND) generate

test-packages:
	$(GO_TEST) ./pkg/...

test-templates:
	$(GO_TEST) ./views/...

test-commands:
	$(GO_TEST) ./cmd/...

format-templates:
	find . -type f -name '*.templ' -exec templ fmt -v {} \;

serve:
	$(WT_OUTPUT_FILE)

test: test-go test-assets

test-assets:
	prettier --check .


test-go: test-commands test-templates test-packages
	golangci-lint run --allow-parallel-runners

screenshots: generate-screenshots screenshots-theme screenshots-responsive screenshots-i18n

generate-screenshots: build-server
	export WT_BIND=[::]:8180 WT_DSN=screenshots.db; \
			$(WT_OUTPUT_FILE) & \
			export SERVER_PID=$$!; \
			sleep 1; \
			K6_BROWSER_ARGS="force-dark-mode" k6 run screenshots.js; \
			kill $${SERVER_PID}

screenshots-i18n:
	magick convert -delay 400 docs/profile-*.png docs/profile.gif

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
	montage \
			-font Liberation-Sans \
			-density 300 \
			-tile 3x0 \
			-geometry +5+5 \
			-background none \
			docs/dashboard-responsive.png docs/single_workout-responsive.png docs/statistics-responsive.png docs/responsive.png

go-cover:
	go test -short -count 1 -mod vendor -covermode=atomic -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm -vf coverage.out

update-deps:
	# Check no changes
	@if [[ "$$(git status --porcelain | wc -l)" -gt 0 ]]; then echo "There are changes; please commit or stash them first"; exit 1; fi
	# Check if branch exists locally or remotely
	@if git show-ref --verify --quiet refs/heads/$(BRANCH_NAME_DEPS); then echo "Branch $(BRANCH_NAME_DEPS) already exists locally. Aborting."; exit 1; fi
	@if git ls-remote --exit-code --heads origin $(BRANCH_NAME_DEPS); then echo "Branch $(BRANCH_NAME_DEPS) already exists remotely. Aborting."; exit 1; fi
	git switch --create $(BRANCH_NAME_DEPS)
	npm update
	go get -u -t ./...
	go mod tidy
	go mod vendor
	git add .
	git commit -m "build(deps): Update Go and frontend dependencies"
	git push origin $(BRANCH_NAME_DEPS)

changelog:
	git cliff -o CHANGELOG.md
	prettier --write CHANGELOG.md
	git commit CHANGELOG.md -m "Update changelog" -m "changelog: ignore"
