package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Encrypt(t *testing.T) {
	//given
	key := "12345678901234561234567890123456"
	pt := "https://example.com/path/path2"
	expected := "000000000000000000000000caeeddd760fd39cbd474a9253094cc1feeea08d017c17d1aafdfba7455d62e5d85fc2c8f101db78d536d54e7c299"
	//when
	ct := EncryptAES([]byte(pt), []byte(key))

	assert.Equal(t, expected, ct)
}

func Test_Decrypt(t *testing.T) {
	//given
	key := "12345678901234561234567890123456"
	ct := "000000000000000000000000caeeddd760fd39cbd474a9253094cc1feeea08d017c17d1aafdfba7455d62e5d85fc2c8f101db78d536d54e7c299"
	expected := "https://example.com/path/path2"
	//when
	pt := DecryptAES([]byte(ct), []byte(key))

	assert.Equal(t, expected, pt)
}
