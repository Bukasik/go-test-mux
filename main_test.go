package main

import (
	"os"
	"testing"
	"log"
)

var a App


func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("postgres"),
		os.Getenv("password"),
		os.Getenv("postgres"))

	ensureTableExists()
	code := m.Run()
	ClearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func ClearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`