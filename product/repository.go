package product

import (
	"github.com/globalsign/mgo"
)

type Repository struct {
	dbSession *mgo.Session
	dbname    string
	tableName string
}
func (r *Repository) GetDBrepository(){
return &Repository{

	dbSession: r.dbSession,
	dbname: "testDB",
	tableName: "testTable"
}
}