package mongo

import (
	"github.com/fmd/ting/backend/response"
	"github.com/fmd/ting/backend"
	"labix.org/v2/mgo/bson"
)

func (r *Repo) PushContentType(t *backend.ContentType) *response.R {
	var err error

	c := r.Db.C(structuresCollection)

	_, err = c.Upsert(bson.M{"_id": t.Id}, t)
	if err != nil {
		return response.Error(err)
	}

	return response.Success(nil)
}

func (r *Repo) ContentTypes() *response.R {
	var err error

	types := make([]string, 0)
	c := r.Db.C(structuresCollection)
	it := c.Find(nil).Iter()
	res := &backend.ContentType{}

	for it.Next(&res) {
		types = append(types, res.Id)
	}

	if err = it.Close(); err != nil {
		return response.Error(err)
	}

	return response.Success(types)
}

func (r *Repo) ContentType(name string) *response.R {
	var err error

	c := r.Db.C(structuresCollection)
	s := &backend.ContentType{}

	err = c.FindId(name).One(&s)
	if err != nil {
		return response.Error(err)
	}

	return response.Success(s)
}
