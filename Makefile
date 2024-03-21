CMD = ./cmd
BIN = ./bin/stamper

clean:
	rm -rf ./bin
	rm -rf ./cmd/license*.go

gen:
	go run ./codegen

test:
	go test ${CMD}

build:
	go build -o ${BIN} ${CMD}

dev:
	go run ${CMD}
