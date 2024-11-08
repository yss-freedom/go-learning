package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
在网络编程中，安全性是至关重要的。
TLS/SSL是一种用于在两个通信应用程序之间提供保密性和数据完整性的协议。
Go语言的crypto/tls包提供了对TLS/SSL的支持。
*/

func helloHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Secure World!")
}

func main() {
	cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatalf("failed to load key pair: %s", err)
	}

	caCert, err := ioutil.ReadFile("ca.pem")
	if err != nil {
		log.Fatalf("failed to read ca certificate: %s", err)
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ClientCAs:    caCertPool,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	}
	tlsConfig.BuildNameToCertificate()

	server := &http.Server{
		Addr:      ":8443",
		Handler:   http.HandlerFunc(helloHandler1),
		TLSConfig: tlsConfig,
	}

	log.Println("Starting server at :8443")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
