package ting

import (
    "labix.org/v2/mgo"
)

//A Repo represents a Mongo session and a database to act upon.
type Repo struct {
    Session *mgo.Session
    Db *mgo.Database
}

//NewMgoSession creates a new Mgo session.
//BUG(Needs to include other credentials for the Mongo session)
//It returns the session and a nil error if successful, or nil and an error if unsuccessful.
func NewSession(hostname string) (*mgo.Session, error) {
    s, err := mgo.Dial(hostname)
    if err != nil {
        return nil, err
    }
    s.SetMode(mgo.Monotonic, true)
    return s, nil
}

//NewRepo creates a *Repo instance.
//BUG(Needs to include other credentials from the Mongo session)
//It returns a nil *Repo and an error if unsuccessful, or a *Repo and a nil error otherwise.
func NewRepo(hostname string, db string) (*Repo, error) {
    var err error
    r := &Repo{}
    r.Session, err = NewSession(hostname)
    if err != nil {
        return nil, err
    }

    r.Db = r.Session.DB(db)
    return r, nil
}