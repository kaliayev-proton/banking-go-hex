package app

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"encoding/json"
	"strings"
	"github.com/gorilla/mux"
	"github.com/kaliayev-proton/banking-go-hex/service"
	"github.com/kaliayev-proton/banking-go-hex/domain"
)


func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getOneCustomer).Methods(http.MethodGet)

	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getOneCustomer).Methods(http.MethodGet)

	// Starting server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

func getOneCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POst request received")
}
func getTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
			loc, err := time.LoadLocation(tz)
			if err != nil {
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
			} else {
					response["current_time"] = time.Now().In(loc).String()
					w.Header().Add("Content-Type", "application/json")
					json.NewEncoder(w).Encode(response)
			}
	 } else {
			for _, tzdb := range timezones {
					loc, err := time.LoadLocation(tzdb)
					if err != nil {
							w.WriteHeader(http.StatusNotFound)
							w.Write([]byte(fmt.Sprintf("invalid timezone %s in input", tzdb)))
							return
					 }
					 now := time.Now().In(loc)
					 response[tzdb] = now.String()
			}
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
	 }
}