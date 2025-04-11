FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-X 'main.version=1.0.0'" -o godocs .

FROM alpine:3.18
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/godocs /usr/local/bin/
ENTRYPOINT ["godocs"]