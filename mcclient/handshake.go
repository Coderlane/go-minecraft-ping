package mcclient

import (
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
	var err error
	if err = hnd.Version.EncodeBinary(writer); err != nil {
		return err
	}
	if err = hnd.Address.EncodeBinary(writer); err != nil {
		return err
	}
	if err = binary.Write(writer, binary.BigEndian, hnd.Port); err != nil {
		return err
	}
	if err = hnd.State.EncodeBinary(writer); err != nil {
		return err
	}
	return nil
}

// DecodeBinary decodes the handshake from the binary format on the wire
func (hnd *Handshake) DecodeBinary(reader io.Reader) error {
	var err error
	if err = hnd.Version.DecodeBinary(reader); err != nil {
		return err
	}
	if err = hnd.Address.DecodeBinary(reader); err != nil {
		return err
	}
	if err = binary.Read(reader, binary.BigEndian, &hnd.Port); err != nil {
		return err
	}
	if err := hnd.State.DecodeBinary(reader); err != nil {
		return err
	}
	return nil
}
