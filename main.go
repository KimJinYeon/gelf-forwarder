package main

import (
	"log"
	"net"
	"strconv"

	"gelf-forwarder/internal"
)

func main() {

	config, err := internal.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Set up UDP listener for GELF messages
	inboundAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort("", strconv.Itoa(config.InboundPort)))
	if err != nil {
		log.Fatalf("Failed to resolve UDP address: %v", err)
	}

	conn, err := net.ListenUDP("udp", inboundAddr)
	if err != nil {
		log.Fatalf("Failed to listen on UDP: %v", err)
	}
	defer conn.Close()

	// Set up UDP connection for forwarding
	destAddr, err := net.ResolveUDPAddr("udp", net.JoinHostPort(config.OutboundHost, strconv.Itoa(config.OutboundPort)))
	if err != nil {
		log.Fatalf("Failed to resolve destination address: %v", err)
	}
	destConn, err := net.DialUDP("udp", nil, destAddr)
	if err != nil {
		log.Fatalf("Failed to connect to destination: %v", err)
	}
	defer destConn.Close()

	buf := make([]byte, 65535)
	for {
		// Read incoming message
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading UDP: %v", err)
			continue
		}

		// Decompress message
		decompressed, err := internal.Decompress(buf[:n])
		if err != nil {
			log.Printf("Failed to decompress message: %v", err)
			continue
		}

		otlpMessage, err := internal.TransformToOTLP(decompressed)
		if err != nil {
			log.Printf("Failed to transform GELF to OTLP: %v", err)
			continue
		}

		// Forward message
		if err := internal.ForwardMessage(otlpMessage, destConn); err != nil {
			log.Printf("Failed to forward message: %v", err)
		}

	}
}
