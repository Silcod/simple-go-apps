package main

import (
	"context"
	"log"

	"github.com/kr/pretty"
	"googlemaps.github.io/maps"
)

type DirectionsRequest struct {
	Origin string `json:"origin"`
	Destination string `json:"destination"`
}

func main() {
	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyBRcp8CWfR7KH6C10iPQ64XtdO-tiu5cjA"))
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}
	r := &DirectionsRequest{
		Origin:      "7b, adekunle fajuyi way, ikeja Gra, Lagos, Nigeria",
		Destination: "10, Balogun street, oshodi, Lagos, Nigeria",
	}
	route, _, err := c.Directions(context.Background(), r) // route, _, err because this function returns three different values
	if err != nil {
		log.Fatalf("fatal error: %s", err)
	}

	pretty.Println(route)
}