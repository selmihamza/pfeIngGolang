package main

import (
	"driver/Helpers"
	routers "driver/Routers"
	_ "driver/docs"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/subosito/gotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Init func
func Init() {
	gotenv.Load()
}

// @title DriversSpace API
// @version 1.0
// @description This is a sample service for managing drivers space
// @termsOfService http://swagger.io/terms/
// @contact.name Selmi hamza
// @contact.email hamza@craftfoundry.tech
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host api-company.herokuapp.com
// @BasePath /
func main() {

	Init()
	Helpers.DbConnect()

	router := routers.InitializeRouter()
	//Set the port
	port := GetPort()
	//swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	//server
	err := http.ListenAndServe(port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"https://yuso-front.herokuapp.com", "http://localhost:3000"}))(router))
	if err != nil {
		panic(err)
	}
}

// GetPort function
func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
		log.Println("[-] No PORT environment variable detected. Setting to ", port)
	}
	return ":" + port
}
