package main

import "github.com/CaiqueRibeiro/4-api/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config)
}
