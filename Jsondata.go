package main
import (
	//"fmt"
   "log"
   "net/http"
	"fmt"
	"encoding/json"
	"database/sql"
) 

//Displays all weather data in html table 
func GetAllData(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: GetAllData")
    json.NewEncoder(w).Encode(AllCountries)
  }


func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: homePage")
	//tmpl.ExecuteTemplate(w, "HomePage","")
}

 //curl http://localhost:8085/WeatherJson   
func GetAllWeatherData(w http.ResponseWriter, r *http.Request){
    var condb *sql.DB
	filterDate := r.URL.Query().Get("date")
   
    condb, errdb := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
    if errdb != nil {
        fmt.Println(" Error open db:", errdb.Error())
      }

	if len(filterDate)<=0 {
		ListWeatherData:=GetWeatherData(condb)
		json.NewEncoder(w).Encode(ListWeatherData)
		
	}else{
		ListWeatherData:=FilterRecords(condb,filterDate,w)	
		json.NewEncoder(w).Encode(ListWeatherData)
	}	

	//http: superfluous response.WriteHeader call from main.GetAllWeatherData (main.go:94)	//	
    //No need to set the status as ok as it is the default
	//w.WriteHeader(200)
	defer condb.Close()
  }

  func FilterRecords(db *sql.DB,wDate string,w http.ResponseWriter) []WeatherData {
	var AllWeatherData []WeatherData
	var eachrow WeatherData
	log.Println("In Filter Show")  
    
    selDB, err := db.Query("select WeatherID , Location, WeatherDate, temperature from weatherdata where weatherdate =?", wDate)
    if err != nil {
		log.Fatal(err)
    }
    log.Println("In Filter Show before loop")
    for selDB.Next() {
		err = selDB.Scan(&eachrow.WeatherID, &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
        if err != nil {
		   log.Fatal(err)
        }
        AllWeatherData = append(AllWeatherData, eachrow)        
    }
	return AllWeatherData
}

  
  func GetWeatherData(db *sql.DB) []WeatherData{
	var AllWeatherData []WeatherData
		var eachrow WeatherData
		fmt.Println("Endpoint Hit: GetWeatherData")
	
		selDB, err := db.Query("select WeatherID ,  WeatherDate,Location, temperature from weatherdata where isactive=1")
		if err != nil {
			panic(err.Error())
		}
		for selDB.Next() {
		   
			err = selDB.Scan(&eachrow.WeatherID, &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
			if err != nil {
				//panic(err.Error())
				log.Fatal(err)
			}	 
			AllWeatherData = append(AllWeatherData, eachrow)
		}
 	return AllWeatherData
 	}