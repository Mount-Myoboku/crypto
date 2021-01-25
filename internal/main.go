package main

import (
	"fmt"
	"github.com/Mount-Myoboku/crypto/caesarcipher"
)

func main() {
	cb := caesarcipher.Init(5)
	output := cb.Encrypt("HELLO")
	fmt.Println(output)
}
