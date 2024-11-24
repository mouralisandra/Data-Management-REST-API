# Étape de construction
FROM golang:1.23.3 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -o main .

# Étape finale
FROM alpine:3.18
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
