package app

import (
	"encoding/json"
	// "encoding/xml"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"github.com/kaliayev-proton/banking-go-hex/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hellor world")
}
func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

		customers, err := ch.service.GetAllCustomers(status)

		if err != nil {
			writeResponse(w, err.Code, err.AsMessage())
		} else {
			writeResponse(w, http.StatusOK, customers)
		}

	// if (r.Header.Get("Content-Type") == "application/xml") {
	// 	w.Header().Add("Content-Type", "application/xml")
	
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {

	// 	w.Header().Add("Content-Type", "application/json")
		
	// 	json.NewEncoder(w).Encode(customers)
	// }

}

func (ch *CustomerHandlers) getOneCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}