package main
import (
	"database/sql"
	"fmt"
	"log"
  )


  func dbConn() (db *sql.DB) {
    var condb *sql.DB
    fmt.Println("Endpoint Hit: homePage connection")
    condb, err := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
    if err != nil {
        log.Fatal(err)
		fmt.Println("Endpoint Hit: Error in db con")
    }
    return condb    
}