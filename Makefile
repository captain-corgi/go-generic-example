all: build run clean

build:
	@gotip build -o generic-example cmd/generic-example/main.go
run:
	@./generic-example
clean:
	@rm generic-example

tidy:
	gotip mod tidy

.PHONY: vendor
vendor:
	gotip mod vendor
clean-vendor:
	rm -rf vendor