package client

import (
	"bytes"
	"reflect"
	"testing"
)

type testStruct struct {
	S string
	B bool

	I8  int8
	I16 int16
	I32 int32
	I64 int64

	V32 int32 `rcon:"variable"`

	U8  uint8
	U16 uint16
	U32 uint32
	U64 uint64
}

func (ts testStruct) Type() Type {
	return 32
}

func TestEncode(t *testing.T) {
	var buf bytes.Buffer

	enc := NewEncoder(&buf)

	input := testStruct{
		S: "test",
		B: true,

		I8:  1,
		I16: 2,
		I32: 3,
		I64: 4,

		V32: 5,

		U8:  6,
		U16: 7,
		U32: 8,
		U64: 9,
	}
	err := enc.Encode(input)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(buf.Bytes())

	dec := NewDecoder(&buf)

	var output testStruct
	err = dec.Decode(&output)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", output)

	if !reflect.DeepEqual(input, output) {
		t.Errorf("Expected: %v Got: %v\n", input, output)
	}
}
