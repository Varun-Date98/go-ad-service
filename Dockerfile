# ---- build stage ----
FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/go-ad-server ./go-ad-server

# ---- runtime stage ----
FROM alpine:3.20

WORKDIR /app
RUN apk add --no-cache ca-certificates

COPY --from=builder /bin/go-ad-server /usr/local/bin/go-ad-server

EXPOSE 8080

CMD ["go-ad-server"]
