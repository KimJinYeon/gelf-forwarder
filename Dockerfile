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
RUN go build -o gelf-forwarder .

# 2. Runtime 단계: 경량 이미지 사용
FROM golang:1.23

WORKDIR /app

# Copy Go binary from build stage
COPY --from=builder /app/gelf-forwarder .

# Default port for GELF listener
EXPOSE 5044

# Run the application
ENTRYPOINT ["./gelf-forwarder"]