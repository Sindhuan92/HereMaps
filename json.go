package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type User struct {
	NoOfConnectors int `json:"totalNumberOfConnectors"`
	SupplierName string `json:"supplierName"`
	ConnectorName string `json:"name"`
        Id int  `json:"id"`
    ChargeCapacity  string `json:"chargeCapacity"`
    
}

func main() {
    http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
        var user User
        json.NewDecoder(r.Body).Decode(&user)

        fmt.Fprintf(w, "%d %s %s %d %s", user.NoOfConnectors, user.SupplierName, user.ConnectorName,user.Id,user.ChargeCapacity)
    })

    http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
        Berlin := User{
            NoOfConnectors: 2,
            SupplierName:  "RWE eRoaming",
			ConnectorName: "IEC 62196-2 type 2 (Mennekes)",
			Id: 31,
			ChargeCapacity:"380-480VAC, 3-phase at max. 32 A",
		
        }

        json.NewEncoder(w).Encode(Berlin)
    })

    http.ListenAndServe(":8080", nil)
}