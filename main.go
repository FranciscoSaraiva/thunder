package main

import (
	"log"
	"net/http"
	"os"

	//"github.com/julienschmidt/httprouter"

	"thunder/home"
	"thunder/input"
	"thunder/config"
)

func main() {
	router := config.NewServerAPI()


	// Show index page
	router.GET("/", home.Index)

	// InBounce
	router.POST("/in_bounce", input.Create)
	router.GET("/in_bounce/:id", input.Show)
	//router.PUT("/in_bounce/:id", input.Update) Not used or needed

	// Static fields
	router.NotFound = http.FileServer(http.Dir("public"))

	// Start server
	log.Println("[ServerAPI] ServerAPI is running in http://localhost:8080/")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("Error: %v", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}
