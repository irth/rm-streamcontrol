package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello")
	o := OBS{
		Address:  "localhost:4455",
		Password: "secret",
	}
	if err := o.Connect(); err != nil {
		panic(err)
	}

	for _, scene := range o.Scenes {
		log.Printf("%d %s", scene.SceneIndex, scene.SceneName)
	}
}
