FROM golang:latest AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/main /usr/local/bin/
COPY ./public /app/
COPY config.yaml /app/
WORKDIR /app
EXPOSE 8080
CMD ["main"]