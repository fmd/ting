package ting

import (
    "time"
    "errors"
)

//A Migration represents a file in the migrations folder.
//Migrations are applied to a database by using the ctl.
type Migration struct {
    Timestamp   int64
    ContentType string
    Action      string
}

//Saving a Migration creates a new file in ./migrations, and saves this Migration to it.
//After a Migration has been saved, it can be applied using the ctl.
func (m *Migration) Save() error {
    if m.ContentType == "" {
        return errors.New("Cannot save a migration that has no ContentType.")
    }

    if m.Action == "" {
        return errors.New("Cannot save a migration that has no Action.")
    }

    m.Timestamp = time.Now().UTC().Unix()
    return nil
}