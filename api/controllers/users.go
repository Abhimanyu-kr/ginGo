package controllers

import (
  "github.com/gin-gonic/gin"
  "log"
	"net/http"
	// "fmt"
	// "bytes"
	"strings"
	"ginGo/api/auth"
	"ginGo/api/models"
	"ginGo/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)


// UserRegister : get post value and validate
func (server *Server) UserRegister(c *gin.Context) {
		// var w http.ResponseWriter = c.Writer

		user := models.TUsers{}
		c.Bind(&user)
		

		user.Prepare()
		err := user.Validate("signup")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"STATUS": "ERROR", "message" : err.Error()})
			return
		}
		// log.Println(user)
		userCreated, err := user.CreateNewUser(server.DB)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			c.JSON(http.StatusOK, gin.H{"STATUS": "ERROR", "message" : formattedError.Error()})
			return
		}
		
		
		
  c.JSON(http.StatusOK, gin.H{"STATUS": "SUCCESS", "message" : "Successfull", "ROW": userCreated})
  // return auth.CreateToken(user.ID)
}

// UserLogin : login and token creation
func (server *Server) UserLogin(c *gin.Context) {
	// var w http.ResponseWriter = c.Writer
	
	user := models.TUsers{}
	c.Bind(&user)
	

	user.Prepare()
	err := user.Validate("login")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"STATUS": "ERROR", "message" : err.Error()})
		return
	}
	log.Println(user.Email)
	log.Println(user.Username)
	// userCreated, err := user.CreateNewUser(server.DB)
	token, err := server.SignIn(user.Email, user.Username, user.Password)
	
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		c.JSON(http.StatusOK, gin.H{"STATUS": "ERROR", "message" : formattedError.Error()})
		return
	}	
	
	
c.JSON(http.StatusOK, gin.H{"STATUS": "SUCCESS", "message" : "Login Successfull", "token": token})
// return auth.CreateToken(user.ID)
}


func (server *Server) SignIn(email, username, password string) (string, error) {

	user := models.TUsers{}
	log.Println("====================")
	log.Println(user)
	// err = server.DB.Table("tUsers").Create(&u).Error
	err := server.DB.Table("tUsers").Where("email = ? OR username= ?", email, username).Take(&user).Error
	if err != nil {
		return "", err
	}
	
	log.Println(user.Password)
	log.Println(user.Username)
	log.Println(password)
	hashedPassword, err := models.Hash(password)
	log.Println(string(hashedPassword))
	// err = models.VerifyPassword("$2a$10$rqHJJTHsxMbtX/5ZjG1mFuWyYbUDW1PLbfwQRN0uChwes38c/0m3e", "")
	err = models.VerifyPassword(strings.TrimSpace(user.Password), strings.TrimSpace(password))
	log.Println(err)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.User_id)
}
