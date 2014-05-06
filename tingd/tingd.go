package main

import (
    "os"
    "fmt"
    "strings"
    "github.com/fitstar/falcore"
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
            tingd [--host=<hostname>] [--db=<dbname]
            tingd --help
            tingd --version

        Options:
            -h | --host     MongoDB host string [default: localhost].
            -d | --db       MongoDB database string [default: %s].
            -p | --port     The port to listen on [default: 8000].
            --help          Show this screen.
            --version       Show version.`, workingDir())
}

func initFalcore(port int) {

}

func main() {
    args, _ := docopt.Parse(usage(), nil, true,fmt.Sprintf("tingd %s", version), false)
    if
}