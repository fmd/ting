package ting

type Backend interface {
    //Basic
    Init(credentials map[string]string)

    //Data
    UpsertContent(contentType string, content interface{})
    GetContent(contentType string, id string)
    GetContents(contentType string, query interface{})

    //Types
    Structure(contentType string)
    ContentTypes()

    //Migrations
    Migrate()
    Rollback()
}