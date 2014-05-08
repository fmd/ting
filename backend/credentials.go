package backend

// ["dbback"] : mongodb | (couchdb)
type Credentials map[string]string

func NewCredentials() Credentials {
	c := make(Credentials)
	c["dbback"] = "mongodb"   //Database backend
	c["dbhost"] = "localhost" //Database host
	c["dbname"] = ""          //Database name
	c["dbuser"] = ""          //Database user
	c["dbpass"] = ""          //Database password

	return c
}
