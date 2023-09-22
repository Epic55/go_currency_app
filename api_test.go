// main_test.go

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

type currency struct {
	ID     int
	a_date string
	title  string
	code   string
	value  string
}

func (a *App) Initialize(user, password, host, port, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/currency/{date1}", a.getCurrency).Methods("GET")
}

func (a *App) getCurrency(w http.ResponseWriter, r *http.Request) {
	count, _ := strconv.Atoi(r.FormValue("count"))
	start, _ := strconv.Atoi(r.FormValue("start"))

	if count > 10 || count < 1 {
		count = 10
	}
	if start < 0 {
		start = 0
	}

	products, err := getCurrency(a.DB, start, count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, products)
}

func getCurrency(db *sql.DB, start, count int) ([]currency, error) {
	statement := fmt.Sprintf("SELECT id, a_date, title, code, value FROM r_currencies")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	currencies := []currency{}

	for rows.Next() {
		var u currency
		if err := rows.Scan(&u.ID, &u.a_date, &u.title, &u.code, &u.value); err != nil {
			return nil, err
		}
		currencies = append(currencies, u)
	}

	return currencies, nil
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("postgres", "1", "localhost", "5432", "db1")

	code := m.Run()

	os.Exit(code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func TestGetCurrency(t *testing.T) {
	addCurrency(1)

	req, _ := http.NewRequest("GET", "/currency/29.08.2023", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
	//clearTable()
}
