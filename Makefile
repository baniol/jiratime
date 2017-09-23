NAME=jiratime
VERSION=0.2.3

DEPEND=golang.org/x/tools/cmd/cover \
		github.com/mattn/goveralls \
       	github.com/Masterminds/glide 

DATE=$(shell date '+%F %T')
TRAVIS_COMMIT?=$(shell git symbolic-ref HEAD | cut -d"/" -f 3)

VFLAG=-X 'main.VERSION=$(VERSION) $(TRAVIS_COMMIT)'

.PHONY: depend clean default

default: $(NAME)
$(NAME): $(shell find . -name \*.go)
	go build -ldflags "$(VFLAG)" -o $(NAME) .

build: $(NAME) \
		build/$(NAME)-darwin-amd64.tgz \
		build/$(NAME)-linux-amd64.tgz

build/$(NAME)-%.tgz: *.go
	rm -rf build/$(NAME)
	mkdir -p build/$(NAME)
	tgt=$*; GOOS=$${tgt%-*} GOARCH=$${tgt#*-} go build -ldflags "$(VFLAG)" -o build/$(NAME)/$(NAME) .
	chmod +x build/$(NAME)/$(NAME)
	cp readme.md build/$(NAME)/
	tar -zcf $@ -C build ./$(NAME)
	rm -r build/$(NAME)

depend:
	go get -v $(DEPEND)
	glide install

lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	for pkg in $$(go list ./... |grep -v /vendor/); do go vet $$pkg; done

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

	
