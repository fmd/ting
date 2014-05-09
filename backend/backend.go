package backend

//Credentials is the credentials, shared across backends.
// ["dbback"] : mongodb | (couchdb)
type Credentials map[string]string

func NewCredentials() Credentials {
    c := make(Credentials)
    c["dbback"] = "mongodb"   //Database backend
    c["dbhost"] = "localhost" //Database host
    c["dbname"] = ""          //Database name
    c["dbuser"] = ""          //Database user
    c["dbpass"] = ""          //Database password

    return c
}

//ContentTypeField is the Content Type Field struct
type ContentTypeField struct {
    Type     string      `bson:"type" json:"type"`
    Required bool        `bson:"required" json:"required"`
    Default  interface{} `bson:"default" json:"default"`
}

//ContentType is the Content Type struct
type ContentType struct {
    Id        string                      `bson:"_id" json:"_id"`
    Structure map[string]ContentTypeField `bson:"structure" json:"structure"`
}

//ContentField is the Content Field struct
type ContentField struct {
    Type string `bson:"type" json:"type"`
    Data string `bson:"data" json:"data"`
}

//Content is the Content struct
type Content struct {
    Id      interface{}             `bson:"_id" json:"_id"`
    Content map[string]ContentField `bson:"content" json:"content"`
}

//B is the interface that connects to one of our supported backends.
type B interface {

    //UpsertContent inserts or updates a piece of content based on its type.
    PushContent(contentType string, content *Content) error

    //Content uses an id to get a piece of content based on its type.
    Content(contentType string, id string) (interface{}, error)

    //Contents gets multiple pieces of content based on a query and a content type.
    Contents(contentType string, query interface{}) ([]interface{}, error)

    //StructureType uses serialized JSON to update the CMS structure of a content type.
    PushContentType(*ContentType) error

    //ContentTypes gets a list of all available content backend.
    ContentTypes() ([]string, error)

    //ContentType gets the structure of a content type by its name.
    ContentType(name string) (interface{}, error)
}
