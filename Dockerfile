FROM golang:1.22 AS builder

WORKDIR /app

RUN apt-get update && apt-get install -y git \
    && go install github.com/swaggo/swag/cmd/swag@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN swag init

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o receipt-processor

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/receipt-processor .
COPY --from=builder /app/docs/swagger.json ./docs/swagger.json

EXPOSE 8080

CMD ["./receipt-processor"]
