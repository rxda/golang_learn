package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// 加密：采用sha1算法加密后转base64格式
func RsaEncryptWithSha1Base64(originalData,publicKey string)(string,error){
	pemData, err := ioutil.ReadFile(publicKey)
	if err != nil {
		panic(err)
	}
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		log.Fatal("failed to decode PEM block containing public key")
	}
	pubKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	encryptedData,err:=rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	return base64.StdEncoding.EncodeToString(encryptedData),err
}

func RsaEncryptWithSha1Hex(originalData,publicKey string)(string,error){
	key, _ := base64.StdEncoding.DecodeString(publicKey)
	pubKey, _ := x509.ParsePKIXPublicKey(key)
	encryptedData,err:=rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(originalData))
	return hex.EncodeToString(encryptedData),err
}

// 解密：对采用sha1算法加密后转base64格式的数据进行解密（私钥PKCS1格式）
func RsaDecryptWithSha1Base64(encryptedData,privateKey string)(string,error){
	privData,err := ioutil.ReadFile(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	privDatablock, _ := pem.Decode(privData)
	pk, err := x509.ParsePKCS1PrivateKey(privDatablock.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	encryptedDecodeBytes,err:=base64.StdEncoding.DecodeString(encryptedData)
	if err!=nil {
		return "",err
	}
	originalData,err:=rsa.DecryptPKCS1v15(rand.Reader,pk,encryptedDecodeBytes)
	return string(originalData),err
}

func RsaDecryptWithSha1Hex(encryptedData,privateKey string)(string,error){
	encryptedDecodeBytes,err:=hex.DecodeString(encryptedData)
	if err!=nil {
		return "",err
	}
	key,_:=base64.StdEncoding.DecodeString(privateKey)
	prvKey,_:=x509.ParsePKCS1PrivateKey(key)
	originalData,err:=rsa.DecryptPKCS1v15(rand.Reader,prvKey,encryptedDecodeBytes)
	return string(originalData),err
}

// 签名：采用sha1算法进行签名并输出为hex格式（私钥PKCS8格式）
func RsaSignWithSha1Hex(data string, prvKey string) (string, error) {
	keyByts, err := hex.DecodeString(prvKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyByts)
	if err != nil {
		fmt.Println("ParsePKCS8PrivateKey err", err)
		return "", err
	}
	h := sha1.New()
	h.Write([]byte([]byte(data)))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA1, hash[:])
	if err != nil {
		fmt.Printf("Error from signing: %s\n", err)
		return "", err
	}
	out := hex.EncodeToString(signature)
	return out, nil
}

//验签：对采用sha1算法进行签名后转base64格式的数据进行验签
func RsaVerySignWithSha1Base64(originalData, signData, pubKey string) error{
	sign, err := base64.StdEncoding.DecodeString(signData)
	if err != nil {
		return err
	}
	public, _ := base64.StdEncoding.DecodeString(pubKey)
	pub, err := x509.ParsePKIXPublicKey(public)
	if err != nil {
		return err
	}
	hash := sha1.New()
	hash.Write([]byte(originalData))
	return rsa.VerifyPKCS1v15(pub.(*rsa.PublicKey), crypto.SHA1, hash.Sum(nil), sign)
}

func main() {
	//server
	//_ = GenRsaKey(1024, "server")
	//clientA
	//_ = GenRsaKey(1024, "clientA")
	//clientB
	//_ = GenRsaKey(1024, "clientB")

	//sha256sign("hello from A")
	//公钥加密
	res, _ := RsaEncryptWithSha1Base64("hello world", "./ssl/clientA/server_public.pem")
	fmt.Println(res)
	//私钥解密
	//res := "Gl3qMuJWFKgZLQ+DB853F6lCvqzOnXcuv3N6uHiLMnbtZiSNWFAOE0RrftNc42XlYA0r7RFzJ7tacYXVelcgFbu2bygzw30OpugDu6JUr9ZH0T9syjDQpT4TQmqyhm2dprXVoVeqRdJK10X2IWOVDAoxRF+TuSrU62zB+ZnP+6w="
	res, _ = RsaDecryptWithSha1Base64(res, "./ssl/server/server_private.pem")
	fmt.Println(res)
}
