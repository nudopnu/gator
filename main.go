package main

import (
	"log"
	"os"

	"github.com/nudopnu/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
		return
	}
	s := state{
		cfg: &cfg,
	}
	cmds := commands{
		handler: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("no command provided")
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
