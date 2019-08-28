package main


import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	)

var wechatCertPath = "tls/cert.pem"
var wechatKeyPath = "tls/key.pem"
var wechatCAPath = "tls/ca.pem"
var wechatRefundURL = "https://sandbox.99bill.com:7445/umgw/common/distribute.html"

var _tlsConfig *tls.Config

func getTLSConfig() (*tls.Config, error) {
	if _tlsConfig != nil {
		return _tlsConfig, nil
	}

	// load cert
	cert, err := tls.LoadX509KeyPair(wechatCertPath, wechatKeyPath)
	if err != nil {
		log.Println("load wechat keys fail", err)
		return nil, err
	}

	// load root ca
	caData, err := ioutil.ReadFile(wechatCAPath)
	if err != nil {
		log.Println("read wechat ca fail", err)
		return nil, err
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
		InsecureSkipVerify: true,
	}
	return _tlsConfig, nil
}

func SecurePost(url string, xmlContent []byte) (*http.Response, error) {
	tlsConfig, err := getTLSConfig()
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: tr}
	return client.Post(
		wechatRefundURL,
		"text/xml",
		bytes.NewBuffer(xmlContent))
}

func main(){
	resp, e:= SecurePost(wechatRefundURL, []byte("12312313"))
	fmt.Println(e)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
