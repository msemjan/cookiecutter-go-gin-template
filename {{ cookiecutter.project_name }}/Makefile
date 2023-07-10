BINARY_NAME=main.out

all: test build

build:
	go build -o ${BINARY_NAME} .

test:
	go test

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

deploy:
	@echo "Error: Not implemented yet"
