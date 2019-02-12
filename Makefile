#!make
include .makerc
export $(shell sed 's/=.*//' .makerc)

NAME=elsa-cli
VERSION=0.0.1
BUILD_DATE=$(shell date)
BUILD_PATH=$(shell pwd)/build
OUTPUT=$(BUILD_PATH)/$(NAME)-$(VERSION)

version:
	@echo "Version: $(VERSION) - $(BUILD_DATE)"

deps:
	go get -u github.com/gobuffalo/packr/...
	go get -u github.com/spf13/cobra/cobra
	go get github.com/spf13/viper
	go get github.com/matishsiao/goInfo
	go get github.com/acobaugh/osrelease
	go get golang.org/x/crypto/acme/autocert
	go get github.com/c4milo/github-release
	go get github.com/mitchellh/gox

clean:
	@echo "Clearing build directory for previous build with current version"
	@rm -rf $(OUTPUT)-*

build: clean
	@echo "Building version $(VERSION)\nBuild date: $(BUILD_DATE)"
	@$(shell perl -pi -e 's#(.*VERSION.*=\x20)(.*)#$${1}"$(VERSION)"#' main.go)
	@$(shell perl -pi -e 's#(.*BUILD_DATE.*=\x20)(.*)#$${1}"$(BUILD_DATE)"#' main.go)

	@gox -ldflags "-X main.VERSION=$(VERSION)" \
		-osarch="linux/amd64" -osarch="linux/386" \
		-output "$(OUTPUT)-{{.OS}}-{{.Arch}}" \
		./...

dist: build
	$(eval FILES := $(shell ls build))
	@rm -rf dist && mkdir -p dist
	@for f in $(FILES); do \
		(cd $(shell pwd)/build/ && tar -cvzf ../dist/$$f.tar.gz $$f); \
		(cd $(shell pwd)/dist/ && shasum -a 512 $$f.tar.gz > $$f.sha512); \
	done

install:
	@echo "Installing $(NAME) to system"
	@install -d -m 755 '/usr/local/bin/'

    ifeq ($(shell uname -p), x86_64)
		install $(OUTPUT)-x64 '/usr/local/bin/ecp'
	else
		install $(OUTPUT)-x86 '/usr/local/bin/ecp'
    endif

release: dist remove-current-release
	@echo "Releasing v$(VERSION) to Github"
	#git tag -s v$(VERSION)
	@latest_tag=$$(git describe --tags `git rev-list --tags --max-count=1`); \
	comparison="$$latest_tag..HEAD"; \
	if [ -z "$$latest_tag" ]; then comparison=""; fi; \
	changelog=$$(git log $$comparison --oneline --no-merges); \
	github-release elsacp/$(NAME) v$(VERSION) "$$(git rev-parse --abbrev-ref HEAD)" "**Changelog**<br/>$$changelog" 'dist/*'; \
	git pull

pre-release: dist remove-current-release
	@echo "Releasing v$(VERSION) to Github"
	#git tag -s v$(VERSION)
	@latest_tag=$$(git describe --tags `git rev-list --tags --max-count=1`); \
	comparison="$$latest_tag..HEAD"; \
	if [ -z "$$latest_tag" ]; then comparison=""; fi; \
	changelog=$$(git log $$comparison --oneline --no-merges); \
	env GITHUB_TOKEN=$(GITHUB_TOKEN) github-release -prerelease \
		$(GITHUB_USER)/$(GITHUB_REPO) v$(VERSION) \
		"$$(git rev-parse --abbrev-ref HEAD)" "**Changelog**<br/>$$changelog" 'dist/*'; \
	git pull

remove-current-release:
	@curl -u $(GITHUB_USER):$(GITHUB_TOKEN) -X DELETE \
	https://api.github.com/repos/$(GITHUB_USER)/$(GITHUB_REPO)/git/refs/tags/v$(VERSION)

commit:
	@git cz -s

.PHONY: version deps clean build dist install release pre-release remove-current-release commit
