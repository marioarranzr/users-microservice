FROM golang

ENV SRC_DIR=/go/src/github.com/marioarranzr/users-microservice

WORKDIR $SRC_DIR
COPY go.mod go.sum ./

ADD . $SRC_DIR
RUN mkdir /app
RUN go build -o ./bin/users-microservice -v main.go; cp ./bin/users-microservice /app/users-microservice

WORKDIR /app/

ENTRYPOINT ["sh", "-c", "/app/users-microservice"]
