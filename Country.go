package main

import (
	//"database/sql"
	//"fmt"
	"log"
	"net/http"
	 // "encoding/json"
  // _ "github.com/denisenkom/go-mssqldb"
  
  )






func CountryList(w http.ResponseWriter, r *http.Request) {
	var ListofCountries []CountryMaster
	var eachrow CountryMaster

    db := dbConn()

    selDB, err := db.Query(" SELECT countryid, countryname FROM CountryMaster")
    if err != nil {
        panic(err.Error())
    }

    for selDB.Next() {
       
        err = selDB.Scan(&eachrow.CountryId, &eachrow.CountryName )
        if err != nil {
            //panic(err.Error())
			log.Fatal(err)

        }
         
        ListofCountries = append(ListofCountries, eachrow)
    }
    tmpl.ExecuteTemplate(w, "Country",ListofCountries)
    defer db.Close()
}










/*
  func GetAllCities(w http.ResponseWriter, r *http.Request){
    var condb *sql.DB
    condb, errdb := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
    if errdb != nil {
        fmt.Println(" Error open db:", errdb.Error())
      }
    GetAllCountriesJson(condb)
    fmt.Println("Endpoint Hit: GetAllData")
    json.NewEncoder(w).Encode(AllCountries)
  }
  */