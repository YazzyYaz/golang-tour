package main

import (
	"fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var image = make([][]uint8, dy)
	fmt.Println(image)
	for i := range image {
		image[i] = make([]uint8, dx)
		for j := range image[i] {
			image[i][j] = uint8(i * j)
		}
	}
	return image
}

func main() {
	pic.Show(Pic)
}
