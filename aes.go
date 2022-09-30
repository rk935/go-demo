package main

import (
	"fmt"
	"github.com/marspere/goencrypt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 500; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束

}

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	// aes decryption
	cipher, err := goencrypt.NewAESCipher([]byte("NzjQKqIfLRVibOda3keZhXBDsT48r70C"), []byte("0000000000000000"), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintHex)
	if err != nil {
		fmt.Println(err)
		return
	}
	cipherText, err := cipher.AESEncrypt([]byte("333"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("加密:", i, cipherText)

	// aes decryption
	plainText, err := cipher.AESDecrypt(cipherText)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("解密: ", i, plainText)
}
