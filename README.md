# GELF OTLP Forwarder

# TL;DR

If you are looking for GELF Receiver for otel collector and can't find it, use this gelf-otlp-forwarder.

## Overview
GELF OTLP Forwarder is a lightweight intermediary server designed to bridge the gap between GELF (Graylog Extended Log Format) and OTLP (OpenTelemetry Protocol). Since the OpenTelemetry Collector does not provide a built-in GELF receiver, GELF Forwarder simplifies the process by receiving logs in GELF format over UDP, optionally decompressing them (GZIP), and forwarding the transformed data in OTLP format over UDP.

## Features
- **GELF to OTLP Transformation**: Automatically converts incoming GELF messages to OTLP format.
- **GZIP Decompression**: Handles GZIP-compressed GELF payloads.
- **Lightweight and Fast**: Built in Go, designed for high performance and low resource consumption.
- **UDP Input and Output**: Supports receiving and forwarding logs over UDP.

## Use Case
If you're using tools or services that emit logs in GELF format but need to integrate with OpenTelemetry for observability, GELF Forwarder provides a simple and efficient solution.

## Installation

### Pre-built Binaries
Download the pre-built binary for your platform from the [Releases](https://github.com/your-repo/gelf-forwarder/releases) page and extract it.

### Building from Source
Ensure you have Go installed (version 1.20 or later), then run:

```bash
git clone https://github.com/your-repo/gelf-otlp-forwarder.git
cd gelf-otlp-forwarder
go build -o gelf-otlp-forwarder .
```

This will create an executable named `gelf-otlp-forwarder` in the project directory.

## Usage

### Configuration
GELF Forwarder can be configured using a `config.yaml` file with the following structure:

```yaml
inbound_port: 5044
outbound_host: otel-collector
outbound_port: 12201
```

### Running the Server with Docker Compose
Create a `docker-compose.yaml` file with the following content:

```yaml
version: '3.8'
services:
  gelf-otlp-forwarder:
    image: gelf-otlp-forwarder:0.0.1
    container_name: go-gelf-forwarder
    environment:
      - CONFIG_PATH=/config/config.yaml
    ports:
      - "5044:5044/udp"
    volumes:
      - ./{your_path}/config.yaml:/config/config.yaml
    depends_on:
      - otel-collector
```

If you wish to set config path, set ENV `CONFIG_PATH` as you like.
This will replace default config path `/config/config.yaml`.

Start the service:

```bash
docker-compose up -d
```

### Example Configuration
`config.yaml` example:

```yaml
inbound_port: 5044
outbound_host: otel-collector
outbound_port: 12201
```

## Example
Send a GELF message using `netcat` for testing:

```bash
echo -n '{"version":"1.1","host":"example.org","short_message":"Test log"}' | \
  nc -u -w0 127.0.0.1 5044
```

The message will be received, transformed, and forwarded to the specified OTLP endpoint.

## Contact
For issues or feature requests, please open an issue here.

