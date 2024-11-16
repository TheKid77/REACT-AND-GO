.DEFAULT_GOAL := run

.PHONY:fmt vet build
fmt:	
	go fmt ./...
vet: fmt
	go vet ./...
build: vet	
	go build
run: build
	go run .