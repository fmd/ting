package backend

import (
	"github.com/fmd/ting/backend/response"
)

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

//ContentType is the Content Type struct
type ContentType struct {
	Id        string      `bson:"_id" json:"_id"`
	Structure interface{} `bson:"structure" json:"structure"`
}

//Backend is the interface that connects to one of our supported backends.
type Backend interface {

	//UpsertContent inserts or updates a piece of content based on its type.
	PushContent(content []byte) *response.R

	//GetContent uses an id to get a piece of content based on its type.
	Content(contentType string, id string) *response.R

	//GetContents gets multiple pieces of content based on a query and a content type.
	Contents(contentType string, query interface{}) *response.R

	//StructureType uses serialized JSON to update the CMS structure of a content type.
	PushType(structure []byte) *response.R

	//ContentTypes gets a list of all available content backend.
	ContentTypes() *response.R
}
