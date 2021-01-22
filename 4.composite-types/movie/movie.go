package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// Actor is a data structure to store actor data.
type Actor struct {
	Name    string
	Sirname string
	DoB     time.Time
}

// Movie is a data structure to store movie data.
type Movie struct {
	Title  string
	Year   int  `json:"Released"`
	Color  bool `json:"Color,omitempty"`
	Actors []Actor
}

func main() {
	var movies = []Movie{
		Movie{Title: "Casablanca", Year: 1942, Color: false, Actors: []Actor{
			Actor{
				Name:    "Bogart",
				Sirname: "Humphrey",
				DoB:     time.Date(1912, 3, 2, 17, 0, 0, 0, &time.Location{}),
			}, Actor{
				Name:    "Bergman",
				Sirname: "Ingrid",
				DoB:     time.Date(1902, 7, 25, 9, 0, 0, 0, &time.Location{}),
			},
		}},
		Movie{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []Actor{
			Actor{
				Name:    "Newman",
				Sirname: "Paul",
				DoB:     time.Date(1935, 7, 12, 19, 0, 0, 0, &time.Location{}),
			},
		}},
		Movie{Title: "Bullitt", Year: 1968, Color: true, Actors: []Actor{
			Actor{
				Name:    "McQueen",
				Sirname: "Steve",
				DoB:     time.Date(1928, 4, 19, 16, 0, 0, 0, &time.Location{}),
			}, Actor{
				Name:    "Bisset",
				Sirname: "Jacqueline",
				DoB:     time.Date(1930, 11, 18, 13, 0, 0, 0, &time.Location{}),
			},
		}},
	}

	// Marshalling
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON encoding failed: %s", err)
	}

	fmt.Printf("%s\n", data)

	// Unmarshalling
	type Title struct {
		Title string
	}

	var titles []Title
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON decoding failed: %s", err)
	}

	fmt.Printf("%#v\n", titles)
}
