package client

// Type represents a minecraft message type
type Type int32

// Message represents a minecraft message
type Message interface {
	Type() Type
}

// Handshake represents the opening handshake with the server
type Handshake struct {
	Version int32 `rcon:"variable"`
	Address string
	Port    uint16
	State   int32 `rcon:"variable"`
}

func findTag(tags []string, tag string) bool {
	for _, knownTag := range tags {
		if knownTag == tag {
			return true
		}
	}
	return false
}
