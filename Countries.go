package main
import (
   "net/http"
	"fmt"
	"encoding/json"
	"database/sql"
)
	func GetAllCountries(w http.ResponseWriter, r *http.Request){
		var condb *sql.DB
		condb, errdb := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
		if errdb != nil {
			fmt.Println(" Error open db:", errdb.Error())
		  }
		  GetCountriesFromDB(condb)
		fmt.Println("Endpoint Hit: GetAllCountries")
		w.WriteHeader(200 )
		json.NewEncoder(w).Encode(AllCountries)
	}
	
    func GetCountriesFromDB(condb *sql.DB){
		var eachrow CountryMaster
		sqlstmt:="SELECT countryid, countryname FROM CountryMaster "
		rows, err := condb.Query(sqlstmt)
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next(){
			rows.Scan(&eachrow.CountryId, &eachrow.CountryName)
			AllCountries = append(AllCountries,eachrow)
		}
	}