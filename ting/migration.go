package ting

import (
    "fmt"
    "time"
    "errors"
    "strings"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"
)

var MigrationsDirName string = "migrations"

//A Migration represents a file in the migrations folder.
//Migrations are applied to a database by using the ctl.
type Migration struct {
    Id          string `bson:"_id"          json:"_id"`
    Timestamp   int64  `bson:"timestamp"    json:"timestamp"`
    Filename    string `bson:"filename"     json:"filename"`
    ContentType string `bson:"content_type" json:"content_type"`
    Action      string `bson:"action"       json:"action"`
    Document    string `bson:"document"     json:"document, omitempty"`
}

//IsValid checks whether this is a valid migration.
//Returns a boolean corresponding to whether it's valid.
func (m *Migration) IsValid() bool {
    if len(m.ContentType) == 0 {
        return false
    }

    if len(m.Action) == 0 {
        return false
    }

    return true
}

//IsComplete checks whether this is a complete migration;
//To be considered complete, it needs to be valid,
//and have an Id, Timestamp and Filename.
//Returns a bool if it's a complete migration.
func (m *Migration) IsComplete() bool {
    if !m.IsValid() {
        return false
    }

    if len(m.Id) == 0 {
        return false
    }

    if m.Timestamp == 0 {
        return false
    }

    if len(m.Filename) == 0 {
        return false
    }

    return true
}

func (m *Migration) IsAppliedToMongo(mc *mgo.Collection) (bool, error) {
    ts := strconv.FormatInt(m.Timestamp, 16)

    r := &Migration{}
    err := mc.Find(bson.M{"_id": ts}).One(&r)
    if err != nil {
        return false, nil
    }

    if len(r.Id) > 0 {
        return true, nil
    }

    return false, nil
}

func (m *Migration) ApplyToMongo(mc *mgo.Collection) (bool, error) {
    if !m.IsComplete() {
        return false, errors.New("Cannot apply incomplete migration.")
    }

    a, err := m.IsAppliedToMongo(mc)
    if a || err != nil {
        return false, err
    }

    err = mc.Insert(m)
    if err != nil {
        return false, err
    }

    return true, nil
}

func (m *Migration) ApplyInit(mc *mgo.Collection, c *mgo.Collection) (bool, error) {
    err := c.Create(&mgo.CollectionInfo{})
    if err != nil {
        return false, err
    }

    applied, err := m.ApplyToMongo(mc)
    if err != nil {
        return false, err
    }

    return applied, nil
}

func (m *Migration) ApplyDocument(mc *mgo.Collection, c *mgo.Collection) (bool, error) {
    return false, nil
}

func (m *Migration) ApplyStructure(mc *mgo.Collection, c *mgo.Collection) (bool, error) {
    return false, nil
}

//Serializes the migration to JSON for saving.
//Returns an empty byte slice and an error if it fails,
//or a byte slice of JSON characters and a nil error if successful.
func (m *Migration) Serialize() ([]byte, error) {
    j, err := json.MarshalIndent(m, "", "    ")
    if err != nil {
        return []byte{}, err
    }

    return j, nil
}

//AllMigrations loads all migrations.
//It returns a slice of *Migrations and a nil error if successful,
//Or nil and an error otherwise.
func AllMigrations() ([]*Migration, error) {
    files, err := ioutil.ReadDir(MigrationsDirName)
    if err != nil {
        return nil, err
    }

    m := []*Migration{}

    for _, fn := range files {
        fileData, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", MigrationsDirName, fn.Name()))
        if err != nil {
            return nil, err
        }

        s := &Migration{}

        err = json.Unmarshal(fileData, s)
        if err != nil {
            return nil, err
        }

        m = append(m, s)
    }

    return m, nil
}

//MigrationIndex gets the current number of migrations in ./migrations and adds one to the figure.
//Returns a string padded up to five chars with zeroes and a nil error if successful,
//or a blank string and an error if unsuccessful.
func MigrationIndex() (string, error) {
    files, err := ioutil.ReadDir(MigrationsDirName)
    if err != nil {
        return "", err
    }

    num := strconv.Itoa(len(files)+1)
    return fmt.Sprintf("%s%s",strings.Repeat("0", 5 - len(num)),num), nil
}

//Saving a Migration creates a new file in ./migrations, and saves this Migration to it.
//After a Migration has been saved, it can be applied using the ctl.
//Returns a nil error if successful, or an error otherwise.
func (m *Migration) Save() error {
    if !m.IsValid() {
        return errors.New("invalid migration, couldn't save to file.")
    }

    //Store the timestamp, id and filename.
    idx, err := MigrationIndex()
    if err != nil {
        return err
    }

    m.Timestamp = time.Now().UTC().UnixNano()
    m.Id = strconv.FormatInt(m.Timestamp, 16)
    m.Filename = fmt.Sprintf("%s_%s_%s_%s.json", idx, m.Id, strings.ToLower(m.ContentType), strings.ToLower(m.Action))

    //Write the migration to a file.
    writePath := fmt.Sprintf("%s/%s", MigrationsDirName, m.Filename)
    data, err := m.Serialize()
    if err != nil {
        return err
    }

    err = ioutil.WriteFile(writePath, data, 0755)
    if err != nil {
        return err
    }

    return nil
}