package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	serverCert, err := ioutil.ReadFile("server.pem")
	if err != nil {
		fmt.Println("Error reading server certificate:", err)
		return
	}

	// Создание пула сертификатов
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(serverCert)

	// Создание конфигурации TLS с указанием пула сертификатов
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	//tlsConfig := &tls.Config{InsecureSkipVerify: true}

	// Создание клиента с использованием TLS конфигурации
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Выполнение запроса к серверу
	resp, err := client.Get("https://localhost:8443/hello")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Response:", string(body))
}
