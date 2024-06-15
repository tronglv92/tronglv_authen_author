package main

import (
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

// func GeneratePKCS1PrivateKey(filename string, bits int) error {
// 	// Generate RSA key
// 	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
// 	if err != nil {
// 		return fmt.Errorf("failed to generate private key: %w", err)
// 	}

// 	// Convert the RSA key to PKCS#1 ASN.1 DER form
// 	privDER := x509.MarshalPKCS1PrivateKey(privateKey)

// 	// Create a PEM block
// 	privBlock := pem.Block{
// 		Type:  "RSA PRIVATE KEY",
// 		Bytes: privDER,
// 	}

// 	// Open file for writing
// 	file, err := os.Create(filename)
// 	if err != nil {
// 		return fmt.Errorf("failed to open file for writing: %w", err)
// 	}
// 	defer file.Close()

// 	// Write the PEM block to file
// 	if err := pem.Encode(file, &privBlock); err != nil {
// 		return fmt.Errorf("failed to write data to file: %w", err)
// 	}

//		return nil
//	}
func main() {

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
