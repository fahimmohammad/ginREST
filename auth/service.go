
package auth

import (
	"errors"
	"time"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)
	
const SECRET_KEY = "hhhgssscejkSSSde61511165ajnk4*7378422idnX"
	
type ServiceInterface interface {
	RegisterUser(user User) (User, error)
	FindByUserEmail(email string) (User, error)
	VerifyWithParseToken(token string) (User, bool, error)
	LoginUser(login Login) (User, string, string, error)
}
	
type Service struct {
	repoService *repoStruct
}
	
func (authService *Service) RegisterUser(user User) (User, error) {
	
	_, err := authService.FindByUserEmail(user.Email)
	
	if err == nil {
		return User{}, errors.New("User Exists")
	}
	user.ID = generateUUID()
	hashedpass, _ := hashPassword(user.Password)
	user.Password = hashedpass

	usr, err := authService.repoService.CreateUser(user)
	if err != nil {
		return User{}, errors.New("Cannot Create User")
	}
	return usr, nil
}
	
func (authService *Service) FindByUserEmail(email string) (User, error) {
	userf, err := authService.repoService.FindByEmail(email)
	if err != nil {
		return User{}, err
	}
	
	return userf, nil
}
	
func (auth *Service) VerifyWithParseToken(token string) (User, bool, error) {
	var usert User
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected signing method")
		}
	
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return User{}, false, err
		}
		return User{}, false, err
	}
	if !tkn.Valid {
		return User{}, false, errors.New("Token invalid")
	}

	mapstructure.Decode(tkn.Claims, &usert)
	return usert, true, nil
}
	
func (authService *Service) LoginUser(login Login) (User, string, string, error) {
	userl, err := authService.FindByUserEmail(login.Email)
	if err != nil {
		return User{}, "", "", err
	}

	match := checkPasswordHash(login.Password, userl.Password)
	if match == false {
		return User{}, "", "", errors.New("Password Not Matched")
	}
	token, err := generateToken(userl)
	if err != nil {
		return User{}, "", "", err
	}
	refreshToken, err := generateRefreshToken(userl)
	if err != nil {
		return User{}, "", "", err
	}
	userl.Password = ""
	return userl, token, refreshToken, nil
}
	
func generateUUID() string {
	return uuid.New().String()
}
	
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 5)
	return string(bytes), err
}
	
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
	
func generateToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["user_name"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * 12).Unix()
	t, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}
	
func generateRefreshToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["email"] = user.Email
	claims["user_name"] = user.UserName
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	t, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return t, nil
}
	
func NewAuthService(repo *repoStruct) *Service {
	return &Service{
		repoService: repo,
	}
}	
