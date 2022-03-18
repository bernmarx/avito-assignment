FROM golang:1.16

WORKDIR /usr/src/microservice

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test ./... -cover
WORKDIR /usr/src/microservice/cmd/microservice
RUN go build -v -o /usr/local/bin/microservice

CMD ["microservice"]
