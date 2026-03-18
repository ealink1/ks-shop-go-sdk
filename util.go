package ks_shop_go_sdk

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

func MessageDecode(content, signKey string) []byte {
	if content == "" || signKey == "" {
		return nil
	}

	key, err := base64.StdEncoding.DecodeString(signKey)
	if err != nil {
		return nil
	}

	cipherText, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}

	blockSize := block.BlockSize()
	if len(cipherText) == 0 || len(cipherText)%blockSize != 0 {
		return nil
	}

	iv := make([]byte, blockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	plain := make([]byte, len(cipherText))
	mode.CryptBlocks(plain, cipherText)

	plain, err = pkcs7Unpad(plain, blockSize)
	if err != nil {
		return nil
	}

	return plain
}

func pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if len(data) == 0 || len(data)%blockSize != 0 {
		return nil, errors.New("invalid padding size")
	}

	padLen := int(data[len(data)-1])
	if padLen == 0 || padLen > blockSize {
		return nil, errors.New("invalid padding")
	}

	for i := len(data) - padLen; i < len(data); i++ {
		if data[i] != byte(padLen) {
			return nil, errors.New("invalid padding")
		}
	}

	return data[:len(data)-padLen], nil
}
