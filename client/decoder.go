package client

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Decoder struct {
	reader io.Reader
}

func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{
		reader: reader,
	}
}

func (dec *Decoder) Decode(msg Message) error {
	var (
		msgTypeID VarInt
		msgLen    VarInt
	)

	if err := msgTypeID.DecodeBinary(dec.reader); err != nil {
		return err
	}
	if err := msgLen.DecodeBinary(dec.reader); err != nil {
		return err
	}

	msgValue := reflect.Indirect(reflect.ValueOf(msg))
	msgType := msgValue.Type()
	for fieldIter := 0; fieldIter < msgValue.NumField(); fieldIter++ {
		field := msgType.Field(fieldIter)
		tags := strings.Split(field.Tag.Get("rcon"), ",")
		if err := decodeSimpleValue(dec.reader, msgValue.Field(fieldIter), tags); err != nil {
			return err
		}
	}
	return nil
}

func decodeSimpleValue(reader io.Reader, value reflect.Value, tags []string) (err error) {
	switch value.Kind() {
	case reflect.String:
		var strLen VarInt
		if err := strLen.DecodeBinary(reader); err != nil {
			return err
		}
		strBytes := make([]byte, strLen)
		err = binary.Read(reader, binary.BigEndian, strBytes)
		value.SetString(string(strBytes))
	case reflect.Bool:
		var valBool bool
		err = binary.Read(reader, binary.BigEndian, &valBool)
		value.SetBool(valBool)
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if findTag(tags, "variable") {
			var valInt VarInt
			if err := valInt.DecodeBinary(reader); err != nil {
				return err
			}
			value.SetInt(int64(valInt))
			return nil
		}
		switch value.Kind() {
		case reflect.Int8:
			var valInt int8
			err = binary.Read(reader, binary.BigEndian, &valInt)
			value.SetInt(int64(valInt))
		case reflect.Int16:
			var valInt int16
			err = binary.Read(reader, binary.BigEndian, &valInt)
			value.SetInt(int64(valInt))
		case reflect.Int32:
			var valInt int32
			err = binary.Read(reader, binary.BigEndian, &valInt)
			value.SetInt(int64(valInt))
		case reflect.Int64:
			var valInt int64
			err = binary.Read(reader, binary.BigEndian, &valInt)
			value.SetInt(valInt)
		}
	case reflect.Uint8:
		var valUint uint8
		err = binary.Read(reader, binary.BigEndian, &valUint)
		value.SetUint(uint64(valUint))
	case reflect.Uint16:
		var valUint uint16
		err = binary.Read(reader, binary.BigEndian, &valUint)
		value.SetUint(uint64(valUint))
	case reflect.Uint32:
		var valUint uint32
		err = binary.Read(reader, binary.BigEndian, &valUint)
		value.SetUint(uint64(valUint))
	case reflect.Uint64:
		var valUint uint64
		err = binary.Read(reader, binary.BigEndian, &valUint)
		value.SetUint(valUint)
	default:
		return fmt.Errorf("unsupported kind: %s", value.Kind())
	}
	return err
}
