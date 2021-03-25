package handlers

import (
	// "fmt"
	"encoding/json"
	"log"
	"net/http"
	// "math"

	"github.com/agk/specialist/c2/sm/models"
	//"github.com/gorilla/mux"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func Grab(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("grab data and creating new ....")
	var Item models.Item
	
	err := json.NewDecoder(request.Body).Decode(&Item)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		err = json.NewEncoder(writer).Encode(msg)
		if err != nil {
			log.Fatal("can't encode message, something very wrong")
		}
		return
	}
	// fmt.Println(Item);
	models.AddItem(Item)	//models.DB = append(models.DB, Item)
	writer.WriteHeader(201)
	// msg := models.Message{Message: "Grabbing ok"}
	// json.NewEncoder(writer).Encode(msg)

	
	
	
	
}

func Solve(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	
	json.NewEncoder(writer).Encode(models.GetItem())
	writer.WriteHeader(200)
}