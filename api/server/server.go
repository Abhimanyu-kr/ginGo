package server

import (
  "ginGo/router"
  "ginGo/api/controllers"
)


var server = controllers.Server{}

// Init to start
func Init() {
  server := staticrouter.StaticRouter()
  server.Run(":3001")
}