GIT_REF ?= $(shell git describe --tags)
GIT_SHA ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u --rfc-3339=seconds)

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
		-ldflags "-X 'main.buildTime=$(BUILD_TIME)' -X 'main.gitCommit=$(GIT_SHA)' -X 'main.gitRef=$(GIT_REF)'" \
		-o ./tmp/main ./

build-docker:
	docker build -t workout-tracker --pull .

build-tw:
	npx tailwindcss -i ./main.css -o ./assets/output.css

watch-tw:
	npx tailwindcss -i ./main.css -o ./assets/output.css --watch

serve:
	./tmp/main

test: test-views test-go test-assets
test-views:
	prettier --check views/
test-assets:
	prettier --check assets/

test-go:
	go test -short -count 1 -mod vendor -covermode=atomic ./...
	golangci-lint run --allow-parallel-runners --fix
