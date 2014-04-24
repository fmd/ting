package ting 

import (
    "labix.org/v2/mgo"
    "fmt"
    "errors"
)

//Repo is the MongoDB structure for a site.
//All Repo data is stored within a MongoDB database.
type Repo struct {
    Db *mgo.Database
}

//RequiredCollections gets a map of required collections for a Repo.
//Any further required collections for Repos should be added here in the format "collection": true
//It returns a map[string]bool of all collections required for a Repo to be considered valid.
func RequiredCollections() map[string]bool {
    return map[string]bool{
        "contentTypes": true,
        "admins":       true,
    }
}

//CollectionError checks to see whether a Repo possesses all required collections.
//It uses RequiredCollections() to get a list of the required collections.
//It returns an error if unsuccessful, and nil otherwise.
func (r *Repo) CollectionError() error {
    names, err := r.Db.CollectionNames()
    
    if err != nil {
        return err
    }

    req := RequiredCollections()

    for _, el := range names {
        if val, ok := req[el]; val && ok {
            req[el] = false
        }
    }

    for idx, notFound := range req {
        if notFound {
            errStr := fmt.Sprintf("Database '%s' does not have required collection '%s'.", r.Db.Name, idx)
            return errors.New(errStr)
        }
    }

    return nil
}
