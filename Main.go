package main
import (
	//"fmt"
   "log"
   "net/http"
	//"fmt"
	//"encoding/json"
	//"database/sql"
) 

type CountryMaster struct {
    CountryId    int
    CountryName  string
  }
var AllCountries []CountryMaster

type  Location struct {
    Cityname string
    State string
    Latitude float64
    Longitude float64
}
var AllCities []Location

type WeatherData struct {
    WeatherID int64
    Location string 
    WeatherDate string 
    Temp string
}
 
 
	
  
func handleRequests() {
    http.HandleFunc("/homepage", homePage)
	http.HandleFunc("/WeatherJson", GetAllWeatherData)
    
    http.HandleFunc("/Weather", Index)
	http.HandleFunc("/erase",DeleteAll)
    http.HandleFunc("/new", New)
    http.HandleFunc("/insert", Insert)

    log.Fatal(http.ListenAndServe(":8085", nil))

    //http.HandleFunc("/", home)
	//http.HandleFunc("/GetAllData", GetAllData)
    //http.HandleFunc("/ListOfCountries", GetAllCountries)	
    //http.HandleFunc("/ListOfLocation", GetAllLocations)
    // http.HandleFunc("/Weather", GetAllWeatherData)
    //http.HandleFunc("/WeatherByID", GetAllWeatherDataByID)
    // http.HandleFunc("/CountryList", CountryList)
    //curl http://localhost:8082/index

}



func main() {
    handleRequests()
}


//http://localhost:8083/Weather?date=08/21/2021  -works 
//http://localhost:8083/Weather?date=21/08/2021  throws null 







