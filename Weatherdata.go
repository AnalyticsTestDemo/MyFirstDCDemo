package main

import (
 	"fmt"
	"log"
	"net/http"
 	"text/template"	 
	"database/sql"
	"encoding/json"
	_ "github.com/denisenkom/go-mssqldb"

  )


var tmpl = template.Must(template.ParseGlob("*"))

//http://localhost:8085/Weather   
//http://localhost:8085/Weather?date=08/21/2021 
func Index(w http.ResponseWriter, r *http.Request) {
	var ListWeatherData []WeatherData
	var eachrow WeatherData
	
	filterDate := r.URL.Query().Get("date")
	 
	db := dbConn()
	if len(filterDate)<=0 {

    selDB, err := db.Query("select WeatherID , Location, WeatherDate, temperature from weatherdata where isactive=1 order by weatherid")
    if err != nil {
        panic(err.Error())
    }

    for selDB.Next() {  
        err = selDB.Scan(&eachrow.WeatherID, &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
        if err != nil {
			log.Fatal(err)
        }
		log.Println("In Weather for loop" ,eachrow.Temp)

        ListWeatherData = append(ListWeatherData, eachrow)
    }
		w.WriteHeader(200)
		tmpl.ExecuteTemplate(w, "Index",ListWeatherData)
		defer db.Close()

	}else
	{
		FilterRecordsByDate(filterDate, w)
	}
}

 func FilterRecordsByDate(wDate string,w http.ResponseWriter) {
		var ListWeatherData []WeatherData
		var eachrow WeatherData
		log.Println("In Filter Show date =" , wDate)

		db := dbConn()
		selDB, err := db.Query("select WeatherID , Location, WeatherDate, temperature from weatherdata where weatherdate =?", wDate )
		if err != nil {
			log.Fatal(err)
		}		
		for selDB.Next() {
			err = selDB.Scan(&eachrow.WeatherID, &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
			if err != nil {
			   log.Fatal(err)
			}
			log.Println("In Filter Show  loop" ,eachrow)
			ListWeatherData = append(ListWeatherData, eachrow)        
		}
		//w.WriteHeader(200)
		tmpl.ExecuteTemplate(w, "Show", ListWeatherData)
		defer db.Close()
	}


func New(w http.ResponseWriter, r *http.Request) {
    tmpl.ExecuteTemplate(w, "New", nil)
}

func DeleteAll(w http.ResponseWriter, r *http.Request){
	db := dbConn()
	log.Println("In DeleteAll . Now deleting all weather data ......")
	_, err := db.Query("Delete from weatherdata where isactive=1")	
    if err != nil {
        panic(err.Error())
    }
	Index(w,r)
}


func Insert(w http.ResponseWriter, r *http.Request) {
	var newLocation =r.FormValue("txtLocation")
	var newDate =r.FormValue("txtwDate")
	var cityname=getJsonValueCity(newLocation)
	db := dbConn()
	log.Println("In Insert new record ")

	isduplicate := IsRecordDuplicate(newDate ,	cityname	,db)

	if isduplicate{
		log.Println("Duplicate Record Exists for -" , cityname , newDate)
	}else{
    if r.Method == "POST" {
		
		var newData WeatherData

		var wID int64
		wDate := newDate
		loc := newLocation
		Temperatures := r.FormValue("txtTemp")
		log.Println("INSERT: Date: " + wDate + " | Loc: " + loc + " temps: " + Temperatures)

		newData.WeatherID = wID
		newData.WeatherDate = wDate
		newData.Location= loc 
		newData.Temp = Temperatures

		insForm, err := db.Prepare(" INSERT INTO WeatherData (weatherdate, location,temperature,isActive) VALUES (?,?,?,?)")
		if err != nil {
           // panic(err.Error())
		   log.Fatal(err)

		}
		insForm.Exec( newData.WeatherDate,newData.Location,newData.Temp, 1)
        log.Println("Record inserted: INSERT: Date: " + newData.WeatherDate + " | Loc: " + newData.Location)
		          
    }
}
    defer db.Close()
    http.Redirect(w, r, "/Weather", 301)
}






//Check if New Weather Data Record is Duplicate 
func  IsRecordDuplicate(newDate string, newLoc string,db *sql.DB) bool{	
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM weatherdata where weatherdate=? and location Like ?", newDate,"%" + newLoc +"%").Scan(&count)
	if err != nil {
        panic(err.Error())
    }
	 if count>0{
		return true
	}else 
	{
		return false
	}
}



func getJsonValueCity(jsonData string)string {
	var info map[string]interface{}
	json.Unmarshal([]byte(jsonData),&info)
	cityname:=fmt.Sprint(info["city"])
 	return  cityname
}
 




 