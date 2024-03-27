FROM golang:1.22 AS build
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/main.go
EXPOSE 8080
CMD ["app"]
