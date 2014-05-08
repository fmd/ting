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
	n, err := r.Db.CollectionNames()
	if err != nil {
		return nil, err
	}

	names := make([]string, 0)

	for _, name := range n {
		if name != "structures" && name != "system.indexes" && len(name) > 0 {
			names = append(names, name)
		}
	}

	return names, nil
}
