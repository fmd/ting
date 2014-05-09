package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"io/ioutil"
	"net/http"
)

func (d *Daemon) Routes() {
	m := d.Martini
	m.Group("/types", func(r martini.Router) {
		r.Get("", d.getContentTypes)
		r.Get("/:name", d.getContentType)
		r.Post("/:name/edit", d.setContentType)
	})
}

func (d *Daemon) getContentTypes(r render.Render) {
	r.JSON(RenderToResponse(d.Backend.ContentTypes()))
}

func (d *Daemon) getContentType(r render.Render, params martini.Params) {
	r.JSON(RenderToResponse(d.Backend.ContentType(params["name"])))
}

func (d *Daemon) setContentType(r render.Render, params martini.Params, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(500, &ResponseWrapper{Data: nil, Status: "error", Message: err.Error()})
	}

	r.JSON(RenderToResponse(d.Backend.PushContentType(params["name"], body)))
}
