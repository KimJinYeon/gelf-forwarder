package internal

import (
	"encoding/json"
	"net"
	"time"
)

// ForwardMessage processes and forwards the GELF message to the destination.
func ForwardMessage(data []byte, destConn *net.UDPConn) error {
	// Parse JSON
	var gelfMessage map[string]interface{}
	if err := json.Unmarshal(data, &gelfMessage); err != nil {
		return err
	}

	// Modify GELF message
	if ts, ok := gelfMessage["timestamp"].(float64); ok {
		gelfMessage["timestamp"] = time.Unix(int64(ts), 0).Format(time.RFC3339)
	}
	gelfMessage["message"] = gelfMessage["short_message"]
	delete(gelfMessage, "short_message")

	// Serialize to JSON
	updatedData, err := json.Marshal(gelfMessage)
	if err != nil {
		return err
	}

	// Send to destination
	_, err = destConn.Write(updatedData)
	return err
}
