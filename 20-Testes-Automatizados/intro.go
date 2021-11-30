package main

import (
	"fmt"
	"introducao/addresses"
)

func main() {
	addressType := addresses.AddressTypes("Avenida Paulista")
	fmt.Println(addressType)
}
