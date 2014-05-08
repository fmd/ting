package main

import (
	"os"
	"github.com/go-martini/martini"
	"github.com/fmd/ting/ting"
	"github.com/fmd/ting/backend"
)

type Daemon struct {
	Ting    *ting.Ting
	Martini *martini.ClassicMartini
}

func NewDaemon(port string, b backend.Credentials) (*Daemon, error) {
	var err error
	d := &Daemon{}

	d.Ting, err = ting.NewTing(b)
	if err != nil {
		return nil, err
	}

	d.Martini = martini.Classic()
	d.Routes()

	err = os.Setenv("PORT", port)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Daemon) Run() error {
	d.Martini.Run()
	return nil
}