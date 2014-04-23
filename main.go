package main

import (
            "github.com/go-martini/martini"
            "labix.org/v2/mgo"
            "flag"
)

var mongoHost *string = flag.String("host","localhost","MongoDB host string")

func getSession(host string) *mgo.Session {
    //Attempt to dial host
    s, err := mgo.Dial(host)

    //If we can't connect, panic
    if err != nil {
        panic(err)
    }

    //At this point we're fine
    s.SetMode(mgo.Monotonic, true)
    return s
}

func main() {

    //Parse flags
    flag.Parse()

    //Create mongo session
    session := getSession(*mongoHost)
    defer session.Close()

    //Create martini and serve
    m := martini.Classic()
    m.Get("/", func() string {
        return "Hello world!"
    })
    m.Run()
}
