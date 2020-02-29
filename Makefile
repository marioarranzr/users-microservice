# Makefile for users-microservice

# Variables
UNAME		:= $(shell uname -s)

.PHONY: help
help:
ifeq ($(UNAME), Linux)
	@grep -P '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
else
	@# this is not tested, but prepared in advance for you, Mac drivers
	@awk -F ':.*###' '$$0 ~ FS {printf "%15s%s\n", $$1 ":", $$2}' \
		$(MAKEFILE_LIST) | grep -v '@awk' | sort
endif

: ### ---- Usage ---- :

.PHONY: build
build: clean ### Clean previous build and build the app
	@go build -i -v github.com/marioarranzr/users-microservice

clean: ### Remove previous build
	@rm -f users-microservice

# Targets
#
.PHONY: test
test:	### Run unit tests
	@go test -cover -v -coverprofile=coverage.out ./... 
	@go tool cover -func=coverage.out                                                                      

.PHONY: run
run:	### Run locally in debug mode
	@go run main.go                                                                       

.PHONY: dockerize
dockerize: build ### Build locally in docker
	@docker build -t users-microservice:latest .

.PHONY: run-locally
run-locally: dockerize ### Build and run locally in docker
	@docker run -p 9091:9091 --rm -it users-microservice
