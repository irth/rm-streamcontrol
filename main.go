package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
)

func main() {
	fmt.Println("hello")
	client, err := goobs.New("localhost:4455", goobs.WithPassword("secret"))
	if err != nil {
		panic(err)
	}
	v, err := client.General.GetVersion()
	fmt.Printf("OBS version: %s\n", v.ObsVersion)
	fmt.Printf("WS version: %s\n", v.ObsWebSocketVersion)
}
