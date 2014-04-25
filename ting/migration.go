package ting

import (
    "fmt"
    "time"
    "errors"
    "strings"
    "strconv"
    "io/ioutil"
    "encoding/json"
)

var MigrationsDirName string = "migrations"

//A Migration represents a file in the migrations folder.
//Migrations are applied to a database by using the ctl.
type Migration struct {
    Timestamp   int64
    ContentType string
    Action      string
}

//Checks whether this is a valid migration.
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

//Saving a Migration creates a new file in ./migrations, and saves this Migration to it.
//After a Migration has been saved, it can be applied using the ctl.
func (m *Migration) Save() error {
    if !m.IsValid() {
        return errors.New("Cannot save: invalid migration.")
    }

    m.Timestamp = time.Now().UTC().Unix()
    data, err := m.Serialize()
    if err != nil {
        return err
    }

    filename := fmt.Sprintf("%s_%s_%s",strconv.FormatInt(m.Timestamp, 10),strings.ToLower(m.ContentType),strings.ToLower(m.Action))
    writePath := fmt.Sprintf("%s/%s",MigrationsDirName,filename)

    err = ioutil.WriteFile(writePath, data, 0755)
    if err != nil {
        return err
    }

    return nil
}