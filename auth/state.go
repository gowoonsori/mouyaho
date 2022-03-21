package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func EncryptAES(pt, key []byte) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	ct := gcm.Seal(nonce, nonce, pt, nil)
	return hex.EncodeToString(ct)
}

func DecryptAES(ct, key []byte) string {
	ct, _ = hex.DecodeString(string(ct))
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}

	nonceSize := gcm.NonceSize()
	nonce, pct := ct[:nonceSize], ct[nonceSize:]

	pt, err := gcm.Open(nil, nonce, pct, nil)

	return string(pt)
}
