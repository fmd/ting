package main

import (
    "os"
    "fmt"
    "errors"
    "reflect"
    "github.com/fmd/goting/ting"
    "github.com/docopt/docopt-go"
)

func usage() string {
    return `tingCTL.

Usage:
    tingctl startproject <name>
    tingctl addtype <name>
    tingctl deltype <name>
    tingctl modtype <type> addprim <prim>
    tingctl modtype <type> delprim <prim>
    tingctl settings get <key>
    tingctl settings set <key> <value>
    tingctl -h | --help
    tingctl --version

Options:
    -h --help   Show this screen.
    --version   Show version.`
}

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
    if err = mkProjectDir("migrations"); err != nil {
        return err
    }

    //Create settings struct instance and save to file.
    s := ting.NewSettings()
    err = s.Save()

    if err != nil {
        return err
    }

    return nil
}

func settingsSet(key string, value string) error {
    s, err := ting.LoadSettings()
    if err != nil {
        return err
    }

    fieldname := s.FieldNameByJsonTag(key)

    if len(fieldname) == 0 {
        return errors.New(fmt.Sprintf("Could not find field in settings with key '%s'", key))
    }

    reflect.ValueOf(s).Elem().FieldByName(fieldname).SetString(value)

    if err = s.Save(); err != nil {
        panic(err)
    }

    return nil
}

func settingsGet(key string) (string, error) {
    return "", nil
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true, "tingCTL v0.1.0", false)

    if args["startproject"].(bool) {
        if err := startProject(args["<name>"].(string)); err != nil {
            panic(err)
        }
        return
    }

    if args["addtype"].(bool) {

    }

    if args["deltype"].(bool) {

    }

    if args["modtype"].(bool) {

    }

    if args["settings"].(bool) {
        if args["set"].(bool) {
            if err := settingsSet(args["<key>"].(string), args["<value>"].(string)); err != nil {
                panic(err)
            }
            return
        } 

        if args["get"].(bool) {
            s, err := settingsGet(args["<key>"].(string))
            if err != nil {
                panic(err)
            }

            fmt.Println(s)
            return
        }
    }
}
