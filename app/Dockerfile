# Use the official Go image as the base image for building the application
FROM golang:1.23.3 AS builder

WORKDIR /app
COPY ./go.mod ./go.sum ./main ./
RUN go mod download
RUN go mod tidy
RUN go get github.com/gorilla/mux
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder ./app/main .

EXPOSE 8000

CMD ["./main"]
