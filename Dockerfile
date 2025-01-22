# 1. Build 단계: Go 애플리케이션 빌드
FROM golang:1.23 AS builder

#ENV GOPROXY=https://build-proxy.azdevops.kt.co.kr/repository/go-group

WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o gelf-otlp-forwarder .

# 2. Runtime 단계: 경량 이미지 사용
FROM scratch

WORKDIR /app

# Copy the static Go binary
COPY --from=builder /app/gelf-otlp-forwarder .

# Default port for GELF listener
EXPOSE 5044

# Run the application
ENTRYPOINT ["/app/gelf-forwarder"]