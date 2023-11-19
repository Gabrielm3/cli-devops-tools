package utils

import "encoding/base64"

func EncodeString(s string) string {
	encode := base64.StdEncoding.EncodeToString([]byte(s))
	return encode
}

func DecodeString(s string) string {
	decode, _ := base64.StdEncoding.DecodeString(s)
	return string(decode)
}