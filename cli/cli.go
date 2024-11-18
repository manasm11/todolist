package cli

import "errors"

type Cli struct{}

func NewCli() *Cli {
	return &Cli{}
}

func (c *Cli) Run() error {
	return errors.New("error")
}
