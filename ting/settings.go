package ting

import (
    "encoding/json"
    "io/ioutil"
    "reflect"
    "fmt"
)

//Settings files are called "settings.json".
var settingsFilename string = "settings.json"

//Settings the struct that contains the schema for a settings file.
//Settings files should *always* be ignored by your VCS!
//All Settings fields should be stored as strings.
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

//NewSettings creates a new instance of Settings.
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

//LoadSettings loads settings from a file.
//It returns a pointer to a Settings instance if successful (nil otherwise),
//As well as an error if it fails (nil otherwise).
func LoadSettings() (*Settings, error) {
    data, err := ioutil.ReadFile(settingsFilename)
    if err != nil {
        return nil, err
    }

    s := &Settings{}
    err = json.Unmarshal(data, s)
    if err != nil {
        return nil, err
    }

    return s, nil
}

//Save is called on a Settings instance to save the settings in json format.
//The data is saved to a file in the current directory by the name of the filename string.
//It returns an error if it fails, or a nil error otherwise.
func (s *Settings) Save() error {
    j, err := json.MarshalIndent(s, "", "    ")

    if err != nil {
        return err
    }
    
    err = ioutil.WriteFile(settingsFilename, j, 0755)

    if err != nil {
        return err
    }

    return nil
}

//FieldNameByJsonTag get a fieldname (obtained via reflection) matching a given tag.
//The tag is passed in as a string.
//It returns an empty string if the tag could not be found, or the name of the field if it could.
func (s *Settings) FieldNameByJsonTag(tag string) string {
    if len(tag) == 0 {
        return ""
    }

    el := reflect.TypeOf(s).Elem()
    numFields := el.NumField()
    for i := 0; i < numFields; i++ {
        field := el.Field(i)
        fmt.Println(field.Tag)
        if field.Tag.Get("json") == tag {
            return field.Name
        }
    }

    return ""
}