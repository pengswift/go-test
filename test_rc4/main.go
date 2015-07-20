package main

import (
	"crypto/rc4"
	"fmt"
)

func main() {
	//明文
	src := []byte("Hello, world!")

	//密钥
	key := []byte("12345")

	cipher, err := rc4.NewCipher(key)
	if err != nil {
		fmt.Println("rc4.NewCipher error:" + err.Error())
	}

	encrypted := make([]byte, len(src))
	cipher.XORKeyStream(encrypted, src)
	fmt.Printf("Encrypting %s : %v -> %v\n", src, []byte(src), encrypted)

	decrypted := make([]byte, len(encrypted))
	cipher, err = rc4.NewCipher(key)
	if err != nil {
		fmt.Println("rc4.NewCipher error:" + err.Error())
	}

	cipher.XORKeyStream(decrypted, encrypted)
	fmt.Printf("Decrypting %v -> %v : %s\n", encrypted, decrypted, decrypted)

}
