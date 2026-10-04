package main

import "fmt"

type Movie struct {
	Title    string
	Duration int
	Rating   float64
}

func main() {
	movie1 := Movie{
		Title:    "Avengers",
		Duration: 210,
		Rating:   9.8,
	}

	movie2 := Movie{
		Title:    "Spider Man",
		Duration: 180,
		Rating:   9.2,
	}

	movie3 := Movie{
		Title:    "Batman",
		Duration: 150,
		Rating:   7.9,
	}

	fmt.Printf("%s %.1f %d\n", movie1.Title, movie1.Rating, movie1.Duration)
	fmt.Printf("%s %.1f %d\n", movie2.Title, movie2.Rating, movie2.Duration)
	fmt.Printf("%s %.1f %d\n", movie3.Title, movie3.Rating, movie3.Duration)

}
