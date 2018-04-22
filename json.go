package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"realeased"`
	Color  bool `json:"color, omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablance", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.MarshalIndent(movies, "", " ")

	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("Json unmarshaling failed: %s", err)
	}
	//fmt.Println(err)
	fmt.Println(titles)
}
