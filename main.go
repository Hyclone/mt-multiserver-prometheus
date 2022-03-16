package main

import (
	"fmt"

	proxy "github.com/HimbeerserverDE/mt-multiserver-proxy"
)

func init() {
	fmt.Println("Hello from plugin!")
	fmt.Println("Players" + fmt.Sprint(proxy.Players()))
}