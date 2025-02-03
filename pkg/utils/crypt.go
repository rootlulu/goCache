package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	//Repeat()函数的功能是把切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("加密字符串错误！")
	} else {
		//获取填充字符串长度
		unpadding := int(origData[length-1])
		//截取切片，删除填充字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}
}

func Encrypt(originalData string, key []byte) (string, error) {
	originalBytes := []byte(originalData)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("the length of key must be 16, 24 or 32, its' length: %d", len(key))
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	originalBytes = PKCS7Padding(originalBytes, block.BlockSize())
	blocMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	crypted := make([]byte, len(originalBytes))
	//执行加密
	blocMode.CryptBlocks(crypted, originalBytes)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func DeCrypt(cypted string, key []byte) (string, error) {
	cypteBytes, err := base64.StdEncoding.DecodeString(cypted)
	if err != nil {
		return "", err
	}
	//创建加密算法实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	//获取块大小
	blockSize := block.BlockSize()
	//创建加密客户端实例
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(cypteBytes))
	//这个函数也可以用来解密
	blockMode.CryptBlocks(origData, cypteBytes)
	//去除填充字符串
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return "", err
	}
	return string(origData), err
}
