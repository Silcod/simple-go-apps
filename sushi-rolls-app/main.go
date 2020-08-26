package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Roll struct {
	ID          string `json:"id,omitempty"`
	ImageNumber string `json:"image_number,omitempty"`
	Name        string `json:"name,omitempty"`
	Ingredients string `json:"ingredients,omitempty"`
}

var rolls []Roll

// getRolls function: The first thing we do is set the headers on the response. Then we use the NewEncoder() and Encode() methods
// to render our rolls slice as json and send it to the response stream.

func getRolls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rolls) // perhaps w because this function is supposed to give a response
}

func getRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // perhaps r because you are supposed to get an id from the user in the request
	for _, item := range rolls {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// The first thing we need to do is create an instance of our Roll struct, as you can see above I’ve called it newRoll.
// Now let’s get some data from the HTTP request and put it in there! To get the data out of our response, we need to
// head back to the json library that we imported. Much like we’ve been using .NewEncoder() and .Encode() to send our
// response; we will use .NewDecoder() and .Decode() to read data from our requests. We call json.NewDecoder() and
// pass it the body of our HTTP request. The decoder reads the data, and we call .Decode() on that passing it a pointer
// to our newRoll struct which allows it to match the json to the appropriate properties of the struct.

// Then to simulate the database we assign it an id, by getting the length of the rolls slice and then calling
// strconv.Itoa() on it which converts it to a string.

// Now that we have our newRoll assembled we need to add it to our slice. To do that we use append(). Append takes
// at least two arguments, the first is the slice we want to add items to and then we can add multiple arguments of
// the correct data type to add them to the slice.
// Finally, we send a response back containing our new roll, just like before.

func createRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newRoll Roll
	json.NewDecoder(r.Body).Decode(&newRoll)
	newRoll.ID = strconv.Itoa(len(rolls) + 1)
	rolls = append(rolls, newRoll)
	json.NewEncoder(w).Encode(newRoll)
}

func updateRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			var newRoll Roll
			json.NewDecoder(r.Body).Decode(&newRoll)
			newRoll.ID = params["id"]
			rolls = append(rolls, newRoll)
			json.NewEncoder(w).Encode(newRoll)
			return
		}
	}
}

func deleteRoll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, item := range rolls {
		if item.ID == params["id"] {
			rolls = append(rolls[:i], rolls[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(rolls)
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/sushi", getRolls).Methods("GET")
	router.HandleFunc("/sushi/{id}", getRoll).Methods("GET")
	router.HandleFunc("/sushi", createRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", updateRoll).Methods("POST")
	router.HandleFunc("/sushi/{id}", deleteRoll).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", router))
}
