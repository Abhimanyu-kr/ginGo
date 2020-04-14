package views

import (
  "github.com/gin-gonic/gin"
  "ginGo/api/controllers"
)

// Index call
func IndexView(route *gin.Engine)  {

  

  // controllers 
  controllers.Controllers(route)
  

}