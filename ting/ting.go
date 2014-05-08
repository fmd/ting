package ting

import (
    "fmt"
    "errors"
    "github.com/fmd/ting/backends"
    "github.com/fmd/ting/backends/mongo"
)

type Ting struct {
    Backend Backend
}

//NewTing creates a new *Ting instance.
//BUG(Needs to not default to mongo)
//Returns a *Ting and a nil error if successful, or a nil *Ting and an error otherwise.
func NewTing(c backends.Credentials) (*Ting, error) {
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