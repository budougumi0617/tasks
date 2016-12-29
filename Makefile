NAME     := tasks
VERSION  := v0.0.1
REVISION := $(shell git rev-parse --short HEAD)

SRCS    := $(shell find . -type f -name '*.go')
LDFLAGS := -ldflags="-s -w -X \"github.com/budougumi0617/tasks/version.Version=$(VERSION)\" -X \"github.com/budougumi0617/tasks/version.Revision=$(REVISION)\" -extldflags \"-static\""

DIST_DIRS := find * -type d -exec

bin/$(NAME): $(SRCS)
	go build $(LDFLAGS) -o bin/$(NAME)

.PHONY: glide
glide:
ifeq ($(shell command -v glide 2> /dev/null),)
	curl https://glide.sh/get | sh
endif

.PHONY: deps
deps: glide
	glide install

.PHONY: install
install:
	go install $(LDFLAGS)

.PHONY: clean
clean:
	rm -rf bin/*
	rm -rf vendor/*

.PHONY: test
test:
	go test -cover -v `glide novendor`

.PHONY: ci-test
ci-test:
	echo "" > coverage.txt
	for d in `glide novendor`; do \
		go test -coverprofile=profile.out -covermode=atomic -v $$d; \
		if [ $$? != 0 ]; then \
			exit 2; \
		else \
			if [ -f profile.out ]; then \
				cat profile.out >> coverage.txt; \
				rm profile.out; \
			fi; \
		fi; \
	done

.PHONY: cross-build
cross-build: deps
	for os in darwin linux windows; do \
		for arch in amd64 386; do \
			GOOS=$$os GOARCH=$$arch CGO_ENABLED=0 go build $(LDFLAGS) -o dist/$$os-$$arch/$(NAME); \
		done; \
	done

.PHONY: dist
dist:
	cd dist && \
	$(DIST_DIRS) cp ../LICENSE {} \; && \
	$(DIST_DIRS) cp ../README.md {} \; && \
	$(DIST_DIRS) tar -zcf $(NAME)-$(VERSION)-{}.tar.gz {} \; && \
	$(DIST_DIRS) zip -r $(NAME)-$(VERSION)-{}.zip {} \; && \
	cd ..


