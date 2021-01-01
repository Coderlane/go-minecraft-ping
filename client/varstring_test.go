package client

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestEncodeDecodeVarString(t *testing.T) {
	type testCase struct {
		Case     VarString
		Expected []byte
	}

	tcases := []testCase{
		{"test", []byte{4, 116, 101, 115, 116}},
		{"", []byte{0}},
	}

	for _, tcase := range tcases {
		t.Run(fmt.Sprintf("%s", tcase.Case), func(t *testing.T) {
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

			var output VarString
			err = output.DecodeBinary(&buf)
			if err != nil {
				t.Fatal(err)
			}
			if tcase.Case != output {
				t.Errorf("Expected: %s Got: %s", tcase.Case, output)
			}
		})
	}
}
