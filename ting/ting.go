package ting

import (
    "fmt"
    "errors"
    "github.com/fmd/ting/ting/backends/mongo"
)

type Ting struct {
    Backend Backend
}

type Credentials map[string]string

func NewCredentials() Credentials {
    c := make(Credentials)
    c["dbback"] = "mongodb"
    c["dbhost"] = "localhost"
    c["dbname"] = ""
    c["dbuser"] = ""
    c["dbpass"] = ""

    return c
}

// ["backend"] : mongodb | (couchdb)
// ["hostname"] : localhost
//

//NewTing creates a new *Ting instance.
//BUG(Needs to not default to mongo)

func NewTing(c Credentials) (*Ting, error) {
    var err error

    hostname := c["dbhost"]
    dbname := c["dbname"]

    t := &Ting{}

    switch c["dbback"] {
        case "mongodb":
            t.Backend, err = mongo.NewRepo(hostname, dbname)
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