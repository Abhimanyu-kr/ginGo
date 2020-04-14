package controllers

import (
	"net/http"

	"ginGo/api/responses"
	// "github.com/foolin/goview"
	// "github.com/foolin/goview/supports/ginview"
	// "github.com/gin-gonic/gin"
	// "html/template"
	// "time"
)

// Home test
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
	
}
