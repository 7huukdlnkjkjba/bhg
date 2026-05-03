package encoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unicode/utf16"
)

func FromUnicode(d []byte) (string, error) {
	// 参考 https://github.com/Azure/go-ntlmssp/blob/master/unicode.go 的逻辑
	if len(d)%2 > 0 {
		return "", errors.New("指定了Unicode (UTF 16 LE),但数据长度不一致")
	}
	s := make([]uint16, len(d)/2)
	err := binary.Read(bytes.NewReader(d), binary.LittleEndian, &s)
	if err != nil {
		return "", err
	}
	return string(utf16.Decode(s)), nil
}

func ToUnicode(s string) []byte {
	// 参考 https://github.com/Azure/go-ntlmssp/blob/master/unicode.go 的逻辑
	uints := utf16.Encode([]rune(s))
	b := bytes.Buffer{}
	binary.Write(&b, binary.LittleEndian, &uints)
	return b.Bytes()
}
