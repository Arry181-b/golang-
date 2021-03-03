package utils

import (
	"bytes"
	"encoding/binary"
)

/**
 *将int类型数据转换为[]byte类型
 */
func Int2Byte(num int64) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)

	return buff.Bytes(), err
}
