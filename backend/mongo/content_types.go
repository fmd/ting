package mongo

import (
	"encoding/json"
	"github.com/fmd/ting/backend"
	"github.com/fmd/ting/backend/response"
	"labix.org/v2/mgo/bson"
)

func (r *Repo) PushContentType(name string, structure []byte) *response.R {
	var err error

	c := r.Db.C(structuresCollection)
	s := &backend.ContentType{Id: name}

	err = json.Unmarshal(structure, &s.Structure)
	if err != nil {
		return response.Error(err)
	}

	_, err = c.Upsert(bson.M{"_id": s.Id}, s)
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
