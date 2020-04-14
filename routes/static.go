package routes

import "github.com/gin-gonic/gin"

func PublicImages(route *gin.Engine)  {
  route.Static("/public/images", "./public/images")
}

func PublicCss(route *gin.Engine)  {
  route.Static("/public/css", "./public/css")
}

func PublicPlugin(route *gin.Engine)  {
  route.Static("/public/plugin", "./public/plugin")
}

func PublicJs(route *gin.Engine)  {
  route.Static("/public/js", "./public/js")
}
