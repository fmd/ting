package main

import (
	"flag"
	"github.com/go-martini/martini"
	"github.com/fmd/goting/ting"
	"labix.org/v2/mgo"
)

var mgoHost *string = flag.String("host", "localhost", "MongoDB hostname string.")
var mgoDb *string = flag.String("db", "test", "MongoDB database to connect to.")

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
	flag.Parse()
	session := getSession(*mgoHost)
	defer session.Close()
	db := session.DB(*mgoDb)
	r := &ting.Repo{db}

	//Check repo for collection errors
    if err := r.CollectionError(); err != nil {
		panic(err)
	}

	//Create martini and serve
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}