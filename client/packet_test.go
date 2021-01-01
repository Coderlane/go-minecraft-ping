package client

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeDecodeFullHandshake(t *testing.T) {
	var buf bytes.Buffer

	input := Handshake{
		Version: -1,
		Address: "test",
		Port:    25565,
		State:   1,
	}

	err := input.EncodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	var output Handshake
	err = output.DecodeBinary(&buf)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(input, output) {
		t.Errorf("Expected: %v Got: %v", input, output)
	}
}
