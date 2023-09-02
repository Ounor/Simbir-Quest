package main

import (
	"github.com/Ounor/Simbir-Quest/internal/welcome"
	_ "image/png"
)

const (
	screenWidth  = 1280
	screenHeight = 720
	title        = "Simbir-Quest"
)

func main() {
	welcome.Show(screenWidth, screenHeight, title)
}
