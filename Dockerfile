FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev sqlite-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main ./cmd/testSolution/main.go

# Финальный образ
FROM alpine:latest

RUN apk --no-cache add ca-certificates sqlite-libs tzdata

WORKDIR /root

COPY --from=builder /app/main .

RUN mkdir -p /data

EXPOSE 8080

CMD ["./main"]
