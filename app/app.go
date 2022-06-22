package app

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"os"
	"encoding/json"
	"strings"
	"github.com/gorilla/mux"
	"github.com/kaliayev-proton/banking-go-hex/service"
	"github.com/kaliayev-proton/banking-go-hex/domain"
	"github.com/jmoiron/sqlx"
)

func sanityCheck() {
	if os.Getenv("BANKING_SERVER_ADDRESS") == "" ||
		os.Getenv("BANKING_SERVER_PORT") == "" {
			log.Fatal("Environment variable not defined...")
		}
}

func Start() {


	sanityCheck()
	// mux := http.NewServeMux()
	router := mux.NewRouter()

	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)

	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// Define routes
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", ch.getOneCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{customer_id:[0-9]+}/account", ah.newAccount).Methods(http.MethodPost)

	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getOneCustomer).Methods(http.MethodGet)

	// Get environment variable from the OS
	address := os.Getenv("BANKING_SERVER_ADDRESS")
	port := os.Getenv("BANKING_SERVER_PORT")

	// Starting server
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
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

func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("BANKING_DB_USER")
	dbPassword := os.Getenv("BANKING_DB_PASSWORD")
	dbAddress := os.Getenv("BANKING_DB_ADDRESS")
	dbPort := os.Getenv("BANKING_DB_PORT")
	dbName := os.Getenv("BANKING_DB_NAME")
	
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}