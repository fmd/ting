package main

import (
    "os"
    "fmt"
    "errors"
    "github.com/fmd/goting/ting"
)

func mkProjectDir(name string) error {

    //Ensure nothing with this name already exists here.
    if _, err := os.Stat(name); err != nil {
        if !os.IsNotExist(err) {
            return err
        }
    } else {
        return errors.New(fmt.Sprintf("File named '%s' already exists in this directory.", name))
    }

    //Create the directory.
    if err := os.Mkdir(name, 0755); err != nil {
        return err
    }

    return nil
}

func chProjectDir(name string) error {
    d, err := os.Open(name)

    if err != nil {
        return err
    }

    defer d.Close()

    err = d.Chdir()

    if err != nil {
        return err
    }

    return nil
}

func startProject(name string) error {
    var err error

    //Make the project directory.
    if err = mkProjectDir(name); err != nil {
        return err
    }

    //Change directory to newly created project directory.
    if err = chProjectDir(name); err != nil {
        return err
    }

    //Create the migrations directory.
    if err = mkProjectDir(ting.MigrationsDirName); err != nil {
        return err
    }

    //Create settings struct instance and save to file.
    s := ting.NewSettings()
    s.MongoDb = name
    s.ProjectName = name
    err = s.Save()

    if err != nil {
        return err
    }

    _, err = ting.NewRepo()
    if err != nil {
        return err
    }

    return nil
}
