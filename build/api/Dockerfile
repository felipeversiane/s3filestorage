FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /app/api ./cmd/main.go

FROM scratch

COPY --from=builder /app/api /api

ENTRYPOINT ["/api"]
