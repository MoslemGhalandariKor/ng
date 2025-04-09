package auth

import (
	"fmt"
	"net/http"
	"nextgen/internals/helpers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	creds.Email = c.PostForm("email")
	creds.Password = c.PostForm("password")

	user, err := Authenticate(creds.Email, creds.Password)
	if err != nil {
		helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "error", Message: "Invalid Credentials"})
		// CHANGE STATUS CODE
		c.Redirect(http.StatusFound, "/login")
		return
	}
	SetSession(c, "Email", user.Email)
	SetSession(c, "Username", user.Username)
	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Successfully logged in"})
	c.Redirect(http.StatusFound, "/dashboard")
}

func RegisterHandler(c *gin.Context) {
	var creds User
	creds.FirstName = c.PostForm("first_name")
	creds.LastName = c.PostForm("last_name")
	creds.Username = c.PostForm("username")
	creds.Email = c.PostForm("email")
	creds.Password = c.PostForm("password")
	creds.PhoneNumber = c.PostForm("phone_number")
	user, err := Register(creds)
	if err != nil {
		if user != nil {
			helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "error", Message: "User already exists please login"})
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	SetSession(c, "Email", user.Email)
	SetSession(c, "Username", user.Username)
	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Successfully logged in"})
	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Please Fil The following form to complete your profile"})
	c.Redirect(http.StatusFound, "/dashboard/personal-profile")
}

func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		fmt.Println(err)
	}
	c.Redirect(http.StatusFound, "/home")
}

func AuthMiddleware(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get("Email")
	fmt.Println(user)
	if user == nil {
		helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "error", Message: "You need to sign in first!!"})
		c.Redirect(302, "/login")
		c.Abort()
		return
	}
	c.Next()
}

// func RegisterHandler(c *gin.Context) {

//     var creds User

// 	creds.Username = c.PostForm("username")
// 	creds.FirstName = c.PostForm("first_name")
// 	creds.LastName = c.PostForm("last_name")
// 	creds.Email = c.PostForm("email")
// 	creds.PhoneNumber = c.PostForm("phone_number")
// 	creds.Password = c.PostForm("password")

//     SetSession(c, "Email", creds.Email)
// 	SetSession(c, "Username", creds.Username)
// 	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Successfully logged in"})
// 	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Please Fil The following form to complete your profile"})
// 	c.Redirect(http.StatusFound, "/dashboard")
// }

// func LoginHandler(c *gin.Context){
// 	var creds struct {
// 		Email       string `json:"email"`
// 		Password    string `json:"password"`
// 	}

// 	creds.Email = c.PostForm("email")
// 	creds.Password = c.PostForm("password")

//     user, err := GetUserByEmail(context.Background(),creds.Email)

// 	if err != nil{
// 		helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "error", Message: "Email or Password is wrong please try again"})
// 		c.Redirect(http.StatusFound, "/login")
// 		c.Abort()
// 		return
// 	}
//     if user.Password != creds.Password{
// 		helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "error", Message: "Email or Password is wrong please try again"})
// 		c.Redirect(http.StatusFound, "/login")
// 		c.Abort()
// 		return
// 	}
// 	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Successfully logged in"})
// 	helpers.SetFlash(c, "flash", helpers.FlashMessage{Type: "success", Message: "Please Fil The following form to complete your profile"})
// 	SetSession(c, "Email", creds.Email)
// 	SetSession(c, "Username", user.Username)
// 	c.Redirect(http.StatusFound, "/dashboard")
// }

// func AuthMiddleware(c *gin.Context) {
// 	session := sessions.Default(c)
// 	user := session.Get("Email")
// 	fmt.Println(user)
// 	if user == nil {
// 		c.Redirect(302, "/login")
// 		c.Abort()
// 		return
// 	}
// 	c.Next()
// }
