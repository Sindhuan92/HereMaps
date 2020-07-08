package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
	FacilityName string `OffJson:"name"`
	Id int  `OffJson:"id"`
	PaymentName  string `OffJson:"name"`
	PaymentId int `OffJson:"id"`
	City string `OffJson:"city"`
	Country string `OffJson:"country"`


    
}

func main() {
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%s %d %s %d %s %s", user.FacilityName, user.Id, user.PaymentName,user.PaymentId,user.City,user.Country)
  })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        Berlin := User{

			FacilityName:"Parking Facility",
			Id: 1,
			PaymentName:"Cash - EUR",
			PaymentId:1000,
			City:"Berlin",
			Country:"DEU",
		
		
        }

        json.NewEncoder(w).Encode(Berlin)
    })

    http.ListenAndServe(":8080", nil)
}


