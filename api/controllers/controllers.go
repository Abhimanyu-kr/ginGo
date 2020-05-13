package controllers

import (
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	// "ginGo/middlewares"
	"html/template"
	"net/http"
	"time"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	
)
var server = Server{}

// Controllers : All route controllers for GET and POST
func Controllers(route *gin.Engine)  {

	// page controller
	app := ginview.NewMiddleware(goview.Config{
		Root:      "templates/",
		Extension: ".html",
		Master:    "layouts/master",
		Partials:  []string{"partials/link",
		   "partials/meta",
		   "partials/title",
		   "partials/scripts_head",
		   "partials/scripts_foot"},
		Funcs: template.FuncMap{
		   "copy": func() string {
			  return time.Now().Format("2006")
		   },
		},
		DisableCache: true,
	 })


	//  db connectivity
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))


	// Front end group
	appGroup := route.Group("/", app)
	{
		{
			appGroup.GET("/", func(ctx *gin.Context) {
				ginview.HTML(ctx, http.StatusOK, "index", gin.H{
					"title": "GinGo | HOME",
				})
			})

			// appGroup.GET("/signup", func(ctx *gin.Context) {
			// 	ginview.HTML(ctx, http.StatusOK, "users/signup", gin.H{
			// 	"title": "GinGo | SIGN UP",
			// 	})
			// })

			// appGroup.GET("/login", func(ctx *gin.Context) {
			// 	ginview.HTML(ctx, http.StatusOK, "users/login", gin.H{
			// 	"title": "GinGo | LOGIN",
			// 	})
			// })
			

		}
	}


	// API group
	apiGroup := route.Group("/api/", app)
	{
		{
			apiGroup.POST("/go_signup", server.UserRegister)

			apiGroup.POST("/go_login", server.UserLogin)
		}
	}

	
  
}