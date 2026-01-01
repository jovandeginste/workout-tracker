fmt:
	go fmt ./...

.PHONY:pb
pb:
	buf generate

test:
	golangci-lint run ./...
	go test -v -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage.out -o=coverage.html

bench:
	go test -v -bench=. ./...

dep-licenses:
	go-licenses save ./ --save_path=THIRD_PARTY_LICENSES 
	build_notice.sh
