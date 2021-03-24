package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	//"strconv"

	"github.com/agk/specialist/c2/hw2/models"
	//"github.com/gorilla/mux"
)



func Grab(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("grab data and creating new ....")
	var Item models.Item
	
	
	err := json.NewDecoder(request.Body).Decode(&Item)
	if err != nil {
		msg := models.Message{Message: "provideed json file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}
	
	Item.Nroots := 0

	if Item.a == 0 {
		if Item.b != 0 {
			Item.Nroots = 1
			// x = fraction(-c,b).toString()
		} else if Item.c == 0 {
			Item.Nroots = 1
			// Все коэффициенты равны нулю; x - любое число.
		} else {
			Item.Nroots = 0
		}
	} else {
		d = math.Pow(Item.b, 2) - 4*Item.a*Item.c
		if d > 0 {
			Item.Nroots = 2
		} else if d == 0 {
			Item.Nroots = 1
		} else if d < 0 {
			Item.Nroots = 0
		}
	}
	

	DB = append(DB, Item)
	writer.WriteHeader(201)
	msg := models.Message{Message: "Grabbing ok"}
	json.NewEncoder(writer).Encode(msg)
	
}
func Solve(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	if len(models.DB) > 0{
		writer.WriteHeader(200)
		json.NewEncoder(writer).Encode(Item)
	}else{
		writer.WriteHeader(404)
		msg := models.Message{Message: "Error: data with that id not grab"}
		json.NewEncoder(writer).Encode(msg)
	} 
}