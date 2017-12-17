package cryptohelper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"github.com/gopyai/go-err"
)

func AESGenerateKey() []byte {
	return GenerateKey(32)
}

func AESEncrypt(key, msg []byte) []byte {
	block, e := aes.NewCipher(key)
	err.Panic(e)
	ciphermsg := make([]byte, aes.BlockSize+len(msg))
	iv := ciphermsg[:aes.BlockSize]
	_, e = rand.Read(iv)
	err.Panic(e)
	cipher.NewCTR(block, iv).XORKeyStream(ciphermsg[aes.BlockSize:], msg)
	return ciphermsg
}

func AESDecrypt(key, ciphermsg []byte) ([]byte, error) {
	block, e := aes.NewCipher(key)
	if e != nil {
		return nil, e
	}
	iv := ciphermsg[:aes.BlockSize]
	msg := make([]byte, len(ciphermsg)-aes.BlockSize)
	cipher.NewCTR(block, iv).XORKeyStream(msg, ciphermsg[aes.BlockSize:])
	return msg, nil
}
