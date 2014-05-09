package mongo

import (
    "github.com/fmd/ting/backend"
)

func (r *Repo) PushContent(content *backend.Content) error {
    return nil
}

func (r *Repo) Content(contentType string, id string) (interface{}, error) {
    return contentType, nil
}

func (r *Repo) Contents(contentType string, query interface{}) ([]interface{}, error) {
    return nil, nil
}