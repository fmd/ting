package ting

import (
    "github.com/fmd/ting/ting/backends/mongo"
)

type Ting struct {
    Backend Backend
}

func NewTing(hostname string, db string) (*Ting, error) {
    var err error
    t := &Ting{}
    t.Backend, err = mongo.NewRepo(hostname, db)
    if err != nil {
        return nil, err
    }

    return t, nil
}