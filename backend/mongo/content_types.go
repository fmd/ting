package mongo

import (
    "github.com/fmd/ting/backend"
    "labix.org/v2/mgo/bson"
)

func (r *Repo) PushContentType(t *backend.ContentType) error {
    var err error

    c := r.Db.C(structuresCollection)

    _, err = c.Upsert(bson.M{"_id": t.Id}, t)
    if err != nil {
        return err
    }

    return nil
}

func (r *Repo) ContentTypes() ([]string, error) {
    var err error

    types := make([]string, 0)
    c := r.Db.C(structuresCollection)
    it := c.Find(nil).Iter()
    res := &backend.ContentType{}

    for it.Next(&res) {
        types = append(types, res.Id)
    }

    if err = it.Close(); err != nil {
        return nil, err
    }

    return types, nil
}

func (r *Repo) ContentType(name string) (interface{}, error) {
    var err error

    c := r.Db.C(structuresCollection)
    s := &backend.ContentType{}

    err = c.FindId(name).One(&s)
    if err != nil {
        return nil, err
    }

    return s, nil
}
