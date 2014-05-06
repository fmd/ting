package main

import (
    "os"
    "fmt"
    "strings"
    "strconv"
    "net/http"
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
            tingd [--host=<hostname>] [--db=<dbname>] [--port=<port>]
            tingd --help
            tingd --version

        Options:
            -h | --host     MongoDB host string [default: localhost].
            -d | --db       MongoDB database string [default: %s].
            -p | --port     The port to listen on [default: 5000].
            --help          Show this screen.
            --version       Show version.`, workingDir())
}

func pipeline() *falcore.Pipeline {
    p := falcore.NewPipeline()
    p.Upstream.PushBack(helloFilter)
    return p
}

func server(port int, pipeline *falcore.Pipeline) *falcore.Server {
    return falcore.NewServer(port, pipeline)
}

func main() {
    args, _ := docopt.Parse(usage(), nil, true,fmt.Sprintf("tingd %s", version), false)
    port, err := strconv.Atoi(args["--port"].(string))
    if err != nil {
        panic(err)
    }
    server := server(port, pipeline())
    if err := server.ListenAndServe(); err != nil {
        panic(err)
    }
}

var helloFilter = falcore.NewRequestFilter(func(req *falcore.Request) *http.Response {
    return falcore.StringResponse(req.HttpRequest, 200, nil, "hello world!")
})