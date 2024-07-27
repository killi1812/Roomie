package Services

import (
	"chatapp/server/Helpers"
	"chatapp/server/models"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"
	"time"
)

const privateKeyPath = "keys/private_key.pem"
const publicKeyPath = "keys/public_key.pem"

func GenerateKeyPair() error {

	_, err := os.Stat(privateKeyPath)
	_, err2 := os.Stat(publicKeyPath)
	if err == nil && err2 == nil {
		return nil
		//	return fmt.Errorf("keys already exist")
	}
	_, err = os.Stat("keys")
	if os.IsNotExist(err) {
		os.Mkdir("keys", 0666)
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2047)
	if err != nil {
		return err
	}
	publicKey := &privateKey.PublicKey

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(publicKey),
	})
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	go Helpers.SaveToFile(privateKeyPath, privateKeyPEM)
	go Helpers.SaveToFile(publicKeyPath, publicKeyPEM)

	return nil
}

func LoadPublicKey() (*rsa.PublicKey, error) {
	publicKeyFile, err := os.Open(publicKeyPath)
	if err != nil {
		return nil, err
	}
	defer publicKeyFile.Close()

	publicKeyPEM := make([]byte, 2048)
	n, err := publicKeyFile.Read(publicKeyPEM)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(publicKeyPEM[:n])
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}

	rsaPubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsaPubKey, nil
}
func LoadPrivateKey() (*rsa.PrivateKey, error) {
	privateKeyFile, err := os.Open(privateKeyPath)
	if err != nil {
		return nil, err
	}
	defer privateKeyFile.Close()

	privateKeyPEM := make([]byte, 2047)
	n, err := privateKeyFile.Read(privateKeyPEM)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(privateKeyPEM[:n])
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

func GenerateCertificate(user models.User) (models.Certificate, error) {
	privateKey, err := LoadPrivateKey()
	if err != nil {
		return models.Certificate{}, err
	}

	cert := models.Certificate{
		Username:  user.Username,
		Email:     user.Email,
		Timestamp: time.Now().Unix(),
	}

	certData, err := json.Marshal(cert)
	if err != nil {
		return models.Certificate{}, err
	}

	hash := sha256.Sum256(certData)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash[:])
	if err != nil {
		return models.Certificate{}, err
	}

	cert.Signature = signature
	return cert, nil
}

func VerifyCertificate(cert models.Certificate) error {
	publicKey, err := LoadPublicKey()
	if err != nil {
		return err
	}
	signature := cert.Signature
	cert.Signature = nil
	certData, err := json.Marshal(cert)
	if err != nil {
		return err
	}

	hash := sha256.Sum256(certData)
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return fmt.Errorf("certificate verification failed: %v", err)
	}

	return nil
}
