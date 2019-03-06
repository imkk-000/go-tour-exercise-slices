package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// initial empty image matrix
	imageResult := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		imageResult[y] = make([]uint8, dx)
	}
	// compute image in matrix
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			imageResult[y][x] = uint8((x + y) / 2)
		}
	}
	return imageResult
}

func main() {
	pic.Show(Pic)
}
