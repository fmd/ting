package main

import (
	"github.com/fmd/ting/backend/response"
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
	r.JSON(d.Ting.ContentTypes())
}

func (d *Daemon) getContentType(r render.Render, params martini.Params) {
	r.JSON(d.Ting.ContentType(params["name"]))
}

func (d *Daemon) setContentType(r render.Render, params martini.Params, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(response.Error(err).Wrap())
	}

	r.JSON(d.Ting.PushContentType(params["name"], body))
}
