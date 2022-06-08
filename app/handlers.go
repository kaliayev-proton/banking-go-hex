package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"github.com/kaliayev-proton/banking-go-hex/service"
)

type Customer struct {
	Name string `json:"full_name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hellor world")
}
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer {
		// 	{Name: "victor", City: "Madrid", Zipcode: "28026"},
		// 	{Name: "victor", City: "Madrid", Zipcode: "28026"},
		// }

		customers, _ := ch.service.GetAllCustomers()

	if (r.Header.Get("Content-Type") == "application/xml") {
		w.Header().Add("Content-Type", "application/xml")
	
		xml.NewEncoder(w).Encode(customers)
	} else {

		w.Header().Add("Content-Type", "application/json")
		
		json.NewEncoder(w).Encode(customers)
	}

}

func (ch *CustomerHandlers) getOneCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	w.Header().Add("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(err.Code)
		// fmt.Fprintf(w, err.Message)
		json.NewEncoder(w).Encode(err.AsMessage())
	} else {
		json.NewEncoder(w).Encode(customer)
	}
}


