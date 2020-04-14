package staticrouter

import (
  "github.com/gin-gonic/gin"
  "ginGo/routes"
  "ginGo/views"
)

func StaticRouter() *gin.Engine  {

  router := gin.New()
  router.Use(gin.Logger())
  router.Use(gin.Recovery())



  // System routes
  routes.NotFound(router)
  routes.NoMethods(router)

  // Static Routes
  routes.PublicPlugin(router)
  routes.PublicCss(router)
  routes.PublicImages(router)
  routes.PublicJs(router)

  // Index view
  views.IndexView(router)

  return router
}