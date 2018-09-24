BRANCH_NAME ?= $(shell git rev-parse --abbrev-ref HEAD)
FORCE: # https://www.gnu.org/software/make/manual/html_node/Force-Targets.html#Force-Targets

bin/%: export GOOS := linux
bin/%: FORCE # forcing build to run everytime
	$(info ==> Building '$@' binary)
	@go build -o $@ cmd/$(@F)/*.go
	@du -h $@

.PHONY: build
build: branch bin/tick bin/fw

.PHONY: branch
branch:
	$(info ==> Branch '$(BRANCH_NAME)')

deploy: build
	sls deploy -v