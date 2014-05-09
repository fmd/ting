package main

import (
    "fmt"
    "github.com/docopt/docopt-go"
    "os"
    "strings"
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
    return fmt.Sprintf(`tingctl.

        Usage:
            tingctl startproject <name>
            tingctl contenttypes 
            tingctl --help
            tingctl --version

        Options:
            --help          Show this screen.
            --version       Show version.`, workingDir())
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true, fmt.Sprintf("tingctl %s", version), false)
    fmt.Println(args)
}
