package client

import (
	"fmt"
	"io"
)

const (
	byteVarIntShift = 7
	maxVarIntShift  = 35
)

// VarInt is a 32 bit integer with variable encoding
type VarInt int32

// EncodeBinary encodes the VarInt in its binary format
func (vint VarInt) EncodeBinary(writer io.ByteWriter) error {
	value := uint32(vint)
	var temp byte
	for {
		temp = (byte)(value & 0b01111111)
		value >>= byteVarIntShift
		if value != 0 {
			temp |= 0b10000000
			writer.WriteByte(temp)
		} else {
			writer.WriteByte(temp)
			return nil
		}
	}
}

// DecodeBinary decodes a VarInt from its binary format
func (vint *VarInt) DecodeBinary(reader io.ByteReader) error {
	output := VarInt(0)
	var shift uint8
	for {
		readByte, err := reader.ReadByte()
		if err != nil {
			return err
		}
		output |= (VarInt(0b01111111 & readByte)) << shift
		if readByte&0b10000000 == 0 {
			break
		}
		shift += byteVarIntShift
		if shift == maxVarIntShift {
			return fmt.Errorf("VarInt exceeded maximum length")
		}
	}
	*vint = output
	return nil
}
