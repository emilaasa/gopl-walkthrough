package eajson

import (
	"encoding/json"
	"fmt"
	"log"
)

// MarshalSomeJson is a basic json marshalling operation from the book
func MarshalSomeJson() {
	// NOTE: Pretty printing for convenience, mostly use Marshal unless presenting
	// json in the terminal
	data, err := json.MarshalIndent(movies, "", "   ")
	if err != nil {
		log.Fatalf("JSON Marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

//Movie json struct lintttt
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 192, Color: false,
		Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true,
		Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true,
		Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
}
