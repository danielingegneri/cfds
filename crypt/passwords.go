package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/ansel1/merry"
	"io"
)

func Encrypt(src string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", merry.Here(err)
	}
	if src == "" {
		return "", merry.New("No input")
	}

	content := PKCS5Padding([]byte(src), block.BlockSize())
	out := make([]byte, block.BlockSize()+len(content))
	iv := out[:block.BlockSize()] // Pointing at start of out for convenience
	// Generate IV
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	ecb := cipher.NewCBCEncrypter(block, iv)
	ecb.CryptBlocks(out[block.BlockSize():], content) // Write after the iv in out
	return base64.StdEncoding.EncodeToString(out), nil
}

func Decrypt(cryptBase64 string, seed string) (string, error) {
	input, err := base64.StdEncoding.DecodeString(cryptBase64)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher([]byte(seed))
	if err != nil {
		return "", merry.Here(err)
	}
	if len(input) == 0 {
		return "", merry.New("No input")
	}

	initialVector := input[0:16]
	ecb := cipher.NewCBCDecrypter(block, initialVector)
	decrypted := make([]byte, len(input)-len(initialVector))
	ecb.CryptBlocks(decrypted, input[len(initialVector):])

	return string(PKCS5Trimming(decrypted)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
