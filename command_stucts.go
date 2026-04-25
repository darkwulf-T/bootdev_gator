package main

import (
	"fmt"

	"github.com/darkwulf-T/bootdev_gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("%s is not a valid command", cmd.name)
	}
	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
