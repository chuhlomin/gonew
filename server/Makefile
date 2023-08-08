.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test
## test: Runs go test
test:
	@go test -cover ./...

.PHONY: lint
## lint: Runs golanglint-ci
lint:
	@golangci-lint run

.PHONY: build
## build: Runs go build (builds `main` binary)
build:
	@go build -o main .

.PHONY: run
## run: Run app
run: build
	@./main

.PHONY: clean
## clean: Clean up
clean:
	@rm main
