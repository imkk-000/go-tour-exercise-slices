package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	imageResult := make([][]uint8, dy)
	for index := 0; index < dy; index++ {
		imageResult[index] = make([]uint8, dx)
	}
	return imageResult
}

func main() {
	pic.Show(Pic)
}
