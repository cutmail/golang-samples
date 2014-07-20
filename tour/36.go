package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
	z := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		dz := make([]uint8, dy)
		for y := 0; y < dy; y++ {
			dz[y] = uint8((x + y) / 2)
		}
		z[x] = dz
	}
	return z
}

func main() {
	pic.Show(Pic)
}
