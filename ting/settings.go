package ting

import (
    "io/ioutil"
)

//This is the struct that contains the schema for a settings file.
//Settings files should *always be ignored by your VCS!
type Settings struct {

    //Mongo Credentials
    MongoHost string `json:"mongo_host"`
    MongoDb   string `json:"mongo_db"`
    MongoUser string `json:"mongo_user"`
    MongoPass string `json:"mongo_pass"`

    //General settings
    ProjectName     string `json:"project_name"`
    ProjectVersion  string `json:"project_version"`
}

//This function creates a new instance of Settings.
//It fills out the instance with useful defaults.
//It returns the created instance.
func NewSettings() *Settings {
    s := &Settings{}
    
    //Mongo Credentials
    s.MongoHost = "localhost"
    s.MongoDb = ""
    s.MongoUser = ""
    s.MongoPass = ""

    //General Settings
    s.ProjectName = ""
    s.ProjectVersion = "0.0.0"

    return s
}

func (s *Settings) Save() err {
    j, err := json.MarshalIndent(s, "", "    ")

    if err != nil {
        return err
    }
    
    err = ioutil.WriteFile("settings.json", j, 0755)

    if err != nil {
        return err
    }

    return nil
}