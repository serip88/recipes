package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"

	router_pkg "csrf-session-rn/router"

	"github.com/gofiber/fiber/v2/middleware/logger" // Middleware for logging HTTP requests
	// "github.com/gofiber/template/html/v2"
	"github.com/gofiber/template/django/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	servicev1 "rain.io/protogen/service/v1"
)

func main() {

	conn, err := grpc.NewClient("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := servicev1.NewAddServiceClient(conn)
	//B test client
	fmt.Println("Start client...")
	user := &servicev1.User{
		Email: "serip88@yahoo.com",
	}
	req := &servicev1.Request{
		Module: servicev1.Module_MODULE_LOGIN,
		User:   user}
	if res, err := client.GetUser(context.Background(), req); err == nil {
		fmt.Println("Res User...", res.User)
	} else {
		fmt.Println("Login fails...", err.Error())
	}
	fmt.Println("End client...")
	//E test client

	// In production, run the app on port 443 with TLS enabled
	// or run the app behind a reverse proxy that handles TLS.
	//
	// It is also recommended that the csrf cookie is set to be
	// Secure and HttpOnly and have the SameSite attribute set
	// to Lax or Strict.
	//
	// In this example, we use the "__Host-" prefix for cookie names.
	// This is suggested when your app uses secure connections (TLS).
	// A cookie with this prefix is only accepted if it's secure,
	// comes from a secure source, doesn't have a Domain attribute,
	// and its Path attribute is "/".
	// This makes these cookies "locked" to the domain.
	//
	// See the following for more details:
	// https://cheatsheetseries.owasp.org/cheatsheets/Cross-Site_Request_Forgery_Prevention_Cheat_Sheet.html
	//
	// It's recommended to use the "github.com/gofiber/fiber/v2/middleware/helmet"
	// middleware to set headers to help prevent attacks such as XSS, man-in-the-middle,
	// protocol downgrade, cookie hijacking, SSL stripping, clickjacking, etc.

	// HTML templates
	// engine := html.New("./views", ".html")
	engine := django.New("./views", ".html")
	// django Doc
	// https://github.com/gofiber/template/blob/master/django/README.md
	// https://docs.djangoproject.com/en/dev/topics/templates/

	//B develop mode
	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	// Debug will print each template that is parsed, good for debugging
	// engine.Debug(true) // Optional. Default: false
	//E develop mode

	// Create a Fiber app
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})
	app.Use(logger.New()) // Use logger middleware to log HTTP requests

	router := router_pkg.New(&client)
	router.SetupRoutes(app)

	certFile := "cert.pem"
	keyFile := "key.pem"

	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		fmt.Println("Self-signed certificate not found, generating...")
		if err := generateSelfSignedCert(certFile, keyFile); err != nil {
			panic(err)
		}
		fmt.Println("Self-signed certificate generated successfully")
		fmt.Println("You will need to accept the self-signed certificate in your browser")
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		panic(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	ln, err := tls.Listen("tcp", "127.0.0.1:8443", config)
	if err != nil {
		panic(err)
	}

	app.Listener(ln)
}

// generateSelfSignedCert generates a self-signed certificate and key
// and saves them to the specified files
//
// This is only for testing purposes and should not be used in production
func generateSelfSignedCert(certFile string, keyFile string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),

		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certOut, err := os.Create(certFile)
	if err != nil {
		return err
	}
	defer certOut.Close()

	_ = pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})

	keyOut, err := os.Create(keyFile)
	if err != nil {
		return err
	}
	defer keyOut.Close()

	_ = pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(priv)})

	return nil
}
