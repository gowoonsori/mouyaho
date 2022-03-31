package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"mouyaho/config"
	"net/http"
)

func CreateCookie(token string) http.Cookie {
	return http.Cookie{
		Name:       "my_session",
		Value:      encryptAES([]byte(token), []byte(config.App.CipherKey)),
		Path:       "/",
		RawExpires: "",
		MaxAge:     1 * 60 * 60 * 30,
		Secure:     true,
		HttpOnly:   true,
	}
}

func DecryptCookie(cookie http.Cookie) string {
	c := cookie.Value
	a := decryptAES([]byte(c), []byte(config.App.CipherKey))
	return a
}

func encryptAES(pt, key []byte) string {
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

func decryptAES(ct, key []byte) string {
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
