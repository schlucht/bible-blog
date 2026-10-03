BINARY_NAME=bibleblog
PORT=5100
MAIN=./cmd/web

install:
	@go mod tidy

docker:
	@docker compose up -d

clean: stop
	@go clean
	@rm -R ./dist

build:
	@go build -o dist/${BINARY_NAME} ${MAIN}

start: build

stop: 
	@-pkill -f ${BINARY_NAME}
	@echo "Backend stopped..."
	

restart: stop build
	clear

run: start
	@env ./dist/${BINARY_NAME} -port=${PORT} &
	@echo "Backend running..."





