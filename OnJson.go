package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
	ParkingId int `OnJson:"id"`
	Capacity int `OnJson:"capacity"`
	 Street string `OnJson:"street"`
	 City string `OffJson:"city"`
	AvailableSpots int `OffJson:"availableSpots"`
	Trend string `OffJson:"trend"`


    
}

func main() {
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%d %d %s %s %d %s", user.ParkingId, user.Capacity, user.Street,user.City,user.AvailableSpots,user.Trend)
  })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        Berlin := User{
			ParkingId:192191,
			Capacity:8,
			Street:"Carrer de Galileu, 314",
			City:"Barcelona",
			AvailableSpots:3,
			Trend:"STATIC",

			
        }

        json.NewEncoder(w).Encode(Berlin)
    })

    http.ListenAndServe(":8080", nil)
}


