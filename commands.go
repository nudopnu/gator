package main

import "fmt"

type command struct {
	name string
	args []string
}
type commands struct {
	handler map[string]func(*state, command) error
}

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
