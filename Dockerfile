FROM golang:1.22.4

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x waitForDB.sh

RUN go build -o final-project ./cmd/main.go

CMD ["./app"]