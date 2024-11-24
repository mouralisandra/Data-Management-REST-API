# Use the official Go image as the base image for building the CustomerDataServicelication
FROM golang:1.23.3 AS builder

WORKDIR /CustomerDataService
COPY ./CustomerDataService/go.mod ./CustomerDataService/go.sum ./CustomerDataService/main ./
RUN go mod init CustomerDataService
RUN go mod download
RUN go mod tidy
RUN go get github.com/gorilla/mux
RUN go build -o main .
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# RUN go build -o main .

FROM alpine:latest

# RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY ./CustomerDataService/main .

EXPOSE 8000

CMD ["./main"]
