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
            //panic(err.Error())
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

	// Print the output from info map.
	//fmt.Println("Printin location details")
	//fmt.Println(info["lat"])
	//fmt.Println(info["lon"])
	cityname:=fmt.Sprint(info["city"])
	//fmt.Println(info["state"])
 	return  cityname
}
 








/*
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
*/

	
/*
func Show(w http.ResponseWriter, r *http.Request) {
	var ListWeatherData []WeatherData
	var eachrow WeatherData

    db := dbConn()
    //Id := r.URL.Query().Get("id")
    //fmt.Println("Endpoint Hit: Show")
	Id:=1
    
    selDB, err := db.Query("select WeatherID , Location, WeatherDate, temperature from weatherdata where weatherid =?", Id)
    if err != nil {
		log.Fatal(err)
    }
    
    for selDB.Next() {
		err = selDB.Scan(&eachrow.WeatherID, &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
        if err != nil {
		   log.Fatal(err)
        }
        ListWeatherData = append(ListWeatherData, eachrow)        
    }
    tmpl.ExecuteTemplate(w, "Show", ListWeatherData)
    defer db.Close()
}


func getJsonValueCity(jsonData string)string {
	var info map[string]interface{}
	json.Unmarshal([]byte(jsonData),&info)

	// Print the output from info map.
	//fmt.Println("Printin location details")
	//fmt.Println(info["lat"])
	//fmt.Println(info["lon"])
	cityname:=fmt.Sprint(info["city"])
	//fmt.Println(info["state"])
 	return  cityname
}
*/

//insForm.Exec( "11/06/2021","Ireland","34.5", 1)

/*
WORKING CODE 
		insForm, err := db.Prepare(" INSERT INTO countryMaster (countryname,isactive) VALUES (?,?)")
		if err != nil {
            panic(err.Error())
		}
		insForm.Exec( "Ireland", 1)
		 
*/
		

//insert into weatherdata (weatherdate, location,temperature,isActive)
//--values ( convert( datetime,'21/06/2021',103), 'Seattle', '67',1)
/*
		 	insForm, err := db.Prepare(" INSERT INTO Weatherdata (weatherdate, location,temperature,isActive) VALUES (?, ?, ?,?) ")
		if err != nil {
            panic(err.Error())
        }
        //insForm.Exec("Convert(datetime,'" + newData.WeatherDate +"',103)", newData.Location, newData.Temp,1)
		//insForm.Exec("TO_DATE('17/12/2015', 'DD/MM/YYYY')", newData.Location, newData.Temp,1)
		insForm.Exec("TO_DATE('17/12/2015', 'DD/MM/YYYY')", "DC","101",1)
		
        log.Println("INSERT: Date: " + newData.WeatherDate + " | Loc: " + newData.Location)
*/


/*
	query := `CREATE TABLE IF NOT EXISTS product(product_id int primary key auto_increment, product_name text, 
	product_price int, created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP)`
	 
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    res, err := db.ExecContext(ctx, query)
    if err != nil {
        log.Printf("Error %s when creating product table", err)
        return err
    }
    rows, err := res.RowsAffected()
    if err != nil {
        log.Printf("Error %s when getting rows affected", err)
        return err
    }
    log.Printf("Rows affected when creating table: %d", rows)
    return nil
*/


/*
func GetAllWeatherData(w http.ResponseWriter, r *http.Request){
    var condb *sql.DB
    condb, errdb := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
    if errdb != nil {
        fmt.Println(" Error open db:", errdb.Error())
      }
	GetWeatherDataFromDB(condb)
	defer condb.Close()

    fmt.Println("Endpoint Hit: GetAllWeatherData")
	tmpl.ExecuteTemplate(w, "Index", AllWeatherData)
    //json.NewEncoder(w).Encode(AllWeatherData)
  }

func GetWeatherDataFromDB(condb *sql.DB){
    var eachrow WeatherData
    sqlstmt:="select WeatherID , Location, WeatherDate, temperature from weatherdata "
    rows, err := condb.Query(sqlstmt)
    if err != nil {
        panic(err)
     }
     for rows.Next(){
         rows.Scan(&eachrow.WeatherID, &eachrow.WeatherDate, &eachrow.Location, &eachrow.Temp )
         AllWeatherData = append(AllWeatherData,eachrow)
		 log.Println(eachrow.Location)
     }
}
 
func GetAllWeatherDataByID(w http.ResponseWriter, r *http.Request){
    var condb *sql.DB
	//id,err = GetID(w,r)
	var id =1
    condb, errdb := sql.Open("mssql", "Data Source=host.docker.internal,1433;database=WeatherDB;User ID=sa;Password=someThingComplicated1234")
    if errdb != nil {
        fmt.Println(" Error open db:", errdb.Error())
      }
	log.Println("filterid =" , id)
	GetFilteredWeatherDataFromDB(id,condb)
    fmt.Println("Endpoint Hit: GetAllWeatherDataByID")
    json.NewEncoder(w).Encode(AllWeatherData)
  }

func GetFilteredWeatherDataFromDB(wid int, condb *sql.DB)(WeatherData, error){
    var eachrow WeatherData
    sqlstmt:="select WeatherID , Location, WeatherDate, temperature from weatherdata where weatherid=?"
    row := condb.QueryRow(sqlstmt, wid)

	if err := row.Scan(&eachrow.WeatherID, &eachrow.Location, &eachrow.WeatherDate, &eachrow.Temp); err != nil {
       if err == sql.ErrNoRows {
            return eachrow, fmt.Errorf("WeatherID %d: no such data", wid)
        }
        return eachrow, fmt.Errorf("WeatherID %d: %v", wid, err)
    }
	         AllWeatherData = append(AllWeatherData,eachrow)

    return eachrow, nil
}

func addNewWeatherData(newData WeatherData,condb *sql.DB) (int64, error) {
    result, err := condb.Exec("INSERT INTO Weatherdata (weatherdate, location,temp) VALUES (?, ?, ?)", newData.WeatherDate, newData.Location, newData.Temp)
    if err != nil {
        return 0, fmt.Errorf("addNewWeatherData: %v", err)
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, fmt.Errorf("addNewWeatherData: %v", err)
    }
    return id, nil
}
///////////////////////////////////////////////
*/