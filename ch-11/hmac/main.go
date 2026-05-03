package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var key = []byte("some random key")

func checkMAC(message, recvMAC []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	calcMAC := mac.Sum(nil)

	return hmac.Equal(calcMAC, recvMAC)
}

func main() {
	// 在实际实现中,我们会从网络源读取消息和HMAC值
	message := []byte("The red eagle flies at 10:00")
	mac, _ := hex.DecodeString("69d2c7b6fbbfcaeb72a3172f4662601d1f16acfb46339639ac8c10c8da64631d")
	if checkMAC(message, mac) {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("NOT EQUAL")
	}
}
