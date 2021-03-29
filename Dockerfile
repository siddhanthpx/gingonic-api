FROM golang:latest

LABEL maintainer="Siddhanth <github.com/siddhanthpx>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

CMD ["./gingonic-api"]