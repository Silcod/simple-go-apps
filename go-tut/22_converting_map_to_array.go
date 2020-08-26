// Converting a Map into an Array of Keys and Values

package main

import (
  "fmt"
)

type NameAge struct{ // A struct in Golang is a user-defined type, that allows to group items of possibly different types into a single type
                    // Any real-world entity which has some set of properties/fields can be represented as a struct.
                    // This concept is generally compared with the classes in object-oriented programming.
                    // It can be termed as a lightweight class that does not support inheritance but supports composition.
    Name string
    Age int
}

func main(){
    var nameAgeSlice []NameAge
    nameAges := map[string]int{ // Here, we map string values to integer values, Map is like dictionaries in other programming languages
      "Micheal": 30,
      "John": 25,
      "Jessica": 26,
      "Ali": 18,
    }

    for key, value := range nameAges{
        nameAgeSlice = append(nameAgeSlice, NameAge {key, value}) // name is the key, value is the age
    }

    fmt.Println(nameAgeSlice)
}
