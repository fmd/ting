package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/fmd/ting/backend/response"
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
	r.JSON(d.Ting.Backend.ContentTypes().ToResponse())
}

func (d *Daemon) getContentType(r render.Render, params martini.Params) {
	r.JSON(d.Ting.Backend.ContentType(params["name"]).ToResponse())
}

func (d *Daemon) setContentType(r render.Render, params martini.Params, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		r.JSON(500, &response.W{Data: nil, Status: "error", Message: err.Error()})
	}

	c, err := d.Ting.ValidateContentType(params["name"], body)
	if err != nil {
		r.JSON(500, &response.W{Data: nil, Status: "error", Message: err.Error()})
	}	

	r.JSON(d.Ting.Backend.PushContentType(c).ToResponse())
}
