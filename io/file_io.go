package io

import (
	"encoding/binary"
	"foxmail_password_recover/decrypt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	content, err := ioutil.ReadAll(file)
	return content
}

func GetClientType(content []byte) decrypt.Type {
	if content[0] == 0xD0 {
		return decrypt.V6
	} else if content[0] == 0x52 {
		return decrypt.V7
	} else {
		return decrypt.UNKNOWN
	}
}

func FindPassWord(content []byte) []byte {
	key := []byte("Password")
	for i := 0; i < len(content)-len(key); i++ {
		if bytesEq(key, content[i:i+8]) {
			// skip "password" and an unknown int32
			offset := i + len(key) + 4
			passwordLen := int(binary.LittleEndian.Uint32(content[offset : offset+4]))
			return content[offset+4 : offset+4+passwordLen]
		}
	}
	panic("can not find password in this file")
}

func bytesEq(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
