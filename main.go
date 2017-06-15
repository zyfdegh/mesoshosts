package main

import (
	"fmt"
	"log"
)

func main() {
	masterIP := "35.185.156.236"
	hostnames, err := mesosHosts(masterIP)
	if err != nil {
		log.Printf("retrieve hostnames error: %v\n", err)
		return
	}
	for _, h := range hostnames {
		fmt.Println(h)
	}
}
