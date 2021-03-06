package ting

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/fmd/ting/backend"
    "github.com/fmd/ting/backend/mongo"
    "github.com/fmd/ting/response"
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
    s.Id = name

    //Attempt to unmarshal the structure into the *ContentType.
    err = json.Unmarshal(structure, &s.Structure)
    if err != nil {
        return nil, err
    }

    //Attempt to get existing content types.
    types, err := t.Backend.ContentTypes()
    if err != nil {
        return nil, err
    }

    //Reserved types, existing types and our own name form a slice.
    types = append(types, ReservedTypes()...)
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

//BUG(Still needs to actually validate against its own structure)
//ValidateContentType makes sure that raw JSON content is valid to be turned into a *backend.Content.
func (t *Ting) ValidateContent(contentType string, id interface{}, content []byte) (*backend.Content, error) {
    var err error

    //Check if the type is reserved.
    if ReservedType(contentType) {

        //We can't push content of reserved types.
        return nil, errors.New(fmt.Sprintf("'%s' is a reserved content type.", contentType))
    }

    c := &backend.Content{}

    c.Id = id

    err = json.Unmarshal(content, &c.Content)
    if err != nil {
        return nil, err
    }

    //Attempt to get existing content types.
    types, err := t.Backend.ContentTypes()
    if err != nil {
        return nil, err
    }

    //Make sure the content type we're trying to push actually exists.
    found := false
    for _, field := range types {
        if contentType == field {
            found = true
            break
        }
    }

    if !found {
        return nil, errors.New(fmt.Sprintf("Content type '%s' does not exist.", contentType))
    }

    //Reserved types, existing types and our own name form a slice.
    types = append(types, ReservedTypes()...)
    types = append(types, contentType)

    //Make sure that every field is valid by ensuring every field is in the slice.
    for _, field := range c.Content {
        found := false
        for _, ty := range types {
            if field.Type == ty {
                found = true
                break
            }
        }

        if !found {
            return nil, errors.New(fmt.Sprintf("Content type '%s' does not exist.", field.Type))
        }
    }

    return c, nil
}

//UpsertContent inserts or updates a piece of content based on its type.
func (t *Ting) PushContent(contentType string, id interface{}, content []byte) (int, response.JSend) {
    c, err := t.ValidateContent(contentType, id, content)
    if err != nil {
        return response.Error(err).Wrap()
    }

    err = t.Backend.PushContent(contentType, c)
    if err != nil {
        return response.Error(err).Wrap()
    }

    return response.Success(nil).Wrap()
}

//Content uses an id to get a piece of content based on its type.
func (t *Ting) Content(contentType string, id string) (int, response.JSend) {
    content, err := t.Backend.Content(contentType, id)
    if err != nil {
        return response.Error(err).Wrap()
    }

    return response.Success(content).Wrap()
}

//Contents gets multiple pieces of content based on a query and a content type.
func (t *Ting) Contents(contentType string, query interface{}) (int, response.JSend) {
    contents, err := t.Backend.Contents(contentType, query)
    if err != nil {
        return response.Error(err).Wrap()
    }

    return response.Success(contents).Wrap()
}

//StructureType uses serialized JSON to update the CMS structure of a content type.
func (t *Ting) PushContentType(name string, body []byte) (int, response.JSend) {
    c, err := t.ValidateContentType(name, body)
    if err != nil {
        return response.Error(err).Wrap()
    }

    err = t.Backend.PushContentType(c)
    if err != nil {
        return response.Error(err).Wrap()
    }

    return response.Success(nil).Wrap()
}

//ContentTypes gets a list of all available content backend.
func (t *Ting) ContentTypes() (int, response.JSend) {
    types, err := t.Backend.ContentTypes()
    if err != nil {
        return response.Error(err).Wrap()
    }

    return response.Success(types).Wrap()
}

//ContentType gets the structure of a content type by its name.
func (t *Ting) ContentType(name string) (int, response.JSend) {
    ty, err := t.Backend.ContentType(name)
    if err != nil {
        return response.Error(err).Wrap()
    }

    return response.Success(ty).Wrap()
}
