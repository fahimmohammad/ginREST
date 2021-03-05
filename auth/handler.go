package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	registerUserHandler(c *gin.Context)
	loginHandler(c *gin.Context)
}

type handlerStruct struct {
	authService *Service
}

func MakeHTTPHandlers(router *gin.RouterGroup, authService *Service) {
	h := &handlerStruct{
		authService: authService,
	}

	router.POST("auth/register", h.registerUserHandler)
	router.POST("auth/login", h.loginHandler)
}

type registerRequest struct {
	User User
}

type registerResponse struct {
	User User   `json:"user" `
	Err  string `json:"err" `
}

func (h *handlerStruct) registerUserHandler(c *gin.Context) {

	var req registerRequest
	if err := c.ShouldBindJSON(&req.User); err != nil {
		c.JSON(http.StatusInternalServerError, registerResponse{
			User: User{},
			Err:  "Error in data binding",
		})
		return
	}

	user, err := h.authService.RegisterUser(req.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, registerResponse{
			User: User{},
			Err:  err.Error(),
		})
		return
	}
	user.Password = ""
	res := registerResponse{
		User: user,
		Err:  "",
	}
	c.JSON(http.StatusOK, &res)
	return
}

type loginRequest struct {
	Login Login
}

type loginResponse struct {
	User         User   `json:"user" `
	Token        string `json:"token" `
	RefreshToken string `json:"refresh_token" `
	Err          string `json:"err" `
}

func (h *handlerStruct) loginHandler(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req.Login); err != nil {
		c.JSON(http.StatusInternalServerError, loginResponse{
			User:         User{},
			Token:        "",
			RefreshToken: "",
			Err:          "Error in data binding",
		})
		return
	}

	user, token, refreshToken, err := h.authService.LoginUser(req.Login)

	if err != nil {
		c.JSON(http.StatusInternalServerError, loginResponse{
			User:         User{},
			Token:        "",
			RefreshToken: "",
			Err:          err.Error(),
		})
		return
	}

	res := loginResponse{
		User:         user,
		Token:        token,
		RefreshToken: refreshToken,
		Err:          "",
	}
	c.JSON(http.StatusOK, &res)
	return

}

// Auth - example middleware
func Auth(auth *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(bearerToken, " ")
		token := splitToken[1]
		_, valid, _ := auth.VerifyWithParseToken(token)
		if !valid {
			c.JSON(401, gin.H{"Err": "invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
