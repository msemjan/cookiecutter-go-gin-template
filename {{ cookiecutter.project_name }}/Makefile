BINARY_NAME={{ cookiecutter.binary_name }}

all: test build

bin:
	mkdir ./bin

build: bin
	go build -o ./bin/${BINARY_NAME} main.go

compile:
	GOOS=freebsd GOARCH=amd64 go build -o bin/freebsd-amd64-${BINARY_NAME} main.go
	GOOS=linux   GOARCH=amd64 go build -o bin/linux-amd64-${BINARY_NAME}   main.go
	GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64-${BINARY_NAME} main.go
	GOOS=darwin  GOARCH=amd64 go build -o bin/darwin-amd64-${BINARY_NAME}  main.go

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out

deps:
	go mod download
	go install golang.org/x/tools/cmd/goimports@latest

vet:
	go vet

fmt:
	goimports

run: build
	./bin/${BINARY_NAME}

dockerBuild:
	docker build .

clean:
	go clean
	rm ./bin/freebsd-amd64-${BINARY_NAME}
	rm ./bin/linux-amd64-${BINARY_NAME}
	rm ./bin/windows-amd64-${BINARY_NAME}
	rm ./bin/darwin-amd64-${BINARY_NAME}

deploy:
	@echo "Error: Not implemented yet"
