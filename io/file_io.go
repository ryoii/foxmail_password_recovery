package io

import (
	"encoding/binary"
	"foxmail_password_recover/decrypt"
	"os"
)

func ReadFile(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
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

// FindPassWord aims to find the latest password saved in the content
func FindPassWord(content []byte) []byte {
	key := []byte("Password")
	var latest []byte
	for i := 1; i < len(content)-len(key); i++ {
		// Find "Password" and exclude that like "SmtpPassword"
		if content[i-1] == 0x00 && bytesEq(key, content[i:i+8]) {
			// skip "Password" and an unknown int32
			offset := i + len(key) + 4
			passwordLen := int(binary.LittleEndian.Uint32(content[offset : offset+4]))
			latest = content[offset+4 : offset+4+passwordLen]
		}
	}
	if latest == nil {
		panic("can not find password in this file")
	}
	return latest
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
