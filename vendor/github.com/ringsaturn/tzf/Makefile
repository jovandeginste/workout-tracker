PROTO_FILES=$(shell find pb -name *.proto)

fmt:
	find pb/ -iname *.proto | xargs clang-format -i --style=Google
	go fmt ./...

.PHONY:pb
pb:
	protoc  --proto_path=. \
			--doc_out=. --doc_opt=html,pb.html,source_relative \
			--go_out=paths=source_relative:. \
			$(PROTO_FILES)

test:
	golangci-lint run ./...
	go test -v -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage.out -o=coverage.html

comparetzpb_gen:
	go run cmd/comparetzpb/main.go

bench:
	go test -v -bench=. ./...
