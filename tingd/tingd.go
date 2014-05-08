package main

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"github.com/fmd/ting/backend"
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
	return fmt.Sprintf(`tingd.

        Usage:
            tingd [--backend=<backend>] [--host=<hostname>] [--db=<dbname>] [--port=<port>]
            tingd --help
            tingd --version

        Options:
            -b | --backend  DB Backend [default: mongodb].
            -h | --host     DB host string [default: localhost].
            -d | --db       DB name string [default: %s].
            -p | --port     The port to listen on [default: 5000].
            --help          Show this screen.
            --version       Show version.`, workingDir())
}

func main() {
	var err error
    args, _ := docopt.Parse(usage(), nil, true, fmt.Sprintf("tingd %s", version), false)

	c := backend.NewCredentials()

	c["dbback"] = args["--backend"].(string)
	c["dbhost"] = args["--host"].(string)
	c["dbname"] = args["--db"].(string)

	port := args["--port"].(string)
	if err != nil {
		panic(err)
	}

	d, err := NewDaemon(port, c)
	if err != nil {
		panic(err)
	}

	err = d.Run()
	if err != nil {
		panic(err)
	}
}
