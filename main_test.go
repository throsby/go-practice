package main

import (
	"log"
	"os"
	"testing"

	_ "bytes"
	_ "encoding/json"
	_ "net/http"
	_ "net/http/httptest"
	_ "strconv"
	// _ "github.com/lib/pq"
	// _ "github.com/throsby/go-practice"
)

var a App

func TestMain(m *testing.M) {
	// a.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"))
	a.Initialize("throsbywells", "", "throsbywells")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL
	name TEXT NOT NULL
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)`
