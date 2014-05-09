package main

func (d *Daemon) Routes() {
	d.Martini.Get("/types", d.getContentTypes)
    d.Martini.Get("/type", d.getContentType)
}

func (d *Daemon) getContentTypes() (int, string) {
	return RenderToJson(d.Ting.Backend.ContentTypes())
}

func (d *Daemon) getContentType() (int, string) {
    return 200, "OK"
}