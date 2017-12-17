package cryptohelper

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/gopyai/go-err"
)

func RSAGenerateKey() *rsa.PrivateKey {
	pri, e := rsa.GenerateKey(rand.Reader, 2048)
	err.Panic(e)
	return pri
}

func RSAEncrypt(pub *rsa.PublicKey, key []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, pub, key, []byte(""))
}

func RSADecrypt(pri *rsa.PrivateKey, cipherkey []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, pri, cipherkey, []byte(""))
}

func RSASign(pri *rsa.PrivateKey, hashed []byte) ([]byte, error) {
	return rsa.SignPSS(rand.Reader, pri, crypto.SHA256, hashed, &rsa.PSSOptions{})
}

func RSAVerify(pub *rsa.PublicKey, hashed, signed []byte) bool {
	return nil == rsa.VerifyPSS(pub, crypto.SHA256, hashed, signed, &rsa.PSSOptions{})
}
