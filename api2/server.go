package api

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"ginGo/api/controllers"
	// "ginGo/api/seed"
	// "ginGo/router"
)

var server = controllers.Server{}

// Run server
func Run() {
	
	var err error

	// if err := godotenv.Load("../.env"); err != nil {
	// 	log.Fatalf("Error getting env, not comming through %v", err)
	// }
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	fmt.Println("IN API -----------------------",os.Getenv("DB_DRIVER"))
	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	// Index view
	// server.initializeRoutes(server)
	// seed.Load(server.DB)

	// server := router.Router()

	// server.Run(":3000")

}
