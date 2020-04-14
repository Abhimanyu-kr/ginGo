package controllers

import (
  "github.com/gin-gonic/gin"
//   "log"
	"net/http"
	// "fmt"
	// "bytes"

	// "github.com/gorilla/mux"
	// "ginGo/auth"
	"ginGo/api/models"
	"ginGo/api/responses"
	"ginGo/api/utils/formaterror"
)


// UserRegister : get post value and validate
func (server *Server) UserRegister(c *gin.Context) {
		var w http.ResponseWriter = c.Writer
	
	
		user := models.TUsers{}
		c.Bind(&user)
		

		user.Prepare()
		err := user.Validate("signup")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"STATUS": "ERROR", "message" : err.Error()})
			return
		}
		return
		userCreated, err := user.CreateNewUser(server.DB)
		if err != nil {
			formattedError := formaterror.FormatError(err.Error())
			c.JSON(http.StatusOK, gin.H{"STATUS": "ERROR", "message" : formattedError.Error()})
			return
		}
		
		responses.JSON(w, http.StatusCreated, userCreated)
		
		
		
  c.JSON(http.StatusOK, gin.H{"STATUS": "SUCCESS", "message" : "Successfull"})
  // return auth.CreateToken(user.ID)
}
