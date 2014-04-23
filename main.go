package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

var mgoHost *string = flag.String("host", "localhost", "MongoDB host string.")
var mgoDb *string = flag.String("db", "test", "MongoDB database to connect to.")

func getRequiredCollections() map[string]bool {
	return map[string]bool{
		"contentTypes": true,
		"admins":       true,
	}
}

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
	session := getSession(*mgoHost)
	defer session.Close()

	//Connect to the database
	d := session.DB(*mgoDb)

	//Attempt to get collection names. If we can't, we connected to a bad DB
	collectionNames, err := d.CollectionNames()

	if err != nil {
		panic(err)
	}

	req := getRequiredCollections()

	//Make sure that have all required collections
	for _, el := range collectionNames {
		if val, ok := req[el]; val && ok {
			req[el] = false
		}
	}

	//If we don't have one of the required collections, panic
	for idx, notFound := range req {
		if notFound {
			errStr := fmt.Sprintf("Database '%s' does not have required collection '%s'.", d.Name, idx)
			err = errors.New(errStr)
			break
		}
	}

	if err != nil {
		panic(err)
	}

	//Create martini and serve
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
