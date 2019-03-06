package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) (imageResult [][]uint8) {
	imageResult = make([][]uint8, dy)
	for index := 0; index < dy; index++ {
		imageResult[index] = make([]uint8, dx)
	}
	return
}

func main() {
	pic.Show(Pic)
}
