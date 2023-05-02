package main

import (
	"fmt"

	"github.com/jclaudiotomasjr/controle-redes/api/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
