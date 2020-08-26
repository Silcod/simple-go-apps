// One of the ways of talking to MongoDB from golang is using the "mgo" package
package main

import (
    "github.com/globalsign/mgo"
    log "github.com/sirupsen/logrus"
    "fmt"
)

const (
    url = "localhost"
)

func main() {
    // connecting to mongodb server using mgo package

    // Note: The mgo("pronounced as mango") package offers a rich MongoDB driver for Go.
    // Detailed documentation of the package is available at GoDoc: https://godoc.org/github.com/globalsign/mgo
    // Usage of the driver revolves around the concept of sessions.
    session, err := mgo.Dial(url)
    if err != nil {
        log.Fatal(err)
    }

    defer session.Close()

    fmt.Printf("Successfully connected to mongodb server at %v", url)

    dbNames, err := session.DatabaseNames()
    if err != nil {
        log.Warn(err)
    }
    for i, v := range dbNames {
        fmt.Printf("[%3v] - %v\n", i + 1, v)
    }
}
