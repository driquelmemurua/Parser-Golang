package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "granvolumen"
	password = "1"
	dbname   = "granvolumen"
)

func main() {
	connection_time_start := time.Now()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	connection_time_end := time.Since(connection_time_start)
	fmt.Println("Successfully connected!")
	query_time_start := time.Now()
	//rows, err := db.Query(`SELECT edad FROM personas1`)
	rows, err := db.Query(`SELECT edad FROM personas2`)
	defer rows.Close()
	var edades [100]uintptr
	for rows.Next() {
		var edad uint16
		err = rows.Scan(&edad)
		edades[edad]++
	}
	query_time_end := time.Since(query_time_start)
	for index, element := range edades {
		fmt.Printf("Edad:%d Personas:%d\n", index, element)
	}
	fmt.Printf("Tiempo de conexion:%s Tiempo de consulta:%s\n", connection_time_end, query_time_end)
}
