.PHONY: all clean test

all: build-all serve

clean:
	rm ./tmp/main ./assets/output.css

build-all: build-tw build-server

build-server:
	go build -o ./tmp/main ./

build-tw:
	npx tailwindcss -i ./assets/main.css -o ./assets/output.css

serve:
	./tmp/main
