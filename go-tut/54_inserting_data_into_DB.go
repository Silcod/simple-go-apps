// Inserting Data Into Database

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

    statement, err := db.Prepare("insert into Profile (FirstName, LastName, Age) values(?,?,?)") // ? are called placeholders
    checkError(err)

    statement.Exec("Jessica", "McArthur", 30)

    var profile Profile
    rows, err := db.Query("select ProfileId, FirstName, LastName, Age from Profile")
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
