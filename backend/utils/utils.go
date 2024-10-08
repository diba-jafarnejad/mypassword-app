package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

var key = []byte("mysecretkey12345") // This is just a demo key. In production, use a secure key!

// Encrypt encrypts the given password using AES encryption
func Encrypt(password string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := []byte("randomInitVec") // Random initialization vector
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ciphertext := aesgcm.Seal(nil, nonce, []byte(password), nil)
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypts the given encrypted password
func Decrypt(encrypted string) (string, error) {
	ciphertext, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	nonce := []byte("randomInitVec")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
