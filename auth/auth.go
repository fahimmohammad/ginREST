package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
)

// Init - initializes package auth
func Init(router *gin.RouterGroup, dbSession *mgo.Session) *Service {
	repoService := NewRepository(dbSession)
	authService := NewAuthService(repoService)
	MakeHTTPHandlers(router, authService)
	return authService
}
