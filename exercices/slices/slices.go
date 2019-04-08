package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	buf := make([][]uint8, dy)
	for y := range buf {
		xs := make([]uint8, dx)
		for x := range xs {
			// xs[x] = uint8((x + y) / 2)
			// xs[x] = uint8(x * y)
			xs[x] = uint8(x + y)
		}
		buf[y] = xs
	}
	return buf
}

func main() {
	pic.Show(Pic)
}
