test:
	go vet -vettool=$(which shadow)
	go test -v -count=1 -race $(go list ./...)
