package main

import (
	"errors"
	"fmt"
	"github.com/fmd/ting/backend"
	"github.com/fmd/ting/backend/mongo"
	"github.com/fmd/ting/backend/response"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"os"
)

//NewBackend creates a Backend instance from the credentials.
//This function uses the c["dbback"] to choose a database backend.
func NewBackend(c backend.Credentials) (backend.B, error) {
	var err error
	var b backend.B

	switch c["dbback"] {
	case "mongodb":
		b, err = mongo.NewRepo(c)
	case "couchdb":
		return nil, errors.New("CouchDB currently unsupported.")
	default:
		return nil, errors.New(fmt.Sprintf("Invalid backend '%s'", c["dbback"]))
	}

	if err != nil {
		return nil, err
	}

	return b, nil
}

type Daemon struct {
	Port    string
	Backend backend.B
	Martini *martini.ClassicMartini
}

func NewDaemon(port string, b backend.Credentials) (*Daemon, error) {
	var err error
	d := &Daemon{}
	d.Port = port
	d.Backend, err = NewBackend(b)
	if err != nil {
		return nil, err
	}

	d.Martini = martini.Classic()
	d.Martini.Use(render.Renderer())
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

//ResponseWrapper is the response wrapper.
//See more at http://labs.omniti.com/labs/jsend
type ResponseWrapper struct {
	Data    interface{} `json:"data"`    //Wrapper around any returned data.
	Status  string      `json:"status"`  // "success" | "fail" | "error"
	Message string      `json:"message"` // Error message.
}

//EncodeResponse encodes a JSON response to send back to the client.
//It takes an interface, attempts to Marshal it, and returns an error otherwise.
//http://labs.omniti.com/labs/jsend
func RenderToResponse(r *response.R) (int, *ResponseWrapper) {
	code := 200
	j := &ResponseWrapper{}

	if r.Error != nil {
		j.Data = nil
		j.Status = "error"
		j.Message = r.Error.Error()
		code = 500
	} else {
		j.Status = r.Status
		j.Data = r.Data
	}

	return code, j
}
