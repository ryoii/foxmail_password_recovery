package decrypt

import "strconv"

type Type uint8

const (
	Rec0KeyV6 = "~draGon~"
	Rec0KeyV7 = "~F@7%m$~"
)

const (
	V6 Type = iota
	V7
	UNKNOWN = 255
)

func PasswordInRec0(clientType Type, decrypt string) string {
	var key []byte
	if clientType == V6 {
		key = []byte(Rec0KeyV6)
	} else if clientType == V7 {
		key = []byte(Rec0KeyV7)
	} else {
		panic("unknown client type")
	}

	inputLen := len(decrypt)
	b := make([]byte, 0, inputLen/2)

	for len(decrypt) > 0 {
		p, err := strconv.ParseInt(decrypt[:2], 16, 0)
		if err != nil {
			panic("password decode error")
		}
		b = append(b, uint8(p))
		decrypt = decrypt[2:]
	}

	c := b
	c[0] = byte(sumBytes(key)%255) ^ c[0]

	d := make([]byte, len(b)-1)
	for i := 0; i < len(d); i++ {
		d[i] = b[i+1] ^ key[i%len(key)]
	}

	e := make([]byte, len(d))
	for i := 0; i < len(e); i++ {
		if d[i] < c[i] {
			e[i] = 0xFF - c[i] + d[i]
		} else {
			e[i] = d[i] - c[i]
		}
	}

	return string(e)
}

func sumBytes(input []byte) int {
	var tmp = 0
	for _, b := range input {
		tmp += int(b)
	}
	return tmp
}
