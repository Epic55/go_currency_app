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
	"testing"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, host, port, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", user, password, host, port, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Error1 - ", err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/currency/{date1}", a.getUsers).Methods("GET")
}

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Initialize("postgres", "1", "localhost", "5432", "db1")

	ensureTableExists()

	code := m.Run()

	//clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM r_currencies WHERE a_date=01.08.2023")
	//a.DB.Exec("ALTER TABLE users AUTO_INCREMENT = 1")
}

const tableCreationQuery = `
CREATE TABLE IF NOT EXISTS users
(
    id serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
)`

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/currency/01.08.2023", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}
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

func TestGetNonExistentCurrency(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/currency/01.08.2023", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "User not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'User not found'. Got '%s'", m["error"])
	}
}

func addUsers(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		statement := fmt.Sprintf("INSERT INTO r_currencies(a_date, title, code, value) VALUES('%s', %s, '%s', '%s')", "28.08.2023", "SSS", "SSS", "1")
		a.DB.Exec(statement)
	}
}

func TestGetCurrency(t *testing.T) {
	//clearTable()
	addUsers(1)

	req, _ := http.NewRequest("GET", "/currency/28.08.2023", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}
