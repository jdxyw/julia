package julia

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math/cmplx"
	"os"
)

type GenFunc func(complex128) complex128

type julia struct {
	h, w int
	x, y float64
	fn   GenFunc
	iter int
	img  *image.RGBA
}

func NewJulia(h, w int, x, y float64, iter int, fn GenFunc) *julia {
	return &julia{
		h:    h,
		w:    w,
		x:    x,
		y:    y,
		fn:   fn,
		iter: iter,
		img:  image.NewRGBA(image.Rect(0, 0, h, w)),
	}
}

func (j *julia) SetFunc(fn GenFunc) {
	j.fn = fn
}

func (j *julia) CleanImage() {
	j.img = nil
	j.img = image.NewRGBA(image.Rect(0, 0, j.h, j.w))
}

func (j *julia) GenerativeGray(maxz float64) {

	for i := 0; i <= j.w; i++ {
		for k := 0; k <= j.h; k++ {
			nit := 0
			z := complex(float64(i)/float64(j.w)*2.0*j.x-j.x, float64(k)/float64(j.h)*2.0*j.y-j.y)

			for cmplx.Abs(z) <= maxz && nit < j.iter {
				z = j.fn(z)
				nit += 1
			}
			ratio := uint8(nit * 255 / j.iter)
			j.img.Set(i, k, color.RGBA{
				R: ratio,
				G: ratio,
				B: ratio,
				A: 255,
			})
		}
	}
}

func (j *julia) ToPng(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := png.Encode(f, j.img); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
}

func (j *julia) ToJpeg(path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := jpeg.Encode(f, j.img, nil); err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
}
