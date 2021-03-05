package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

// Init - initializes package
func Init(router *gin.RouterGroup, dbSession *mgo.Session) {
	repoService := newRepositoryService(dbSession)
	todoService := NewTodoService(repoService)
	MakeHTTPHandlers(router, todoService)
}
