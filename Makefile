LINTIGNOREDEPS='vendor/.+\.go'

PIGEON_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/vendor/")

all: help

help:
	@echo "Please use \`make <target>' where <target> is one of"
	@echo "  build                   to go build the pigeon"
	@echo "  unit                    to run unit tests"
	@echo "  verify                  to verify tests"
	@echo "  lint                    to lint the pigeon"
	@echo "  vet                     to vet the pigeon"

build:
	@echo "go build pigeon and vendor packages"

unit: verify test

test:
	@go test $(PIGEON_ONLY_PKGS)

verify: lint vet

lint:
	@echo "go lint packages"
	@lint=`golint ./...`; \
	lint=`echo "$$lint" | grep -E -v -e ${LINTIGNOREDEPS}`; \
	echo "$$lint"; \
	if [ "$$lint" != "" ]; then exit 1; fi

vet:
	@echo "go vet packages"
	@go tool vet -all -structtags -shadow $(shell ls -d */ | grep -v "vendor")

vendor:
	glide install

install: install-pigeon install-pigeon-app
install-pigeon:
	@go install github.com/kaneshin/pigeon/tools/cmd/pigeon
install-pigeon-app:
	@go install github.com/kaneshin/pigeon/tools/cmd/pigeon-app

run-pigeon: install-pigeon
	@lime -bin=/tmp/pigeon-bin -ignore-pattern=\(\\.git\|vendor\) -build-pattern=.* ./tools/cmd/pigeon
run-pigeon-app: install-pigeon-app
	@lime -i -bin=/tmp/pigeon-bin -ignore-pattern=\(\\.git\|vendor\) ./tools/cmd/pigeon-app -port=8080 -- -label

