package main

import (
    "labix.org/v2/mgo"
    "fmt"
    "errors"
)

type Repo struct {
    Db *mgo.Database
}

func RequiredCollections() map[string]bool {
    return map[string]bool{
        "contentTypes": true,
        "admins":       true,
    }
}

func (r *Repo) CollectionError() error {
    names, err := r.Db.CollectionNames()
    
    if err != nil {
        return err
    }

    req := RequiredCollections()

    //Make sure that have all required collections
    for _, el := range names {
        if val, ok := req[el]; val && ok {
            req[el] = false
        }
    }

    //If we don't have one of the required collections, panic
    for idx, notFound := range req {
        if notFound {
            errStr := fmt.Sprintf("Database '%s' does not have required collection '%s'.", r.Db.Name, idx)
            return errors.New(errStr)
        }
    }

    return nil
}