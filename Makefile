.PHONY: gen build
BINARY_NAME ?= url_shortener
BINARY_PATH ?= bin
COVERAGE_FILE ?= coverage.txt
PORT=8080

check:
	go vet ./...
	golangci-lint run

gen_mocks:
	minimock -g -i ./service/database.* -o ./service/database -s _mock.go
	minimock -g -i ./service/domain.* -o ./service/domain -s _mock.go

precompile:
	go build ./...

compile:
	go build -o bin/binary cmd/${BINARY_NAME}/main.go

compile_linux:
	GOOS=linux GOARCH=amd64 go build -v -o bin/binary cmd/${BINARY_NAME}/main.go

short_test:
	go test -v -short -race -coverprofile=${COVERAGE_FILE} -covermode=atomic ./...

test:
	go test -v -race -coverprofile=${COVERAGE_FILE} -covermode=atomic ./...

test_local:
	set -o allexport &&	source .env && \
	export DB_HOST=127.0.0.1 && \
	go test -v -race -coverprofile=${COVERAGE_FILE} -covermode=atomic ./...

docker: compile_linux
	docker build --no-cache -t ${BINARY_NAME} .

docker-compose: docker
	docker-compose up -d db
	sleep 3
	docker-compose up db_migrator
	docker-compose up service



