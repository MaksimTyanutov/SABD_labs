package main

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func loadPrivateKeyWithPassword(keyFile, password string) (*rsa.PrivateKey, error) {
	keyData, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(keyData)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the key")
	}

	// Decrypt the private key using the provided password
	der, err := x509.DecryptPEMBlock(block, []byte(password))
	if err != nil {
		return nil, err
	}

	key, err := x509.ParsePKCS1PrivateKey(der)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func main() {
	http.HandleFunc("/", handler)

	// Загрузка сертификата
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		fmt.Println("Error loading certificate:", err)
		return
	}

	// Настройка конфигурации сервера с использованием TLS
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}

	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on https://localhost:8443")
	err = server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
