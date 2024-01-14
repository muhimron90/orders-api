BINARY_APP=api
MAIN_FILE=main.go

build:
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_APP} ${MAIN_FILE}

run: build
	./${BINARY_APP}

dev:
	SERVER_PORT=3000 go run ${MAIN_FILE}

clean:
	go clean
	rm ${BINARY_APP}

seed:
	py scripts/seed-orders.py

.PHONY: build run dev clean seed
