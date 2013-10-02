package cli

import (
	"flag"
)

type Dispatcher struct {
	root *Module
}

func NewDispatcher() *Dispatcher {
	root := NewModule(nil, "")
	root.AddCommand("hello", NewHello())
	root.AddCommand("register", NewKd())
	kite := root.AddModule("kite", "Includes commands related to kites")
	kite.AddCommand("create", NewCreate())
	kite.AddCommand("run", NewRun())

	return &Dispatcher{root: root}
}

func (d *Dispatcher) Run() error {
	command, err := d.findCommand()
	if err != nil {
		return err
	}
	if command != nil {
		return (*command).Exec()
	}
	return nil
}

func (d *Dispatcher) findCommand() (*Command, error) {
	flag.Parse()
	args := flag.Args()
	module, err := d.root.FindModule(args)
	if err != nil {
		return nil, err
	}
	if module != nil {
		return module.Command, nil
	}
	return nil, nil
}
