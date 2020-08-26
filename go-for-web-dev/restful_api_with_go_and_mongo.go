package main

import (
  "context"
  "fmt"
  "net/http"
  "time"
  "encoding/json"

  "github.com/gorilla/mux"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
  // Note bson is used to marshal and unmarshal our data structure into something mongodb understands
  // While json is used to marshal and unmarshal our data structure into something our web browser understands
    ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
    Firstname string `json:"firstname,omitempty" bson:"firstname,omitempty"`
    Lastname string `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func CreatePersonEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Add("content-type", "application/json")
    var person Person
    json.NewDecoder(request.Body).Decode(&person)
    collection := client.Database("thepolyglotdeveloper").Collection("people")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    result, _ := collection.InsertOne(ctx, person)
    json.NewEncoder(response).Encode(result)
}

func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
    response.Header().Add("content-type", "application/json")
    var people []Person
    collection := client.Database("thepolyglotdeveloper").Collection("people") // Note ctx, collection, then collection.Find or collection.Insert
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        response.WriteHeader(http.StatusInternalServerError)
        response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
        return
    }
    defer cursor.Close(ctx)

    for cursor.Next(ctx) {
        var person Person
        cursor.Decode(&person)
        people = append(people, person)
    }

    if err := cursor.Err(); err != nil {
      response.WriteHeader(http.StatusInternalServerError)
      response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
      return
    }
    json.NewEncoder(response).Encode(people)
}


func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) {}


func main() {
    fmt.Println("Starting the application...")
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)  //(parent, timeout)
    client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    router := mux.NewRouter()
    router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
    router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
    router.HandleFunc("/person/{id}", GetPersonEndpoint).Methods("GET")
    http.ListenAndServe(":12345", router)
}
