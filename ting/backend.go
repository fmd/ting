package ting

type Backend interface {

	//UpsertContent inserts or updates a piece of content based on its type.
	UpsertContent(contentType string, content interface{})

	//GetContent uses an id to get a piece of content based on its type.
	GetContent(contentType string, id string)

	//GetContents gets multiple pieces of content based on a query and a content type.
	GetContents(contentType string, query interface{})

	//StructureType uses serialized JSON to update the CMS structure of a content type.
	StructureType(structure []byte) error

	//ContentTypes gets a list of all available content types.
	ContentTypes() ([]string, error)
}
