.PHONY: help
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: build
build: ## Build a wasm image from a TinyGo filter.
	tinygo build -o charlescd.wasm -scheduler=none -target=wasi ./cmd/charlescd

.PHONY: test
test: build ## Test the wasm image from the end user’s experience.
	cp charlescd.wasm e2e/tests/integration/request-routing/deploy/charlescd.wasm
	cd e2e; \
	docker build --build-arg html=webpage/red.html -f webpage/Dockerfile -t webpage:red .; \
	docker build --build-arg html=webpage/blue.html -f webpage/Dockerfile -t webpage:blue .; \
	kubectl kuttl test
