package mcclient

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/Coderlane/go-minecraft-ping/client"
)

// Handshake represents the opening handshake with the server
type Handshake struct {
	Version client.VarInt
	Address client.VarString
	Port    uint16
	State   client.VarInt
}

// EncodeBinary encodes the handshake in the binary format for the wire
func (hnd Handshake) EncodeBinary(writer io.Writer) error {
	var (
		err error
		buf bytes.Buffer
	)
	if err = hnd.Version.EncodeBinary(&buf); err != nil {
		return err
	}
	if err = hnd.Address.EncodeBinary(&buf); err != nil {
		return err
	}
	if err = binary.Write(&buf, binary.BigEndian, hnd.Port); err != nil {
		return err
	}
	if err = hnd.State.EncodeBinary(&buf); err != nil {
		return err
	}
	pkt := client.Packet{
		ID:   0,
		Data: buf.Bytes(),
	}
	return pkt.EncodeBinary(writer)
}

// DecodeBinary decodes the handshake from the binary format on the wire
func (hnd *Handshake) DecodeBinary(reader io.Reader) error {
	var (
		err error
		pkt client.Packet
	)
	if err = pkt.DecodeBinary(reader); err != nil {
		return err
	}
	buf := bytes.NewReader(pkt.Data)
	if err = hnd.Version.DecodeBinary(buf); err != nil {
		return err
	}
	if err = hnd.Address.DecodeBinary(buf); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, &hnd.Port); err != nil {
		return err
	}
	if err := hnd.State.DecodeBinary(buf); err != nil {
		return err
	}
	return nil
}
