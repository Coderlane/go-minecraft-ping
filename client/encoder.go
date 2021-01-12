package client

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Encoder struct {
	writer io.Writer
}

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{
		writer: writer,
	}
}

func (enc *Encoder) Encode(msg Message) error {
	var buf bytes.Buffer
	msgValue := reflect.ValueOf(msg)
	msgType := reflect.TypeOf(msg)
	for fieldIter := 0; fieldIter < msgValue.NumField(); fieldIter++ {
		field := msgType.Field(fieldIter)
		tags := strings.Split(field.Tag.Get("rcon"), ",")
		if err := encodeSimpleValue(&buf, msgValue.Field(fieldIter), tags); err != nil {
			return err
		}
	}
	if err := VarInt(buf.Len()).EncodeBinary(enc.writer); err != nil {
		return err
	}
	if err := VarInt(msg.Type()).EncodeBinary(enc.writer); err != nil {
		return err
	}
	_, err := io.Copy(enc.writer, &buf)
	return err
}

func encodeSimpleValue(writer io.Writer, value reflect.Value, tags []string) error {
	switch value.Kind() {
	case reflect.String:
		if err := VarInt(value.Len()).EncodeBinary(writer); err != nil {
			return err
		}
		if err := binary.Write(writer, binary.BigEndian, []byte(value.String())); err != nil {
			return err
		}
	case reflect.Bool:
		if err := binary.Write(writer, binary.BigEndian, value.Bool()); err != nil {
			return err
		}
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if findTag(tags, "variable") {
			if err := VarInt(value.Int()).EncodeBinary(writer); err != nil {
				return err
			}
		} else {
			if err := binary.Write(writer, binary.BigEndian, value.Interface()); err != nil {
				return err
			}
		}
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if err := binary.Write(writer, binary.BigEndian, value.Interface()); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unsupported kind: %s", value.Kind())
	}
	return nil
}
