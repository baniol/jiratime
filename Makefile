# source: https://github.com/rightscale/go-boilerplate/blob/master/Makefile

NAME=jiratime

VERSION=0.2.3

DEPEND=golang.org/x/tools/cmd/cover \
		github.com/mattn/goveralls \
       	github.com/Masterminds/glide 
# github.com/golang/lint/golint

# PACKAGES=$(go list ./... |grep -v /vendor/)

# TRAVIS_BRANCH?=master
DATE=$(shell date '+%F %T')
TRAVIS_COMMIT?=$(shell git symbolic-ref HEAD | cut -d"/" -f 3)

# VERSION=$(NAME) $(DATE) - $(TRAVIS_COMMIT)
VFLAG=-X 'main.VERSION=$(VERSION) $(TRAVIS_COMMIT)'

.PHONY: depend clean default

default: $(NAME)
$(NAME): $(shell find . -name \*.go)
	go build -ldflags "$(VFLAG)" -o $(NAME) .

# the standard build produces a "local" executable, a linux tgz, and a darwin (macos) tgz
# uncomment and join the windows zip if you need it
build: $(NAME) \
		build/$(NAME)-darwin-amd64.tgz \
		build/$(NAME)-linux-amd64.tgz
# build/$(NAME)-linux-arm.tgz \
# build/$(NAME)-windows-amd64.zip

# create a tgz with the binary and any artifacts that are necessary
# note the hack to allow for various GOOS & GOARCH combos, sigh
build/$(NAME)-%.tgz: *.go
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -ldflags "$(VFLAG)" -o build/$(NAME)/$(NAME) .
	chmod +x build/$(NAME)/$(NAME)
	cp readme.md build/$(NAME)/
	tar -zcf $@ -C build ./$(NAME)
	rm -r build/$(NAME)

# build/$(NAME)-%.zip: *.go
# 	touch $@

# Installing build dependencies. You will need to run this once manually when you clone the repo
depend:
	go get -v $(DEPEND)
	glide install

# run gofmt and complain if a file is out of compliance
# run go vet and similarly complain if there are issues
# run go lint and complain if there are issues
# TODO: go tool vet is currently broken with the vendorexperiement
lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	for pkg in $$(go list ./... |grep -v /vendor/); do go vet $$pkg; done
#go tool vet -v -composites=false *.go
#go tool vet -v -composites=false **/*.go
# for pkg in $$(go list ./... |grep -v /vendor/); do golint $$pkg; done

travis-test: cover

test:
	go test $$(go list ./... |grep -v /vendor/) -v -cover

cover: lint
	echo "mode: count" > coverage-all.out
	for pkg in `go list ./... |grep -v /vendor/`; do \
		go test -v -covermode=count -coverprofile=coverage.out $$pkg; \
		tail -n +2 coverage.out >> coverage-all.out; \
	done
	goveralls -coverprofile=coverage-all.out -service travis-ci -repotoken $(COVERALLS_TOKEN)

cover-local:
	echo "mode: count" > coverage-all.out
	for pkg in `go list ./... |grep -v /vendor/`; do \
		go test -v -covermode=count -coverprofile=coverage.out $$pkg; \
		tail -n +2 coverage.out >> coverage-all.out; \
	done
	go tool cover -html=coverage-all.out
	rm coverage.out
	rm coverage-all.out

	
