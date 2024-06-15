package rsa

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func ReadPrivateKeyFromFile(filename string) (*rsa.PrivateKey, error) {
	keyBytes, err := os.ReadFile(fmt.Sprintf("keys/%s", filename))
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKeyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privateKey, ok := privateKeyInterface.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast private key to *rsa.PrivateKey")
	}

	return privateKey, nil
}

func ParsePKFromPEM(keyPem string) (*rsa.PrivateKey, error) {
	kp, err := base64.StdEncoding.DecodeString(keyPem)
	if err != nil {
		return nil, fmt.Errorf("error decoding private key: %v", err)
	}

	block, _ := pem.Decode(kp)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the private key")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return key, nil
}
