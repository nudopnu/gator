package main

import (
	"fmt"

	"github.com/nudopnu/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg.SetUser("nudopnu")
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cfg)
}
