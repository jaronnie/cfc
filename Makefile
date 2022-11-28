.PHONY: build
build:
	@cd cmd/cfctl; go install

.PHONY: fmt.all
fmt.all:
	@sh scripts/goimports.sh all