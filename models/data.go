package models
import (
	"log"
	"math"
)
var DB []Item

type Item struct {
	A      int `json:"A"`
	B      int `json:"B"`
	C      int `json:"C"`
	Nroots int `json:"Nroots"`
}

func AddItem(i Item) {
	//i.Nroots := 0
	if i.A == 0 {
		if i.B != 0 {
			i.Nroots = 1
		} else if i.C == 0 {
			i.Nroots = 1
		} else {
			i.Nroots = 0
		}
	} else {
		d := int(math.Pow(float64(i.B), 2)) - 4*i.A*i.C
		if d > 0 {
			i.Nroots = 2
		} else if d == 0 {
			i.Nroots = 1
		} else if d < 0 {
			i.Nroots = 0
		}
	}
	DB = append(DB, i)
}

func GetItem() Item {
	if len(DB) == 0 {
		log.Fatal("Error: data with that id not grab")
	} 
	return DB[len(DB)-1]
}