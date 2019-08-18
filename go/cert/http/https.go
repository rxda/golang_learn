package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func pfxToPem(){
	pfxPath := "go/cert/demofile/10000001/kuaiqian-merchant.pfx"
	pfxBytes, err := ioutil.ReadFile(pfxPath)
	if err != nil{
		log.Fatal(err)
	}
	pemBlock, err :=pkcs12.ToPEM(pfxBytes, "123456")
	if err != nil{
		log.Fatal(err)
	}

	var pemData []byte
	for _, b := range pemBlock {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}
	fmt.Println(string(pemData))
}

func main() {
	caCert, err := ioutil.ReadFile("/Users/rxda/Documents/Codes/golang/golang_learn/go/cert/http/k.crt")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				InsecureSkipVerify: true,
			},
		},
	}

	_, err = client.Post("https://sandbox.99bill.com:7445/umgw/common/distribute.html",
		"application/json",
		strings.NewReader("hello"),
		)
	if err != nil {
		panic(err)
	}
}

func request(certFile, keyFile, caFile string)  {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Load CA cert
	caCert, err := ioutil.ReadFile(caFile)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Setup HTTPS client
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	// Do GET something
	resp, err := client.Get("https://goldportugal.local:8443")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Dump response
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}