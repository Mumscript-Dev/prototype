package main

import (
	"image/color"
	"os"

	"github.com/skia-dev/go2ts"
)

func main() {
	type Direction string

	const (
		Up    Direction = "up"
		Down  Direction = "down"
		Left  Direction = "left"
		Right Direction = "right"
	)

	AllDirections := []Direction{Up, Down, Left, Right}

	type Position struct {
		X int
		Y int
	}

	type Turtle struct {
		Position  `json:"Coordinates"`
		Color     color.Alpha
		Direction Direction
	}

	type Turtle2 struct {
		Position  `json:"Coordinates"`
		Color     color.Alpha
		Direction Direction
	}

	type User struct {
		Name string
		Age  int
	}
	type Pie struct {
		Title string `json:"title"`
		Value int    `json:"value"`
	}

	// Open a file for writing
	file, err := os.Create("types/types.ts")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Generate TypeScript types
	generator := go2ts.New()
	generator.Add(Turtle{})
	generator.Add(Turtle2{})
	generator.Add(User{})
	generator.Add(Pie{})
	generator.AddUnion(AllDirections)
	generator.Render(file) 
}