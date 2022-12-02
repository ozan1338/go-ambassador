FROM golang:1.18

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod download github.com/onsi/ginkgo

COPY . .
