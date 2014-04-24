package main

import (
    "os"
    "fmt"
    "errors"
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

func startProject(name string) {
    if _, err := os.Stat(name); err != nil {
        if !os.IsNotExist(err) {
            panic(err)
        }
    } else {
        panic(errors.New(fmt.Sprintf("File named '%s' already exists in this directory.", name)))
    }
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true, "tingCTL v0.1.0", false)

    if args["startproject"].(bool) {
        startProject(args["<name>"].(string))
        return
    }

    fmt.Println(args)

    if args["addtype"].(bool) {

    }

    if args["deltype"].(bool) {

    }

    if args["modtype"].(bool) {

    }
}
