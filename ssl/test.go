package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
)

//A sign server verify

func sha256sign(originalData string){
	//server public
	pemData,err := ioutil.ReadFile("./ssl/clientA/client_a_public.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	// A private
	privData,err := ioutil.ReadFile("./ssl/clientA/client_a_private.pem")
	if err != nil {
		log.Fatal(err)
	}
	privDatablock, _ := pem.Decode(privData)

	privateKey, err := x509.ParsePKCS1PrivateKey(privDatablock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	//先公钥加密
	encryptData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	//fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
	hash := sha1.New()
	hash.Write([]byte(encryptData))
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash.Sum(nil))

	fmt.Println("sign=", hex.EncodeToString(sign))
	fmt.Println("data=",hex.EncodeToString(encryptData))
}


func sha256verify(originalData string){
	//public
	pemData,err := ioutil.ReadFile("./ssl/clientA/client_a_public.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, rest := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	pemData,err = ioutil.ReadFile("./ssl/clientA/client_a_private.pem")
	if err != nil {
		log.Fatal(err)
	}
	block, _ = pem.Decode(pemData)

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Printf("Got a %T, with remaining data: %q", pub, rest)
	hash := sha1.New()
	hash.Write([]byte(originalData))
	encryptedData, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hash.Sum(nil))
	fmt.Println(hex.EncodeToString(encryptedData))
}
