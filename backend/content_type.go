package backend

type ContentType struct {
    Id string             `bson:"_id" json:"_id"`
    Structure interface{} `bson:"structure" json:"structure"`
}