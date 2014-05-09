package main

func (d *Daemon) Routes() {
	d.Martini.Get("/types", d.getContentTypes)
}

func (d *Daemon) getContentTypes() (int, string) {
	t, err := d.Ting.ContentTypes()
	if err != nil {
		return 500, "Error getting content types."
	}

	return d.EncodeResponse(t)
}