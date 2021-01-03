package mcclient

//go:generate mockgen -destination=mock_minecraft_client.go -package=mcclient -self_package=github.com/Coderlane/go-minecraft-ping/mcclient github.com/Coderlane/go-minecraft-ping/mcclient MinecraftClient

import (
	"bytes"
	"encoding/json"
	"net"
	"strconv"

	"github.com/Coderlane/go-minecraft-ping/client"
)

// ClientState represents the connection state of the client
type ClientState int

const (
	// ClientStateUnknown is the default state
	ClientStateUnknown = 0
	// ClientStateStatus is for sending unauthenticated ping and status messages
	ClientStateStatus = 1
	// ClientStateLogin is for authenticating with the client
	ClientStateLogin = 2
)

// MinecraftClient represents a client connection to a Minecraft server
type MinecraftClient interface {
	Handshake(ClientState) error
	Status() (*StatusResponse, error)
	Close() error
}

type mcclient struct {
	conn  client.Client
	host  string
	port  int
	state ClientState
}

// NewMinecraftClient wraps an existing client connection as a MinecraftClient
func NewMinecraftClient(client client.Client) (MinecraftClient, error) {
	host, strPort, err := net.SplitHostPort(client.Addr())
	if err != nil {
		return nil, err
	}
	port, err := strconv.Atoi(strPort)
	if err != nil {
		return nil, err
	}
	return &mcclient{
		conn: client,
		host: host,
		port: port,
	}, nil
}

// NewMinecraftClientFromAddress creates a new client connection with a
// Minecraft server
func NewMinecraftClientFromAddress(address string) (MinecraftClient, error) {
	client, err := client.NewClient(address)
	if err != nil {
		return nil, err
	}
	return NewMinecraftClient(client)
}

// Handshake performs the opening handshake with the server. The state
// indicates what the next intended state is.
func (cln *mcclient) Handshake(state ClientState) error {
	hnd := Handshake{
		Version: -1,
		Address: client.VarString(cln.host),
		Port:    uint16(cln.port),
		State:   client.VarInt(state),
	}
	var buf bytes.Buffer
	if err := hnd.EncodeBinary(&buf); err != nil {
		return err
	}
	pkt := client.Packet{
		ID:   0,
		Data: buf.Bytes(),
	}
	return cln.conn.Send(pkt)
}

// Status requests the status from the minecraft server
func (cln *mcclient) Status() (*StatusResponse, error) {
	request := client.Packet{
		ID: 0,
	}
	if err := cln.conn.Send(request); err != nil {
		return nil, err
	}
	response, err := cln.conn.Recv()
	if err != nil {
		return nil, err
	}
	buf := bytes.NewReader(response.Data)
	var data client.VarString
	if err := data.DecodeBinary(buf); err != nil {
		return nil, err
	}
	var status StatusResponse
	err = json.Unmarshal([]byte(data), &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

// Close closes the connection with the minecraft server
func (cln *mcclient) Close() error {
	return cln.conn.Close()
}
