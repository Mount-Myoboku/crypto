package main

import (
	"crypto/caesarcipher"
	"fmt"
)

func main()  {
	fmt.Println(caesarcipher.Encrypt("HELLO"))
	fmt.Println(caesarcipher.Decrypt("EBIIL"))
}
