package main

import (
	"fmt"
	"log"
)

func main() {
	s := "qwe\\45"
	unpacked, err := unpack(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unpacked)
}
