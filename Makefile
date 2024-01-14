BINARY_APP=api

build:
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_APP} main.go

run: build
	./${BINARY_APP}

clean:
	go clean
	rm ${BINARY_APP}

.PHONY: build run clean
