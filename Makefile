TARGET_ONLY_PKGS=$(shell go list ./... 2> /dev/null | grep -v "/vendor/")
IGNORE_DEPS_GOLINT='vendor/.+\.go'
IGNORE_DEPS_GOVET='vendor/.+\.go'
IGNORE_DEPS_GOCYCLO='vendor/.+\.go'
HAVE_GOLINT:=$(shell which golint)
HAVE_GOCYCLO:=$(shell which gocyclo)
HAVE_GOCOV:=$(shell which gocov)

TARGETS=$(addprefix github.com/kaneshin/pigeon/cmd/,pigeon)
OBJS=$(notdir $(TARGETS))

all: $(TARGETS)

$(TARGETS):
	@go install -v $@

.PHONY: unit unit-report
unit: lint vet cyclo test
unit-report: lint vet cyclo test-report

.PHONY: lint vet cyclo test coverage test-report
lint: golint
	@echo "go lint"
	@lint=`golint ./...`; \
		lint=`echo "$$lint" | grep -E -v -e ${IGNORE_DEPS_GOLINT}`; \
		echo "$$lint"; if [ "$$lint" != "" ]; then exit 1; fi

vet:
	@echo "go vet"
	@vet=`go tool vet -all -structtags -shadow $(shell ls -d */ | grep -v "vendor") 2>&1`; \
		vet=`echo "$$vet" | grep -E -v -e ${IGNORE_DEPS_GOVET}`; \
		echo "$$vet"; if [ "$$vet" != "" ]; then exit 1; fi

cyclo: gocyclo
	@echo "gocyclo -over 20"
	@cyclo=`gocyclo -over 20 . 2>&1`; \
		cyclo=`echo "$$cyclo" | grep -E -v -e ${IGNORE_DEPS_GOCYCLO}`; \
		echo "$$cyclo"; if [ "$$cyclo" != "" ]; then exit 1; fi

test:
	@go test $(TARGET_ONLY_PKGS)

coverage: gocov
	@gocov test $(TARGET_ONLY_PKGS) | gocov report

test-report:
	@echo "Invoking test and coverage"
	@echo "" > coverage.txt
	@for d in $(TARGET_ONLY_PKGS); do \
		go test -coverprofile=profile.out -covermode=atomic $$d || exit 1; \
		[ -f profile.out ] && cat profile.out >> coverage.txt && rm profile.out || true; done

.PHONY: golint gocyclo gocov
golint:
ifndef HAVE_GOLINT
	@echo "Installing linter"
	@go get -u github.com/golang/lint/golint
endif

gocyclo:
ifndef HAVE_GOCYCLO
	@echo "Installing gocyclo"
	@go get -u github.com/fzipp/gocyclo
endif

gocov:
ifndef HAVE_GOCOV
	@echo "Installing gocov"
	@go get -u github.com/axw/gocov/gocov
endif
