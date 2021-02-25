package main

import (
	"julia"
)

func julia1(z complex128) complex128 {
	c := complex(-0.1, 0.651)

	z = z*z + c

	return z
}

func main() {
	j := julia.NewJulia(800, 800, 1.5, 1.5, 800, 40, julia1)
	j.Generative(julia.Inferno)
	j.ToPng("julia4.png")
}
