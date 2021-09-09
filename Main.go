//https://golang.org/doc/tutorial/database-access
//https://tutorialedge.net/golang/golang-orm-tutorial/

//how to dockerize golang app 
//https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

  
//https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html


//To do 
//1 Check all functions are working --- >today get/ set httpcodes 
//2 move app to docker  --> try today 
//3 Move app and sql both on same container 


package main
import (
	//"fmt"
   "log"
   "net/http"
	"fmt"
	"encoding/json"
	"database/sql"
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
   // tmpl.ExecuteTemplate(w, "Show", ListWeatherData)
    // defer db.Close()
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
func handleRequests() {
    //http.HandleFunc("/", home)
	http.HandleFunc("/homepage", homePage)
	//http.HandleFunc("/GetAllData", GetAllData)
    //http.HandleFunc("/ListOfCountries", GetAllCountries)
	http.HandleFunc("/WeatherJson", GetAllWeatherData)
	
	
    //http.HandleFunc("/ListOfLocation", GetAllLocations)
    // http.HandleFunc("/Weather", GetAllWeatherData)
    //http.HandleFunc("/WeatherByID", GetAllWeatherDataByID)


    http.HandleFunc("/Weather", Index)
	http.HandleFunc("/erase",DeleteAll)
    http.HandleFunc("/new", New)
    http.HandleFunc("/insert", Insert)
   // http.HandleFunc("/CountryList", CountryList)
    
    //curl http://localhost:8082/index

    log.Fatal(http.ListenAndServe(":8085", nil))
}


/*
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
*/

func main() {

    handleRequests()
}

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

//http://localhost:8083/Weather?date=08/21/2021  -works 
//http://localhost:8083/Weather?date=21/08/2021  throws null 








 
/*
curl -d weatherdate=06/12/2021 location='{
    lat: 37.5407,
    lon: 77.436,
    city: RICHMOND,
    state: VA
}' Temperature='1,2,3,4,5,6,7,8,9,10,11,12,13'   http://localhost:8085/insert

curl -d weatherdate=06/12/2021 location="{
    lat: 37.5407,
    lon: 77.436,
    city: RICHMOND,
    state: VA
}" Temperature="1,2,3,4,5,6,7,8,9,10,11,12,13"   http://localhost:8085/insert

curl -d weatherdate=06/12/2021 location="{
    lat: 37.5407,
    lon: 77.436,
    city: RICHMOND,
    state: VA
}" Temperature='1,2,3,4,5,6,7,8,9,10,11,12,13'   http://localhost:8085/insert

*/


/*

RUN mkdir /app
ADD . /app
WORKDIR /app 
COPY go.mod .
COPY go.sum .

#RUN go build -o main .
RUN go build . /app go 
#RUN go mod download
CMD ["./app"]
#ENTRYPOINT ["/Weather"]
*/