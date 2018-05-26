package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"1", "2"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("json marshling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	dataFormat, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		log.Fatalf("json marshling failed: %s", err)
	}
	fmt.Printf("%s\n", dataFormat)

	var title []struct{ Title string }

	if err := json.Unmarshal(data, &title); err != nil {
		log.Fatalf("JSON unmarshl failed: %s", err)
	}
	fmt.Println(title)
}
