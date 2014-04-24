package ting 

import (
    "fmt"
    "errors"
    "labix.org/v2/mgo"
)

//Repo is the MongoDB structure for a site.
//All Repo data is stored within a MongoDB database.
type Repo struct {
    Session *mgo.Session
    Db *mgo.Database
    Settings *Settings
}

//RequiredCollections gets a map of required collections for a Repo.
//Any further required collections for Repos should be added here in the format "collection": true
//It returns a map[string]bool of all collections required for a Repo to be considered valid.
func RequiredCollections() map[string]bool {
    return map[string]bool{
        "contentTypes": true,
        "migrations":   true,
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

//NewMgoSession creates a new Mgo session.
//BUG(Needs to include other credentials from the Repo session)
//It returns the session and a nil error if successful, or nil and an error if unsuccessful.
func NewMgoSession(hostname string) (*mgo.Session, error) {
    s, err := mgo.Dial(hostname)
    if err != nil {
        return nil, err
    }

    s.SetMode(mgo.Monotonic, true)
    return s, nil
}

//LoadRepo initialises a Repo instance based on the current directory.
//It attempts to load a settings file into a Settings instance,
//and attempts to connect to the mongoDB instance specified by the settings.
//It returns a *Repo and a nil error on success, or nil and an error on fail.
func LoadRepo() (*Repo, error) {
    var err error

    r := &Repo{}
    r.Settings, err = LoadSettings()
    if err != nil {
        return nil, err
    }

    r.Session, err = NewMgoSession(r.Settings.MongoHost)
    if err != nil {
        return nil, err
    }

    defer r.Session.Close()
    r.Db = r.Session.DB(r.Settings.MongoDb)

    if err = r.CollectionError(); err != nil {
        return nil, err
    }

    return r, nil
}

func (r *Repo) AddContentType(name string) error {
    fmt.Println(fmt.Sprintf("Added content type named '%s'.", name))
    return nil
}

func (r *Repo) DeleteContentType(name string) error {
    fmt.Println(fmt.Sprintf("Removed content type named '%s'.", name))
    return nil
}