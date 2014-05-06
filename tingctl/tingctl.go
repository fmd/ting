package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/docopt/docopt-go"
)

func workingDir() string {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    split := strings.Split(dir, "/")
    return split[len(split)-1]
}

func usage() string {
    return fmt.Sprintf(`tingCTL.

Usage:
    tingctl startproject <name>
    tingctl --help
    tingctl --version

Options:
    -h | --host MongoDB host string [default: 127.0.0.1].
    -d | --db   MongoDB database string [default: %s].
    --help      Show this screen.
    --version   Show version.`, workingDir())
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true, "gomictl v0.1.0", false)
    fmt.Println(args)
}