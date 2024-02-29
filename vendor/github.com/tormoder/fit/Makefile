PWD		:= $(shell pwd)
BIN		:= bin/

GO		:= go
GO_BIN		:= GOBIN=$(PWD)/$(BIN)
GOMODULE_OFF	:= GO111MODULE=off

FIT_PKGS 	:= ./...
FIT_DIRS 	:= $(shell find . -type f -not -path "*vendor*" -not -path "./.git*" -not -path "*testdata*" -name "*.go" -printf "%h\n" | sort -u)
FIT_PKG_PATH 	:= github.com/tormoder/fit
FITGEN_REL_PATH := ./cmd/fitgen

GOFUZZ_PKG_PATH	:= github.com/dvyukov/go-fuzz

DECODE_BENCH_NAME := DecodeActivity$$/Small
DECODE_BENCH_TIME := 5s

.PHONY: all
all: build test testrace checkfull

.PHONY: build
build:
	@echo "$(GO) build:"
	$(GO) build -v $(FIT_PKGS)

.PHONY: test
test:
	@echo "$(GO) test:"
	$(GO) test -v -cpu=2 $(FIT_PKGS)

.PHONY: testrace
testrace:
	@echo "$(GO) test -race:"
	$(GO) test -v -cpu=1,2,4 -race $(FIT_PKGS)

.PHONY: bench
bench:
	$(GO) test -v -run=^$$ -bench=. -benchtime=5s $(FIT_PKGS)

.PHONY: fitgen
fitgen:
	$(GO) install $(FITGEN_REL_PATH)

.PHONY: gofuzz
gofuzz:
	@echo "This target must be ran with the repo located under GOPATH with GOBIN set"
	@echo "Use '-workdir=workdir' to use inital corpus copied by 'gofuzzclean' target"
	$(GOMODULE_OFF) $(GO) get -u $(GOFUZZ_PKG_PATH)/go-fuzz
	$(GOMODULE_OFF) $(GO) get -u $(GOFUZZ_PKG_PATH)/go-fuzz-build
	go-fuzz-build $(FIT_PKG_PATH)

.PHONY: gofuzzclean
gofuzzclean: gofuzz
	rm -rf workdir/
	mkdir -p workdir/corpus
	find testdata -name \*.fit -exec cp {} workdir/corpus/ \;

.PHONY: clean
clean:
	$(GO) clean -i ./...
	rm -f fit-fuzz.zip
	find . -name '*.prof' -type f -exec rm -f {} \;
	find . -name '*.test' -type f -exec rm -f {} \;
	find . -name '*.current' -type f -exec rm -f {} \;
	find . -name '*.current.gz' -type f -exec rm -f {} \;

.PHONY: gcoprofile 
gcoprofile:
	git checkout types.go messages.go profile.go types_string.go

.PHONY: profcpu
profcpu:
	$(GO) test -run=^$$ -cpuprofile=cpu.prof -bench=$(DECODE_BENCH_NAME) -benchtime=$(DECODE_BENCH_TIME)
	$(GO) tool pprof fit.test cpu.prof

.PHONY: profmem
profmem:
	$(GO) test -run^$$ =-memprofile=allocmem.prof -bench=$(DECODE_BENCH_NAME) -benchtime=$(DECODE_BENCH_TIME)
	$(GO) tool pprof -alloc_space fit.test allocmem.prof

.PHONY: profobj
profobj:
	$(GO) test -run=^$$ -memprofile=allocobj.prof -bench=$(DECODE_BENCH_NAME) -benchtime=$(DECODE_BENCH_TIME)
	$(GO) tool pprof -alloc_objects fit.test allocobj.prof

.PHONY: mdgen
mdgen:
	godoc2md $(FIT_PKG_PATH) Fit Header CheckIntegrity > MainApiReference.md

.PHONY: check
check:
	@echo "check (basic)":
	@echo "gofmt (simplify)"
	@gofmt -s -l .
	@echo "$(GO) vet"
	@$(GO) vet $(FIT_PKGS)

.PHONY: checkfull
checkfull: checkdeps
	@echo "check (full):"
	@echo "gofmt (simplify)"
	@! gofmt -s -l . | grep -vF 'No Exceptions'
	@echo "goimports"
	@! $(BIN)goimports -l . | grep -vF 'No Exceptions'
	@echo "gofumpt"
	@! $(BIN)gofumpt -l . | grep -vE '(types.go|types_string.go|types_man.go)'
	@echo "vet"
	@$(GO) vet $(FIT_PKGS)
	@echo "vet --shadow"
	@$(GO) vet -vettool=$(which shadow) $(FIT_PKGS)
	@echo "errcheck"
	@$(BIN)errcheck -ignore 'fmt:Fprinf*,bytes:Write*,archive/zip:Close,io:Close,Write' $(FIT_PKGS)
	@echo "ineffassign"
	@for dir in $(FIT_DIRS); do \
		$(BIN)ineffassign -n $$dir ; \
	done
	@echo "unconvert"
	@! $(BIN)unconvert $(FIT_PKGS) | grep -vF 'messages.go'
	@echo "misspell"
	@! $(BIN)misspell ./**/* | grep -vE '(messages.go|/vendor/|profile/testdata)'
	@echo "staticcheck"
	@$(BIN)staticcheck $(FIT_PKGS)

.PHONY: checkdeps
checkdeps:
	$(GO_BIN) $(GO) install github.com/client9/misspell/cmd/misspell
	$(GO_BIN) $(GO) install github.com/gordonklaus/ineffassign
	$(GO_BIN) $(GO) install github.com/kisielk/errcheck
	$(GO_BIN) $(GO) install github.com/mdempsky/unconvert
	$(GO_BIN) $(GO) install golang.org/x/tools/cmd/goimports
	$(GO_BIN) $(GO) install honnef.co/go/tools/cmd/staticcheck
	$(GO_BIN) $(GO) install mvdan.cc/gofumpt
