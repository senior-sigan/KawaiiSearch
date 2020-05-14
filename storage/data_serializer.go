package storage

import (
	"bytes"
	"encoding/binary"
)

func MarshalFloats(farray []float32) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, int64(len(farray)))
	if err != nil {
		return nil, err
	}
	err = binary.Write(buf, binary.BigEndian, farray)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnmarshallFloats(data []byte) ([]float32, error) {
	buf := bytes.NewReader(data)
	var size int64
	err := binary.Read(buf, binary.BigEndian, &size)
	if err != nil {
		return nil, err
	}
	array := make([]float32, size)
	err = binary.Read(buf, binary.BigEndian, array)
	if err != nil {
		return nil, err
	}
	return array, nil
}
