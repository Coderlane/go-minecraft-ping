package client

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecodeFullPacket(t *testing.T) {
	var buf bytes.Buffer

	input := Packet{
		ID:   1,
		Data: []byte("test"),
	}

	err := input.EncodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	var output Packet
	err = output.DecodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(input, output) {
		t.Errorf("Expected: %v Got: %v", input, output)
	}
}
