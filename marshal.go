package cryptohelper

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"
)

func MarshalPrivateKey(pri *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(pri)
}

func UnmarshalPrivateKey(b []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(b)
}

func MarshalPublicKey(pub *rsa.PublicKey) ([]byte, error) {
	return x509.MarshalPKIXPublicKey(pub)
}

func UnmarshalPublicKey(b []byte) (*rsa.PublicKey, error) {
	p, e := x509.ParsePKIXPublicKey(b)
	if e != nil {
		return nil, e
	}
	return p.(*rsa.PublicKey), nil
}

func WritePrivateKeyToFile(f *os.File, pri *rsa.PrivateKey) error {
	return pem.Encode(f, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: MarshalPrivateKey(pri)})
}

func SavePrivateKey(fileName string, pri *rsa.PrivateKey) error {
	f, e := os.Create(fileName)
	if e != nil {
		return e
	}
	defer f.Close()
	return WritePrivateKeyToFile(f, pri)
}

func ReadPrivateKeyFromFile(f *os.File) (*rsa.PrivateKey, error) {
	b, e := ioutil.ReadAll(f)
	if e != nil {
		return nil, e
	}
	block, _ := pem.Decode(b)
	return UnmarshalPrivateKey(block.Bytes)
}

func LoadPrivateKey(fileName string) (*rsa.PrivateKey, error) {
	f, e := os.Open(fileName)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	return ReadPrivateKeyFromFile(f)
}

func WritePublicKeyToFile(f *os.File, pub *rsa.PublicKey) error {
	b, e := MarshalPublicKey(pub)
	if e != nil {
		return e
	}
	return pem.Encode(f, &pem.Block{Type: "PUBLIC KEY", Bytes: b})
}

func SavePublicKey(fileName string, pub *rsa.PublicKey) error {
	f, e := os.Create(fileName)
	if e != nil {
		return e
	}
	defer f.Close()
	return WritePublicKeyToFile(f, pub)
}

func ReadPublicKeyFromFile(f *os.File) (*rsa.PublicKey, error) {
	b, e := ioutil.ReadAll(f)
	if e != nil {
		return nil, e
	}
	block, _ := pem.Decode(b)
	return UnmarshalPublicKey(block.Bytes)
}

func LoadPublicKey(fileName string) (*rsa.PublicKey, error) {
	f, e := os.Open(fileName)
	if e != nil {
		return nil, e
	}
	defer f.Close()
	return ReadPublicKeyFromFile(f)
}
