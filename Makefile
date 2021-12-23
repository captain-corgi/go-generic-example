all: build run clean

build:
	@go build -o generic-example cmd/generic-example/main.go
run:
	@./generic-example
clean:
	@rm generic-example

tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor
clean-vendor:
	rm -rf vendor