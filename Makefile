.DEFAULT_GOAL := build

.PHONY:fmt vet build clean ex
fmt:
	cd ${CHEX} && go fmt ./...

vet: fmt
	cd ${CHEX} && go vet ./...

build: vet
	cd ${CHEX} && go build

run: build
	cd ${CHEX} && go run main.go

clean:
	cd ${CHEX} && go clean

ex:
	mkdir -p ${CHEX}
	cd ${CHEX} && go mod init ${CHEX}
	echo 'package main\n\nfunc main() {\n\n}\n' >${CHEX}/main.go

play:
	open http://go.dev/play
