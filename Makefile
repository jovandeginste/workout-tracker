GIT_TAG ?= $(shell git describe --tags)

.PHONY: all clean test build

all: clean install-deps test build

install-deps:
	npm install

clean:
	rm -fv ./tmp/main ./assets/output.css ./workout-tracker
	rm -rf ./node_modules/

dev:
	air

build: build-tw build-server build-docker

build-server: build-tw
	go build -ldflags "-X main.version=$(GIT_TAG)" -o ./tmp/main ./

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
	go test -short -count 1 -mod vendor ./...
	golangci-lint run --allow-parallel-runners --fix
