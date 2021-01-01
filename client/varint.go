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

// Length represents the encoded length of the integer
func (vint VarInt) Length() int {
	value := uint32(vint)
	length := 0
	for {
		value >>= byteVarIntShift
		length += 1
		if value == 0 {
			return length
		}
	}
}

// EncodeBinary encodes the VarInt in its binary format
func (vint VarInt) EncodeBinary(writer io.Writer) error {
	value := uint32(vint)
	temp := make([]byte, 1)
	for {
		temp[0] = (byte)(value & 0b01111111)
		value >>= byteVarIntShift
		if value != 0 {
			temp[0] |= 0b10000000
			writer.Write(temp)
		} else {
			writer.Write(temp)
			return nil
		}
	}
}

// DecodeBinary decodes a VarInt from its binary format
func (vint *VarInt) DecodeBinary(reader io.Reader) error {
	output := VarInt(0)
	var shift uint8
	readByte := make([]byte, 1)
	for {
		_, err := reader.Read(readByte)
		if err != nil {
			return err
		}
		output |= (VarInt(0b01111111 & readByte[0])) << shift
		if readByte[0]&0b10000000 == 0 {
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
