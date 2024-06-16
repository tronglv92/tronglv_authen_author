package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"flag"
	"fmt"
	"github/tronglv_authen_author/helper/server"
	"github/tronglv_authen_author/internal/config"
	"github/tronglv_authen_author/internal/handler"
	"github/tronglv_authen_author/internal/registry"
	isvc "github/tronglv_authen_author/internal/server"

	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func GeneratePKCS1PrivateKey() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("Error generating private key:", err)
		return
	}

	// Convert the private key to PKCS#1 ASN.1 PEM format
	privateKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	)

	// Encode the PEM to a base64 string
	privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKeyPEM)

	fmt.Println("Base64 Encoded PEM Private Key:")
	fmt.Println(privateKeyBase64)

	publicKey := &privateKey.PublicKey

	// Convert the public key to PKIX ASN.1 PEM format
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("Error marshalling public key:", err)
		return
	}
	publicKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		},
	)

	fmt.Println("PEM Encoded Public Key:")
	fmt.Println(string(publicKeyPEM))

}
func main() {

	// GeneratePKCS1PrivateKey()

	c := config.Load(configFile)

	svcGroup := service.NewServiceGroup()

	svcGroup.Add(server.NewGrpcServer(c.Server,
		isvc.NewGrpcHandler(registry.NewServiceContext(c)),
	))
	svcGroup.Add(server.NewHttpServer(c.Server,
		handler.NewRestHandler(registry.NewServiceContext(c)),
	))
	defer svcGroup.Stop()
	fmt.Printf("Starting server at %s:%d...\n", c.Server.Http.Host, c.Server.Http.Port)
	svcGroup.Start()

}
