package main

import (
    "os"
    "fmt"
    "io/ioutil"
    "errors"
    "encoding/json"
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
    tingctl -h | --help
    tingctl --version

Options:
    -h --help   Show this screen.
    --version   Show version.`
}

func startProject(name string) error {
    if _, err := os.Stat(name); err != nil {
        if !os.IsNotExist(err) {
            return err
        }
    } else {
        return errors.New(fmt.Sprintf("File named '%s' already exists in this directory.", name))
    }

    if err := os.Mkdir(name, 0755); err != nil {
        return err
    }

    //Change directory to newly created project directory
    d, err := os.Open(name)

    if err != nil {
        return err
    }

    defer d.Close()

    err = d.Chdir()

    if err != nil {
        return err
    }

    //Create settings struct instance and save to file
    s := ting.NewSettings()
    err = s.Save()

    if err != nil {
        return err
    }

    return nil
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
}
