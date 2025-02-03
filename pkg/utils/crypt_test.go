package utils

import (
	"fmt"
	"testing"
)

// var str = "luban666."
var str = "root:luban666.@tcp(localhost:3306)/user"
var key = "sjguejglsjejguej"

// var encryptData = "xNTpIpkLKZSgUvfhTqu/8g=="
var encryptData = "wLDXTtRYJcOhS0Zo7lCmF17gQ92Rd+qQqwy0T71B2XotwohNzHvNO+1WCK924xrB"

func TestEncrypt(t *testing.T) {
	t.Run(
		"encrypt the data", func(t *testing.T) {
			encryptData, err := Encrypt(str, []byte(key))
			if err != nil {
				t.Errorf("encrypt the %s failed: %v", str, err)
			}
			fmt.Println(string(encryptData))
		},
	)
}

func TestDecrypt(t *testing.T) {
	t.Run(
		"encrypt the data", func(t *testing.T) {
			orginalData, err := DeCrypt(encryptData, []byte(key))
			if err != nil {
				t.Errorf("decrypt the %s failed: %v", str, err)
			}
			if string(orginalData) != str {
				t.Errorf("decrypt the %s failed, output: %s", str, string(orginalData))
			}
			fmt.Println(string(orginalData))
		},
	)
}
