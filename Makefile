APP_NAME=PFBackend

build-dev:
	go build -o ${APP_NAME} ./cmd/main.go

build:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ${APP_NAME} ./cmd/main.go

start:
	./${APP_NAME}

restart: build start

dev:
	air -c ./.air.toml

run: start

