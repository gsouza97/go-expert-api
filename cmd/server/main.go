package main

import "github.com/gsouza97/go-expert-api/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
