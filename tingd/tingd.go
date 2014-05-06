package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
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
            tingd [--host=<hostname>] [--db=<dbname>] [--port=<port>]
            tingd --help
            tingd --version

        Options:
            -h | --host     DB host string [default: localhost].
            -d | --db       DB database string [default: %s].
            -p | --port     The port to listen on [default: 5000].
            --help          Show this screen.
            --version       Show version.`, workingDir())
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true,fmt.Sprintf("tingd %s", version), false)

    host := args["--host"].(string)
    db := args["--db"].(string)
    port, err := strconv.Atoi(args["--port"].(string))
    if err != nil {
        panic(err)
    }

    d, err := NewDaemon(host, db, port)
    if err != nil {
        panic(err)
    }

    err = d.ListenAndServe()
    if err != nil {
        panic(err)
    }
}