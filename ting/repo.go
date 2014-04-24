package ting 

import (
    "fmt"
    "errors"
    "labix.org/v2/mgo"
)

//Repo is the structure for a site.
//All Repo data is stored within a MongoDB database.
//All Repo settings are stored within a settings file.
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

//RepoFromCwd creates a *Repo instance and sets up a Mongo session and settings file.
//It uses the current working directory to find the settings file.
//It does not connect to the database.
//It returns the *Repo and a nil error on success, or nil and an error on failure.
func RepoFromCwd() (*Repo, error) {
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

    return r, nil
}

//NewRepo sets up a Repo in Mongo from an existing settings file.
//LoadRepo will not work unless the Repo has been set up with NewRepo.
//It returns the *Repo and a nil error on success, or nil and an error on failure.
func NewRepo() (*Repo, error) {
    r, err := RepoFromCwd()
    if err != nil {
        return nil, err
    }
    defer r.Session.Close()
    r.Db = r.Session.DB(r.Settings.MongoDb)
    

    return r, nil
}

//LoadRepo initialises a Repo instance based on the current directory.
//It attempts to load a settings file into a Settings instance,
//and attempts to connect to the mongoDB instance specified by the settings.
//It returns a *Repo and a nil error on success, or nil and an error on fail.
func LoadRepo() (*Repo, error) {
    r, err := RepoFromCwd()
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

//AddContentType adds a content type to this Repo.
//It adds the migration file and also adds the type to Mongo.
//It returns an error if it was not successful.
func (r *Repo) AddContentType(name string) error {
    fmt.Println(fmt.Sprintf("Added content type named '%s'.", name))
    return nil
}

//DeleteContentType deletes a content type to this Repo.
//It adds the migration file and also deletes the type from Mongo.
//It returns an error if it was not successful.
func (r *Repo) DeleteContentType(name string) error {
    fmt.Println(fmt.Sprintf("Removed content type named '%s'.", name))
    return nil
}