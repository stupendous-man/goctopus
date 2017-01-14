//TODO: Change to a shared library
package main

//TODO: Remove mgo dependency
import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

//Created by stupendous-man
func main() {

	//Establish session with Mongo server
	//TODO: Replace localhost with env variable set in docker container
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		//Retry a few times before going into panic
		log.Print(err)
		for retries := 1; retries <= 5; retries++ {
			log.Printf("Retrying Mongo connection. Attempt %d...", retries)
			session, err = mgo.Dial("localhost")

			if err != nil {
				time.Sleep(100 * time.Millisecond)
			} else {
				break
			}
		}
	}

	if err != nil {
		panic(err)
	}

	//Deferred function to close session
	defer session.Close()

	//Establish connection with database/collection
	c := session.DB("goctupusdb").C("messages")

	//Insert document in database/collection
	//TODO: Move to separate function taking document as input parameter
	err = c.Insert(&Document{"Not all who wander are lost..."})

	if err != nil {
		log.Fatal(err)
	}
}

//TODO: Expand document structure
type Document struct {
	Message string
}

//TODO: Create CRUD functions
