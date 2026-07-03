FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY server/ ./server/
RUN go build -o mediago-webui server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /app/mediago-webui .
COPY web/ ./web/

EXPOSE 8080
CMD ["./mediago-webui"]
