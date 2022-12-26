.PHONY:
.SILENT:

build:
	go build -o ./.bin/weather cmd/main.go

run: build
	./.bin/weather