package auth

import (
	"fmt"
	"testing"
)

func Test_GenerateKey(t *testing.T) {
	keys := generateKey()
	fmt.Printf("cipherKey : %s,\ncipherIvKey : %s", keys[0], keys[1])
}
