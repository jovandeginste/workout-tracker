GIT_REF ?= $(shell git describe --tags)
GIT_COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u --rfc-3339=seconds)
OUTPUT_FILE ?= tmp/main

.PHONY: all clean test build

all: clean install-deps test build

install-deps:
	npm install

clean:
	rm -fv ./assets/output.css ./workout-tracker
	rm -rf ./tmp/ ./node_modules/

dev:
	air

build: build-tw build-server build-docker

build-server:
	go build \
		-ldflags "-X 'main.buildTime=$(BUILD_TIME)' -X 'main.gitCommit=$(GIT_COMMIT)' -X 'main.gitRef=$(GIT_REF)'" \
		-o $(OUTPUT_FILE) ./

build-docker:
	docker build -t workout-tracker --pull \
		--build-arg BUILD_TIME="$(BUILD_TIME)" \
		--build-arg GIT_COMMIT="$(GIT_COMMIT)" \
		--build-arg GIT_REF="$(GIT_REF)" \
		.

build-tw:
	npx tailwindcss -i ./main.css -o ./assets/output.css

watch-tw:
	npx tailwindcss -i ./main.css -o ./assets/output.css --watch

serve:
	$(OUTPUT_FILE)

test: test-views test-go test-assets
test-views:
	prettier --check views/
test-assets:
	prettier --check assets/

test-go:
	go test -short -count 1 -mod vendor -covermode=atomic ./...
	golangci-lint run --allow-parallel-runners --fix
