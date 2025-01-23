package auth

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"errors"
	"golang.org/x/crypto/bcrypt"
)
func Authenticate(email, password string) (*User, error) {
	user, err := FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func Register(creds User) (*User, error) {

	user, err := FindByEmail(creds.Email)
	if user != nil {
		return user, errors.New("user already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user = &User{
		FirstName: creds.FirstName,
		LastName: creds.LastName,
		Username: creds.Username,
		Email: creds.Email,
		Password: string(hashedPassword),
		PhoneNumber: creds.PhoneNumber,
	}

	return CreateUser(user)

}



func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	err := session.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}
