package cryptohelper

import (
	"crypto/rand"

	"github.com/gopyai/go-err"
)

func GenerateKey(numBytes int) []byte {
	key := make([]byte, numBytes)
	_, e := rand.Read(key)
	err.Panic(e)
	return key
}
