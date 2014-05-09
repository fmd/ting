package ting

import (
    "github.com/fmd/ting/backend/mongo"
    "github.com/fmd/ting/backend"
    "encoding/json"
    "errors"
    "fmt"
)

type Ting struct {
    Backend backend.B
}

//NewTing creates a Backend instance from the credentials.
//This function uses the c["dbback"] to choose a database backend.
func NewTing(c backend.Credentials) (*Ting, error) {
    var err error

    t := &Ting{}

    switch c["dbback"] {
    case "mongodb":
        t.Backend, err = mongo.NewRepo(c)
    case "couchdb":
        return nil, errors.New("CouchDB currently unsupported.")
    default:
        return nil, errors.New(fmt.Sprintf("Invalid backend '%s'", c["dbback"]))
    }

    if err != nil {
        return nil, err
    }

    return t, nil
}

//ReservedTypes returns reserved types used internally by Ting.
func ReservedTypes() []string {
    return []string{
        "string",
        "int",
        "bool",
        "list",
        "array",
    }
}

//ReservedType checks if a type is in the reserved list. Shorthand func.
func ReservedType(t string) bool {
    for _, ty := range ReservedTypes() {
        if t == ty {
            return true
        }
    }
    return false
}

//ValidateContentType makes sure that a raw JSON content type is valid to be turned into a *backend.ContentType.
func (t *Ting) ValidateContentType(name string, structure []byte) (*backend.ContentType, error) {
    var err error

    //Check if the type is reserved.
    if ReservedType(name) {

        //We can't edit the structure of reserved types.
        return nil, errors.New(fmt.Sprintf("'%s' is a reserved content type id.", name))
    }

    s := &backend.ContentType{}

    //Attempt to unmarshal the structure into the *ContentType.
    err = json.Unmarshal(structure, &s.Structure)
    if err != nil {
        return nil, err
    }

    //Attempt to get existing content types.
    resp := t.Backend.ContentTypes()
    if resp.Error != nil {
        return nil, resp.Error
    }

    //Reserved types, existing types and our own name form a slice.
    types := append(resp.Data.([]string), ReservedTypes()...)
    types = append(types, name)

    //Make sure that every field is valid by ensuring every field is in the slice.
    for _, field := range s.Structure {
        found := false
        for _, ty := range types {
            if field.Type == ty {
                found = true
            }
        }

        if !found {
            return nil, errors.New(fmt.Sprintf("Content type '%s' does not exist.", field.Type))
        }
    }

    //Return our content type.
    return s, nil
}