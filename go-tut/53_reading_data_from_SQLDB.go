// Reading Data From SQL Databases

package main

import (
    "fmt"
    _ "github.com/mattn/go-sqlite3"
    "database/sql"
)

type Profile struct{
    ProfileId int
    FirstName string
    LastName string
    Age int
}

func main(){
    db, err := sql.Open("sqlite3", "./personal.db") // sqlite3 is (driverName), ./personal.db is (dataSourceName)
    checkError(err)

    /*
      db.Query() syntax:
      db.Query(query string, args....interface{})
      Query executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters  in the query
      or when the argument has a variable value
    */

    var profile Profile
    ab2 := 2
    rows, err := db.Query("select ProfileId, FirstName, LastName, Age from Profile where ProfileId = ?", ab2)
    checkError(err)

    for rows.Next(){
        err := rows.Scan(&profile.ProfileId, &profile.FirstName, &profile.LastName, &profile.Age)
        checkError(err)

        fmt.Println(profile)
    }
    rows.Close()
    db.Close()
}

func checkError(err error){
    if (err != nil) {
      panic(err)
    }
}
