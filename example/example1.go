package main

import (
	"julia"
)

func julia1(z complex128) complex128 {
	c := complex(-0.1, 0.65)

	z = z*z + c

	return z
}

func main() {
	j := julia.NewJulia(800, 800, 1.5, 1.5, 1000, 10, julia1)
	j.Generative(julia.Twilight)
	j.ToPng("julia1.png")
}
