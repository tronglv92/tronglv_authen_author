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

	privateKeyInterface, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// privateKey, ok := privateKeyInterface.(*rsa.PrivateKey)
	// if !ok {
	// 	return nil, fmt.Errorf("failed to cast private key to *rsa.PrivateKey")
	// }

	return privateKeyInterface, nil
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

func ConvertPrivateKeyToPEM(privateKey *rsa.PrivateKey) (string, error) {
	// Convert the RSA private key to DER format
	der := x509.MarshalPKCS1PrivateKey(privateKey)
	if der == nil {
		return "", fmt.Errorf("failed to marshal RSA private key")
	}

	// Create a PEM block with the DER encoded private key
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: der,
	}

	// Encode the PEM block to a string
	pemData := pem.EncodeToMemory(block)
	if pemData == nil {
		return "", fmt.Errorf("failed to encode PEM block")
	}

	return string(pemData), nil
}
