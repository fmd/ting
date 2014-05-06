package credentials

// ["dbback"] : mongodb | (couchdb)
// ["dbhost"] : localhost
// ["dbname"] : ""
// ["dbuser"] : ""
// ["dbpass"] : ""
type Credentials map[string]string

func NewCredentials() Credentials {
	c := make(Credentials)
	c["dbback"] = "mongodb"
	c["dbhost"] = "localhost"
	c["dbname"] = ""
	c["dbuser"] = ""
	c["dbpass"] = ""

	return c
}
