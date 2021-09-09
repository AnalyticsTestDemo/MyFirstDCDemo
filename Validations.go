package main
import (
	//"database/sql"
	//"encoding/json"
//	"fmt"
)

//https://golang.org/doc/tutorial/database-access
//https://tutorialedge.net/golang/golang-orm-tutorial/

//how to dockerize golang app 
//https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golang-application-80f3fb59a15e

//https://www.golangprograms.com/example-of-golang-crud-using-mysql-from-scratch.html

 
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
