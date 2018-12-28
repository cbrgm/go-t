.PHONY: all check-license build install test lint fmt vet clean

all: check-license build test

PROJECT=go-t
GITHUB_URL=github.com/cbrgm/go-t
BIN?=t
VERSION?=$(shell cat VERSION)

GO := CGO_ENABLED=0 go
GOOS?=$(shell uname -s | tr A-Z a-z)
GOARCH?=$(subst x86_64,amd64,$(patsubst i%86,386,$(shell uname -m)))
OUT_DIR=_output

DATE := $(shell date -u '+%FT%T%z')

LDFLAGS += -X main.Version=$(VERSION)
LDFLAGS += -X main.Revision=$(DRONE_COMMIT)
LDFLAGS += -X "main.BuildDate=$(DATE)"
LDFLAGS += -extldflags '-static'

PACKAGES = $(shell go list ./... | grep -v /vendor/)

check-license:
	@echo ">> checking license headers"
	@./scripts/check_license.sh

build:
	@$(eval OUTPUT=$(OUT_DIR)/$(GOOS)/$(GOARCH)/$(BIN))
	@echo ">> building for $(GOOS)/$(GOARCH) to $(OUTPUT)"
	@mkdir -p $(OUT_DIR)/$(GOOS)/$(GOARCH)
	$(GO) build --installsuffix cgo -o $(OUTPUT) -v -ldflags '-w $(LDFLAGS)' $(GITHUB_URL)

install: build
	@$(eval OUTPUT=$(OUT_DIR)/$(GOOS)/$(GOARCH)/$(BIN))
	@echo ">> copying $(BIN) into $(GOPATH)/bin/$(BIN)"
	@cp $(OUTPUT) $(GOPATH)/bin/$(BIN)

test:
	@echo ">> running all tests"
	$(GO) test $(PKGS) --cover
	@for PKG in $(PACKAGES); do $(GO) test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

fmt:
	$(GO) fmt $(PACKAGES)

vet:
	$(GO) vet $(PACKAGES)

clean:
	$(GO) clean -i ./...
	rm -rf dist/
	rm -rf _output/

release:
	@which gox > /dev/null; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mitchellh/gox; \
	fi
	CGO_ENABLED=0 gox -arch="386 amd64 arm" -verbose -ldflags '-w $(LDFLAGS)' -output="dist/$(PROJECT)-${DRONE_TAG}-{{.OS}}-{{.Arch}}" .