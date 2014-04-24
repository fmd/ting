package main

import (
    "fmt"
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

func main() {
    args, _ := docopt.Parse(usage(), nil, true, "tingCTL v0.1.0", false)

    if args["startproject"].(bool) {
        if err := startProject(args["<name>"].(string)); err != nil {
            panic(err)
        }

        return
    }

    if args["settings"].(bool) {
        if args["set"].(bool) {
            if err := settingsSet(args["<key>"].(string), args["<value>"].(string)); err != nil {
                panic(err)
            }
            
        } else if args["get"].(bool) {
            s := settingsGet(args["<key>"].(string))
            if len(s) > 0 {
                fmt.Println(s)
            }
        }

        return
    }

    //The following functions all require a loaded Repo.
    r, err := ting.LoadRepo()
    if err != nil {
        panic(err)
    }

    if args["addtype"].(bool) {
        r.AddContentType(args["<name>"].(string))
    }

    if args["deltype"].(bool) {
        r.DeleteContentType(args["<name>"].(string))
    }

    if args["modtype"].(bool) {

    }
}