package main

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo"
)

/*type Hello struct {
	Xy string //sr
}*/

func (db *mgo.Session) addProducHandler(write http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(write, db.DB())
	//fmt.Println(req.Url)
}

func main() {

	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		fmt.Println("Cannot connect to database......")
		return
	}

	fmt.Println("Server Started.......")

	http.HandleFunc("/anexa/addProduct", session.addProducHandler)
	http.ListenAndServe(":8080", nil)

}
