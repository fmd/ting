package mongo

import (
	"encoding/json"
	"github.com/fmd/ting/backend"
	"labix.org/v2/mgo/bson"
)

func (r *Repo) StructureType(structure []byte) error {
	var err error
	c := r.Db.C(structuresCollection)
	s := &backend.ContentType{}
	err = json.Unmarshal(structure, &s)
	if err != nil {
		return err
	}

	_, err = c.Upsert(bson.M{"_id": s.Id}, s)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) ContentTypes() ([]string, error) {
	types := make([]string, 0)
	c := r.Db.C(structuresCollection)
	it := c.Find(nil).Iter()
	res := &backend.ContentType{}

	for it.Next(&res) {
		types = append(types, res.Id)
	}

	if err := it.Close(); err != nil {
		return nil, err
	}

	return types, nil
}
