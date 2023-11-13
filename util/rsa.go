package util

import (
	"carservice/infra/config"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

// http://www.metools.info/code/c80.html
// 生成PKCS#1的公私钥

func RsaEncrypt(origData []byte) ([]byte, error) {
	var publicKey = []byte(config.GetServerConfig().Rsa.PublicKey)
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return nil, errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	var privateKey = []byte(config.GetServerConfig().Rsa.PrivateKey)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func RsaEncryptBase64(origData string) (string, error) {
	var publicKey = []byte(config.GetServerConfig().Rsa.PublicKey)
	block, _ := pem.Decode(publicKey) //将密钥解析成公钥实例
	if block == nil {
		return "", errors.New("public key error")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	cypherbytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	if err != nil {
		return "", err
	}
	cypherText := base64.RawStdEncoding.EncodeToString(cypherbytes)
	return cypherText, nil
}

func RsaDecryptBase64(ciphertext string) (string, error) {
	cypherBytes, err := base64.RawStdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	var privateKey = []byte(config.GetServerConfig().Rsa.PrivateKey)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("private key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	originalBytes, err := rsa.DecryptPKCS1v15(rand.Reader, priv, cypherBytes)
	if err != nil {
		return "", err
	}

	return string(originalBytes), nil
}
