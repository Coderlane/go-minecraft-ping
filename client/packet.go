package client

import (
	"bufio"
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
