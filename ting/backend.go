package ting

type Backend interface {

    //Data
    UpsertContent(contentType string, content interface{})
    GetContent(contentType string, id string)
    GetContents(contentType string, query interface{})

    //Types
    StructureType(contentType string)
    ContentTypes()
}