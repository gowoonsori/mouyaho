package fonts

import (
	"encoding/base64"
	"log"
)

var (
	arialFont []byte
	veraFont  []byte
)

func init() {
	var err error
	arialFont, err = base64.StdEncoding.DecodeString(arialBase64)
	if err != nil {
		log.Fatalf("Couldn't decode base64 font data: %s\n", err)
	}

	veraFont, err = base64.StdEncoding.DecodeString(veraSansBase64)
	if err != nil {
		log.Fatalf("Couldn't decode base64 font data: %s\n", err)
	}
}

func GetArialFont() []byte {
	return arialFont
}

func GetVeraFont() []byte {
	return veraFont
}
