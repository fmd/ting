package backend

import (
	"github.com/fmd/ting/backend/response"
)

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