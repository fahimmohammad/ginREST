package auth

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/haquenafeem/boilerplate-gin/configurations"
)

type repoInterface interface {
	CreateUser(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type repoStruct struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repoStruct) CreateUser(user User) (User, error) {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Insert(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *repoStruct) FindByEmail(email string) (User, error) {
	var user User
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"email": email}).One(&user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func NewRepository(dbSession *mgo.Session) *repoStruct {
	return &repoStruct{
		DBSession: dbSession,
		DBName:    configurations.DBName,
		DBTable:   configurations.AuthTable,
	}
}
