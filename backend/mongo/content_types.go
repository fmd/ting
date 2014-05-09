package mongo

import (
	"github.com/fmd/ting/backend/response"
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"errors"
	"fmt"
)

//Reserved content types
const (
	CT_STRING string = "string"
	CT_BOOL   string = "bool"
	CT_INT	  string = "int"
	CT_LIST   string = "list"
	CT_ARRAY  string = "array"
)

func ReservedType(t string) bool {
	if  t == CT_STRING ||
	  t == CT_BOOL ||
	  t == CT_INT ||
	  t == CT_LIST ||
	  t == CT_ARRAY {
		return true
	}
	return false
}

//ContentTypeField is the Content Type Field struct
type ContentTypeField struct {
	ContentType string      `bson:"content_type" json:"content_type"`
	Required    bool        `bson:"required" json:"required"`
	Default     interface{} `bson:"default" json:"default"`
}

//ContentType is the Content Type struct
type ContentType struct {
	Id        string      `bson:"_id" json:"_id"`
	Structure interface{} `bson:"structure" json:"structure"`
}

func (r *Repo) PushContentType(name string, structure []byte) *response.R {
	var err error

	if ReservedType(name) {
		return response.Error(errors.New(fmt.Sprintf("'%s' is a reserved content type id.", name)))
	}

	c := r.Db.C(structuresCollection)
	s := &ContentType{Id: name}

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
	res := &ContentType{}

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
	s := &ContentType{}

	err = c.FindId(name).One(&s)
	if err != nil {
		return response.Error(err)
	}

	return response.Success(s)
}