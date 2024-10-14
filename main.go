package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/nudopnu/gator/internal/config"
)

type (
	state struct {
		cfg *config.Config
	}
	command struct {
		name string
		args []string
	}
	commands struct {
		handler map[string]func(*state, command) error
	}
)

func (c *commands) register(name string, f func(*state, command) error) {
	c.handler[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handler[cmd.name]
	if !ok {
		return fmt.Errorf("handler for command '%s' not found", cmd.name)
	}
	return handler(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("expecting name as an argument")
	}
	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("set username to '%s'\n", username)
	return nil
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
