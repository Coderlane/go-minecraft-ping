package client

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
)

// Packet encapulsates all messages on the wire
type Packet struct {
	ID   VarInt
	Data []byte
}

// EncodeBinary encodes the packet in the binary format for the wire
func (pkt Packet) EncodeBinary(writer io.Writer) error {
	var err error
	buf := bufio.NewWriter(writer)
	length := VarInt(pkt.ID.Length() + len(pkt.Data))
	if err = length.EncodeBinary(buf); err != nil {
		return err
	}
	if err = pkt.ID.EncodeBinary(buf); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, pkt.Data); err != nil {
		return err
	}
	return buf.Flush()
}

// DecodeBinary decodes the packet from the binary format on the wire
func (pkt *Packet) DecodeBinary(reader io.Reader) error {
	var (
		err    error
		length VarInt
	)
	buf := bufio.NewReader(reader)
	if err = length.DecodeBinary(buf); err != nil {
		return err
	}
	if err = pkt.ID.DecodeBinary(buf); err != nil {
		return err
	}
	length -= VarInt(pkt.ID.Length())
	data := make([]byte, length)
	if err = binary.Read(buf, binary.BigEndian, data); err != nil {
		return err
	}
	pkt.Data = data
	return nil
}

// Handshake represents the opening handshake with the server
type Handshake struct {
	Version VarInt
	Address VarString
	Port    uint16
	State   VarInt
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
	pkt := Packet{
		ID:   0,
		Data: buf.Bytes(),
	}
	return pkt.EncodeBinary(writer)
}

// DecodeBinary decodes the handshake from the binary format on the wire
func (hnd *Handshake) DecodeBinary(reader io.Reader) error {
	var (
		err error
		pkt Packet
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
