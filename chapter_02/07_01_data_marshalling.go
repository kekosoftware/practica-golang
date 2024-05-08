package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Marshal a Person object to JSON
	p := Person{Name: "John Doe", Age: 30}
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Error marshalling to JSON: %s", err)
	}
	fmt.Println(string(jsonData))
	// Unmarshal JSON to a Person object
	var p2 Person
	err = json.Unmarshal(jsonData, &p2)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %s", err)
	}
	fmt.Println(p2)
}