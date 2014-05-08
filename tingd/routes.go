package main

import (
	"encoding/json"
)

func (d *Daemon) Routes() {
	d.Martini.Get("/types", d.getContentTypes)
}

func (d *Daemon) getContentTypes() (int, string) {
	t, err := d.Ting.Backend.ContentTypes()

	if err != nil {
		return 500, "Error getting content types."
	}

	m, err := json.Marshal(t)

	if err != nil {
		return 500, "Error encoding content types JSON."
	}

	return 200, string(m)
}
