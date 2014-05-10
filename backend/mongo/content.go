package mongo

import (
    "errors"
    "fmt"
    "github.com/fmd/ting/backend"
    "labix.org/v2/mgo/bson"
)

func (r *Repo) PushContent(contentType string, content *backend.Content) error {
    var err error

    c := r.Db.C(contentType)

    if content.Id == nil {
        content.Id = bson.NewObjectId()
    }

    _, err = c.UpsertId(content.Id, content)

    if err != nil {
        return err
    }

    return nil
}

func (r *Repo) Content(contentType string, id string) (interface{}, error) {
    var err error

    c := r.Db.C(contentType)
    content := &backend.Content{}

    if !bson.IsObjectIdHex(id) {
        return nil, errors.New(fmt.Sprintf("Invalid object id '%s'", id))
    }

    err = c.FindId(bson.ObjectIdHex(id)).One(&content)
    if err != nil {
        return nil, err
    }

    return content, nil
}

//BUG(Doesn't actually use a query... just gets everything)
func (r *Repo) Contents(contentType string, query interface{}) ([]*backend.Content, error) {
    var err error

    contents := make([]*backend.Content, 0)

    c := r.Db.C(contentType)
    it := c.Find(nil).Iter()

    var res *backend.Content
    for it.Next(&res) {
        contents = append(contents, res)
        res = &backend.Content{}
    }

    if err = it.Close(); err != nil {
        return nil, err
    }

    return contents, nil
}
