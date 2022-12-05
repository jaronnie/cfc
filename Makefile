.PHONY: build
build:
	@go install github.com/jaronnie/cfc/cmd/cfctl

.PHONY: fmt.all
fmt.all:
	@sh scripts/goimports.sh all