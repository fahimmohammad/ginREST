package todo

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/haquenafeem/boilerplate-gin/configurations"
)

type repositoryInterface interface {
	createTodo(todo Todo) (Todo, error)
	readAllTodo() ([]Todo, error)
	readSingleTodo(todoID string) (Todo, error)
	deleteTodo(todoID string) error
	updateTodo(todoID string, todo Todo) (Todo, error)
}

type repository struct {
	DBSession *mgo.Session
	DBName    string
	DBTable   string
}

func (r *repository) createTodo(todo Todo) (Todo, error) {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Insert(&todo)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (r *repository) readAllTodo() ([]Todo, error) {
	var todo []Todo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{}).All(&todo)
	if err != nil {
		return []Todo{}, err
	}
	return todo, nil
}

func (r *repository) readSingleTodo(todoID string) (Todo, error) {
	var todo Todo
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Find(bson.M{"_id": todoID}).One(&todo)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

func (r *repository) deleteTodo(todoID string) error {
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	err := coll.Remove(bson.M{"_id": todoID})
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) updateTodo(todoID string, todo Todo) (Todo, error) {
	todo.ID = todoID
	coll := r.DBSession.DB(r.DBName).C(r.DBTable)
	selector := bson.M{"_id": todoID}
	err := coll.Update(selector, bson.M{"$set": todo})
	if err != nil {
		return Todo{}, err
	}
	updatedTodo, err := r.readSingleTodo(todoID)
	if err != nil {
		return Todo{}, err
	}
	return updatedTodo, nil
}

func newRepositoryService(dbSession *mgo.Session) *repository {
	return &repository{
		DBSession: dbSession,
		DBName:    configurations.DBName,
		DBTable:   configurations.TodoTable,
	}
}
