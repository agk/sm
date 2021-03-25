package main

import (
	"log"
	"net/http"
	"os"

	"github.com/agk/specialist/c2/sm/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	port	string
	grabResourcePrefix string = "/grab"
	solveResourcePrefix string = "/solve"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildResource(router, grabResourcePrefix)

	log.Println("Router initalizing successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
