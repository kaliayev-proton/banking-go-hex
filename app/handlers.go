package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}



func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hellor world")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{Name: "victor", City: "Madrid", Zipcode: "28026"},
		{Name: "victor", City: "Madrid", Zipcode: "28026"},
	}

	if (r.Header.Get("Content-Type") == "application/xml") {
		w.Header().Add("Content-Type", "application/xml")
	
		xml.NewEncoder(w).Encode(customers)
} else {

	w.Header().Add("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(customers)
}

}
