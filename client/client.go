package client

import (
	"net"
)

// Client handles low level communication with Minecraft servers
type Client interface {
	Send(Packet) error
	Recv() (*Packet, error)
	Close() error
}

type client struct {
	conn net.Conn
}

// NewClient creates a new client connection
func NewClient(address string) (Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &client{
		conn: conn,
	}, nil
}

// Send a Packet on the connection
func (cln *client) Send(pkt Packet) error {
	return pkt.EncodeBinary(cln.conn)
}

// Recv a Packet on the connection
func (cln *client) Recv() (*Packet, error) {
	var pkt Packet
	if err := pkt.DecodeBinary(cln.conn); err != nil {
		return nil, err
	}
	return &pkt, nil
}

// Close the connection
func (cln *client) Close() error {
	return cln.conn.Close()
}
