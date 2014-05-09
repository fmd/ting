package mongo

import (
    "github.com/fmd/ting/response"
)

func (r *Repo) PushContent(content []byte) *response.R {
    return &response.R{}
}

func (r *Repo) Content(contentType string, id string) *response.R {
    return &response.R{}
}

func (r *Repo) Contents(contentType string, query interface{}) *response.R {
    return &response.R{}
}
