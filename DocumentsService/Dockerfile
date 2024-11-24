# Use the official Go image as the base image for building the DocumentsServicelication
FROM golang:1.23.3 AS builder

WORKDIR /DocumentsService
COPY ./DocumentsService/go.mod ./DocumentsService/go.sum ./DocumentsService/main ./
RUN go mod download
RUN go mod init DocumentsService
# RUN go mod tidy
RUN go get github.com/gorilla/mux
RUN go build -o main .
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

FROM alpine:latest

# RUN apk --no-cache add ca-certificates
WORKDIR /root/

# COPY --from=builder ./DocumentsService/main .

EXPOSE 7000

CMD ["./main"]
