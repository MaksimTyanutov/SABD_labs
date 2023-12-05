package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	// Создаем новый маршрутизатор
	router := http.NewServeMux()

	// Определяем прокси-маршрут для перенаправления запросов
	router.Handle("/hello/", http.StripPrefix("/hello", reverseProxy()))

	// Запускаем прокси-сервер на порту 8081 без TLS
	fmt.Println("Proxy server listening on :8081")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func reverseProxy() http.Handler {
	// Задаем URL сервера, на который будут перенаправляться запросы
	targetURL, _ := url.Parse("https://localhost:8443")

	// Чтение сертификата сервера
	serverCert, err := ioutil.ReadFile("server.pem")
	if err != nil {
		fmt.Println("Error reading server certificate:", err)
		return nil
	}

	// Создание пула сертификатов и добавление сертификата
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(serverCert)

	// Создаем прокси с настройками для перенаправления запросов
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// Изменяем хост-заголовок для корректного перенаправления
	proxy.Director = func(req *http.Request) {
		req.URL.Scheme = "https"
		req.URL.Host = targetURL.Host
		req.Host = targetURL.Host

	}

	// Настройка TLS для прокси с использованием созданного пула сертификатов
	proxy.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: certPool,
		},
	}

	// Модификация ответа перед отправкой клиенту
	proxy.ModifyResponse = func(resp *http.Response) error {
		// Добавляем заголовок "Test" со значением "TestSample" к ответу
		resp.Header.Add("Test", "TestSample")
		return nil
	}

	// Выполняем запрос через прокси
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})
}
