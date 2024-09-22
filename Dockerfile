# Build
FROM golang:alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o quickshare-backend ./cmd/api-server

# Run
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/quickshare-backend .
COPY .aws /root/.aws

EXPOSE 8080

CMD ["./quickshare-backend"]
