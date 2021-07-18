package encodeAndDecodeStrings

import (
	"encoding/base64"
	"strings"
)

type Codec struct {
	empty bool
}

var encoing = base64.StdEncoding

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	if len(strs) == 0 {
		codec.empty = true
		return ""
	}

	encode := encoing.EncodeToString(([]byte(strs[0])))
	for i := 1; i < len(strs); i++ {
		encode += "|" + encoing.EncodeToString(([]byte(strs[i])))
	}
	return encode
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	if codec.empty {
		return []string{}
	}

	if len(strs) == 0 {
		return []string{""}
	}

	result := strings.Split(strs, "|")
	for i, str := range result {
		d, _ := encoing.DecodeString(str)
		result[i] = string(d)
	}
	return result
}
