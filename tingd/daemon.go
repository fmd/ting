package main

import (
    "net/http"
    "github.com/fitstar/falcore"
)

type Daemon struct {
    Server *falcore.Server
    Pipeline *falcore.Pipeline
}

func NewDaemon(port int) *Daemon {
    d := &Daemon{}
    d.Pipeline = falcore.NewPipeline()
    d.Server = falcore.NewServer(port, d.Pipeline)

    d.InitPipeline()

    return d
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