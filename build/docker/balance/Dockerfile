FROM golang:1.16

WORKDIR /usr/src/balance

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go test ./... -cover
WORKDIR /usr/src/balance/cmd/balance
RUN go build -v -o /usr/local/bin/balance

CMD ["balance"]
