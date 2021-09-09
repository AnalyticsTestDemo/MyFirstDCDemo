package main

import (
	"database/sql"
	"fmt"
	//"log"
	"net/http"
	  "encoding/json"  
  )



func GetAllLocations(w http.ResponseWriter, r *http.Request){
    var condb *sql.DB
    condb, errdb := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
    if errdb != nil {
        fmt.Println(" Error open db:", errdb.Error())
      }
	  GetLocationFromDB(condb)
    fmt.Println("Endpoint Hit: GetAllLocations")
    json.NewEncoder(w).Encode(AllCities)
  }

func GetLocationFromDB(condb *sql.DB){
    var eachrow Location
    sqlstmt:="select c.cityname, s.statename , c.Latitude, c.longitude from citymaster c ,StateMaster s where s.stateid = c.stateid"
    rows, err := condb.Query(sqlstmt)
    if err != nil {
        panic(err)
     }
     defer rows.Close()
     for rows.Next(){
         rows.Scan(&eachrow.Cityname, &eachrow.State, &eachrow.Latitude, &eachrow.Longitude)
         AllCities = append(AllCities,eachrow)
     }
}





/*
func LocationList(w http.ResponseWriter, r *http.Request) {
	var ListofLocations []Location
	var eachrow Location

    db := dbConn()

    selDB, err := db.Query(" SELECT LocationID, Locationname, Latitude, Longitude FROM CountryMaster")
    if err != nil {
        panic(err.Error())
    }

    for selDB.Next() {
       
        err = selDB.Scan(&eachrow.LocationID, &eachrow.Locationname,&eachrow.Latitude, &eachrow.Longitude )
        if err != nil {
            //panic(err.Error())
			log.Fatal(err)

        }
         
        ListofCountries = append(ListofCountries, eachrow)
    }
    tmpl.ExecuteTemplate(w, "Country",ListofCountries)
    defer db.Close()
}
*/






