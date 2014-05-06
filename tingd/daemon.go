package main

import (
    "net/http"
    "github.com/fmd/ting/ting"
    "github.com/fmd/ting/credentials"
    "github.com/fitstar/falcore"
)

type Daemon struct {
    Ting *ting.Ting
    Server *falcore.Server
    Pipeline *falcore.Pipeline
}

func NewDaemon(c credentials.Credentials, port int) (*Daemon, error) {
    var err error
    d := &Daemon{}
    d.Pipeline = falcore.NewPipeline()
    d.Server = falcore.NewServer(port, d.Pipeline)

    d.InitPipeline()

    d.Ting, err = ting.NewTing(c)
    if err != nil {
        return nil, err
    }

    return d, nil
}

func (d *Daemon) ListenAndServe() error {
    err := d.Server.ListenAndServe()
    if err != nil {
        return err
    }
    return nil
}

func (d *Daemon) InitPipeline() {
    d.Pipeline.Upstream.PushBack(helloFilter)
}

var helloFilter = falcore.NewRequestFilter(func(req *falcore.Request) *http.Response {
    return falcore.StringResponse(req.HttpRequest, 200, nil, "hello world!")
})