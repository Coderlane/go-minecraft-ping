package client

import (
	"encoding/binary"
	"io"
)

// VarString is a string prefixed with a VarInt
type VarString string

// Length represents the encoded length of the string
func (vstr VarString) Length() int {
	length := VarInt(len(vstr))
	return length.Length() + len(vstr)
}

// EncodeBinary encodes the VarString in its binary format
func (vstr VarString) EncodeBinary(writer io.Writer) error {
	length := VarInt(len(vstr))
	if err := length.EncodeBinary(writer); err != nil {
		return err
	}
	return binary.Write(writer, binary.BigEndian, []byte(vstr))
}

// DecodeBinary decodes a VarString from its binary format
func (vstr *VarString) DecodeBinary(reader io.Reader) error {
	var length VarInt
	if err := length.DecodeBinary(reader); err != nil {
		return err
	}
	strBytes := make([]byte, length)
	if err := binary.Read(reader, binary.BigEndian, strBytes); err != nil {
		return err
	}
	*vstr = VarString(strBytes)
	return nil
}
