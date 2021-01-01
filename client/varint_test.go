package client

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeDecodeVarInt(t *testing.T) {
	type testCase struct {
		Case     VarInt
		Expected []byte
	}

	tcases := []testCase{
		{0, []byte{0}},
		{1, []byte{1}},
		{255, []byte{255, 1}},
		{-1, []byte{255, 255, 255, 255, 15}},
	}

	for _, tcase := range tcases {
		t.Run(fmt.Sprintf("%d", tcase.Case), func(t *testing.T) {
			var buf bytes.Buffer

			input := tcase.Case
			err := input.EncodeBinary(&buf)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(tcase.Expected, buf.Bytes()) {
				t.Errorf("Expected: %v Got: %v", tcase.Expected, buf.Bytes())
			}
			if len(tcase.Expected) != input.Length() {
				t.Errorf("Expected: %v Got: %v", len(tcase.Expected), input.Length())
			}

			var output VarInt
			err = output.DecodeBinary(&buf)
			if err != nil {
				t.Fatal(err)
			}
			if tcase.Case != output {
				t.Errorf("Expected: %d Got: %d", tcase.Case, output)
			}
		})
	}
}

func TestDecodeTooLong(t *testing.T) {
	buf := bytes.NewBuffer([]byte{255, 255, 255, 255, 255, 255})
	var output VarInt
	err := output.DecodeBinary(buf)
	if err == nil {
		t.Fatalf("Expected an error.")
	}
}

func TestDecodeTooShort(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	var output VarInt
	err := output.DecodeBinary(buf)
	if err == nil {
		t.Fatalf("Expected an error.")
	}
}
