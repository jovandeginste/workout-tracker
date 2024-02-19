GIT_TAG ?= $(shell git describe --tags)

.PHONY: all clean test

all: build-all serve

clean:
	rm ./tmp/main ./assets/output.css

build-all: build-tw build-server build-docker

build-server:
	go build -ldflags "-X main.version=$(GIT_TAG)" -o ./tmp/main ./

build-docker:
	docker build -t workouts --pull .

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
