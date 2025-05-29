package main

import (
	"fmt"

	"github.com/zigzagalex/gator/internal/config"
)

func main() {
	conf, _ := config.Read()
	conf.DBURL = "postgres://example"
	conf.SetUser("alex")

	conf1, _ := config.Read()
	fmt.Printf("%v", conf1)

}
