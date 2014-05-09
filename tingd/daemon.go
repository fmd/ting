package main

import (
	"github.com/fmd/ting/backend"
	"github.com/fmd/ting/ting"
	"github.com/go-martini/martini"
	"encoding/json"
	"fmt"
	"os"
)

type Daemon struct {
	Port    string
	Ting    *ting.Ting
	Martini *martini.ClassicMartini
}

func NewDaemon(port string, b backend.Credentials) (*Daemon, error) {
	var err error
	d := &Daemon{}
	d.Port = port
	d.Ting, err = ting.NewTing(b)
	if err != nil {
		return nil, err
	}

	d.Martini = martini.Classic()
	d.Routes()

	err = os.Setenv("PORT", d.Port)
	if err != nil {
		return nil, err
	}

	return d, nil
}

func (d *Daemon) Run() error {
	d.Martini.Run()
	return nil
}

func (d *Daemon) EncodeResponse(in interface{}) (int, string) {
    bytes, err := json.Marshal(in)
    if err != nil {
        return 500, fmt.Sprintf("Error encoding response to JSON: %s", err)
    }

    return 200, string(bytes)
}