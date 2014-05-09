package mongo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/fmd/ting/backend/response"
	"labix.org/v2/mgo/bson"
)

func ReservedTypes() []string {
	return []string{
		"string",
		"int",
		"bool",
		"list",
		"array",
	}
}

func ReservedType(t string) bool {
	for _, ty := range ReservedTypes() {
		if t == ty {
			return true
		}
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
	Id        string                      `bson:"_id" json:"_id"`
	Structure map[string]ContentTypeField `bson:"structure" json:"structure"`
}

func (r *Repo) PushContentType(name string, structure []byte) *response.R {
	var err error

	//Check if the type is reserved.
	if ReservedType(name) {

		//We can't edit the structure of reserved types.
		return response.Error(errors.New(fmt.Sprintf("'%s' is a reserved content type id.", name)))
	}

	//Get the collection and create a *ContentType to unmarshal into.
	c := r.Db.C(structuresCollection)
	s := &ContentType{Id: name}

	//Attempt to unmarshal the structure into the *ContentType.
	err = json.Unmarshal(structure, &s.Structure)
	if err != nil {
		return response.Error(err)
	}

	//Make sure that every field is valid:
	//Get all types, and ensure that every field refers to either an existing type, or the type that we are pushing.
	resp := r.ContentTypes()
	if resp.Error != nil {
		return response.Error(resp.Error)
	}

	types := append(resp.Data.([]string), ReservedTypes()...)
	types = append(types, name)

	for _, field := range s.Structure {
		found := false
		for _, ty := range types {
			if field.ContentType == ty {
				found = true
			}
		}

		if !found {
			return response.Error(errors.New(fmt.Sprintf("Content type '%s' does not exist.", field.ContentType)))
		}
	}

	//Now that we've validated the structure, we can upsert to Mongo.
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
