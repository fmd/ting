package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/docopt/docopt-go"
)

var version string = "v0.1.0"

func workingDir() string {
    dir, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    split := strings.Split(dir, "/")
    return split[len(split)-1]
}

func usage() string {
    return fmt.Sprintf(`tingd.

    Usage:
        tingctl [--host=<hostname>] [--db=<dbname]
        tingctl --help
        tingctl --version

    Options:
        -d | --db       MongoDB database string [default: %s].
        -h | --host     MongoDB host string [default: localhost].
        --help          Show this screen.
        --version       Show version.`, workingDir())
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true,fmt.Sprintf("tingd %s", version), false)
    fmt.Println(args)
}