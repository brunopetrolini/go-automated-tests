package main

import (
	"automatedtests/address"
	"fmt"
)

func main() {
	addressType := address.ValidateAddressType("Rua dos Campos")
	fmt.Println(addressType)
}
