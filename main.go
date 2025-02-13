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

	// Open a file for writing
	file, err := os.Create("types.ts")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Generate TypeScript types
	generator := go2ts.New()
	generator.Add(Turtle{})
	generator.AddUnion(AllDirections)
	generator.Render(file) // Write output to file instead of Stdout
}