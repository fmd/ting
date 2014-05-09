package ting

import (
	"errors"
	"fmt"
	"github.com/fmd/ting/backend"
	"github.com/fmd/ting/backend/mongo"
)

type Ting struct {
	Backend backend.Backend
}

//NewTing creates a new *Ting instance.
//BUG(Needs to not default to mongo)
//Returns a *Ting and a nil error if successful, or a nil *Ting and an error otherwise.
func NewTing(c backend.Credentials) (*Ting, error) {
	var err error

	t := &Ting{}

	switch c["dbback"] {
	case "mongodb":
		t.Backend, err = mongo.NewRepo(c)
	case "couchdb":
		return nil, errors.New("CouchDB currently unsupported.")
	default:
		return nil, errors.New(fmt.Sprintf("Invalid backend '%s'", c["dbback"]))
	}

	if err != nil {
		return nil, err
	}

	return t, nil
}