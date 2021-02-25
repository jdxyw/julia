//Package julia provides interface to draw a julia set with customize complementary sets and color map.
package julia

import (
	"image"
	"image/jpeg"
	"image/png"
	"math/cmplx"
	"os"
)

// GenFunc defines a func type used by julia set.
type GenFunc func(complex128) complex128

type julia struct {
	h, w int
	x, y float64
	maxz float64
	fn   GenFunc
	iter int
	img  *image.RGBA
}

// NewJulia returns a julia struct.
func NewJulia(h, w int, x, y float64, iter int, z float64, fn GenFunc) *julia {
	return &julia{
		h:    h,
		w:    w,
		x:    x,
		y:    y,
		fn:   fn,
		iter: iter,
		maxz: z,
		img:  image.NewRGBA(image.Rect(0, 0, h, w)),
	}
}

// CleanImage clean the image and create a new RGBA object.
func (j *julia) CleanImage() {
	j.img = nil
	j.img = image.NewRGBA(image.Rect(0, 0, j.h, j.w))
}

// Generative draws the julia set with specified color map.
func (j *julia) Generative(cm ColorMap) {

	for i := 0; i <= j.w; i++ {
		for k := 0; k <= j.h; k++ {
			nit := 0
			z := complex(float64(i)/float64(j.w)*2.0*j.x-j.x, float64(k)/float64(j.h)*2.0*j.y-j.y)

			for cmplx.Abs(z) <= j.maxz && nit < j.iter {
				z = j.fn(z)
				nit += 1
			}
			idx := uint8(nit*255/j.iter) % uint8(len(cm)-1)
			j.img.Set(i, k, cm[idx])
		}
	}
}

// ToPng saves the image to local with PNG format.
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

	return nil
}

// ToJpeg saves the image to local with Jpeg format.
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

	return nil
}
