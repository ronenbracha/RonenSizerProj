package imaging

import (
	"fmt"
	"image"
	"testing"
)

func TestResize(t *testing.T) {
	testCases := []struct {
		name string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Resize 2x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Resize 2x2 1x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0x80,
					0x00, 0x80, 0x80, 0xff,
				},
			},
		},
		{
			"Resize 2x2 2x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			2, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0xff, 0x00, 0x80, 0x80, 0x00, 0x80, 0xff,
				},
			},
		},
		{
			"Resize 2x2 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			2, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 3x1 1x1 nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 2, 0),
				Stride: 3 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 1,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x00, 0xff, 0x00, 0xff},
			},
		},
		{
			"Resize 2x2 0x4 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 4,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 2x2 4x0 linear",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			4, 0,
			Linear,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x40, 0xff, 0x00, 0x00, 0xbf, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0x40, 0x6e, 0x6d, 0x25, 0x70, 0xb0, 0x14, 0x3b, 0xcf, 0xbf, 0x00, 0x40, 0xff,
					0x00, 0xff, 0x00, 0xbf, 0x14, 0xb0, 0x3b, 0xcf, 0x33, 0x33, 0x99, 0xef, 0x40, 0x00, 0xbf, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0xbf, 0x40, 0xff, 0x00, 0x40, 0xbf, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Resize 0x0 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, -1, -1),
				Stride: 0,
				Pix:    []uint8{},
			},
			1, 1,
			Box,
			&image.NRGBA{},
		},
		{
			"Resize 2x2 0x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 0,
			Box,
			&image.NRGBA{},
		},
		{
			"Resize 2x2 -1x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			-1, 0,
			Box,
			&image.NRGBA{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Resize(tc.src, tc.w, tc.h, tc.f)
			if !compareNRGBA(got, tc.want, 0) {
				t.Fatalf("got result %#v want %#v", got, tc.want)
			}
		})
	}
}

func TestResampleFilters(t *testing.T) {
	for _, filter := range []ResampleFilter{
		NearestNeighbor,
		Box,
		Linear,
		Hermite,
		MitchellNetravali,
		CatmullRom,
		BSpline,
		Gaussian,
		Lanczos,
		Hann,
		Hamming,
		Blackman,
		Bartlett,
		Welch,
		Cosine,
	} {
		t.Run("", func(t *testing.T) {
			src := image.NewNRGBA(image.Rect(-1, -1, 2, 3))
			got := Resize(src, 5, 6, filter)
			want := image.NewNRGBA(image.Rect(0, 0, 5, 6))
			if !compareNRGBA(got, want, 0) {
				t.Fatalf("got result %#v want %#v", got, want)
			}
			if filter.Kernel != nil {
				if x := filter.Kernel(filter.Support + 0.0001); x != 0 {
					t.Fatalf("got kernel value %f want 0", x)
				}
			}
		})
	}
}

func TestResizeGolden(t *testing.T) {
	for name, filter := range map[string]ResampleFilter{
		"out_resize_nearest.png": NearestNeighbor,
		"out_resize_linear.png":  Linear,
		"out_resize_catrom.png":  CatmullRom,
		"out_resize_lanczos.png": Lanczos,
	} {
		got := Resize(testdataBranchesPNG, 150, 0, filter)
		want, err := Open("testdata/" + name)
		if err != nil {
			t.Fatalf("failed to open image: %v", err)
		}
		if !compareNRGBA(got, toNRGBA(want), 0) {
			t.Fatalf("resulting image differs from golden: %s", name)
		}
	}
}

func TestFit(t *testing.T) {
	testCases := []struct {
		name string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Fit 2x2 1x10 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			1, 10,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Fit 2x2 10x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			10, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Fit 2x2 10x10 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			10, 10,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
		},
		{
			"Fit 0x0 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, -1, -1),
				Stride: 0,
				Pix:    []uint8{},
			},
			1, 1,
			Box,
			&image.NRGBA{},
		},
		{
			"Fit 2x2 0x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			0, 0,
			Box,
			&image.NRGBA{},
		},
		{
			"Fit 2x2 -1x0 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 1),
				Stride: 2 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
				},
			},
			-1, 0,
			Box,
			&image.NRGBA{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Fit(tc.src, tc.w, tc.h, tc.f)
			if !compareNRGBA(got, tc.want, 0) {
				t.Fatalf("got result %#v want %#v", got, tc.want)
			}
		})
	}
}

func TestFill(t *testing.T) {
	testCases := []struct {
		name string
		src  image.Image
		w, h int
		a    Anchor
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Fill 4x4 2x2 Center Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			2, 2,
			Center,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0x14, 0x15, 0x16, 0x17, 0x1c, 0x1d, 0x1e, 0x1f,
					0x34, 0x35, 0x36, 0x37, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
		},
		{
			"Fill 4x4 1x4 TopLeft Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			1, 4,
			TopLeft,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 4),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03,
					0x10, 0x11, 0x12, 0x13,
					0x20, 0x21, 0x22, 0x23,
					0x30, 0x31, 0x32, 0x33,
				},
			},
		},
		{
			"Fill 4x4 8x2 Bottom Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			8, 2,
			Bottom,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 8, 2),
				Stride: 8 * 4,
				Pix: []uint8{
					0x30, 0x31, 0x32, 0x33, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f, 0x3c, 0x3d, 0x3e, 0x3f,
					0x30, 0x31, 0x32, 0x33, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
		},
		{
			"Fill 4x4 2x8 Top Nearest",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			2, 8,
			Top,
			NearestNeighbor,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 8),
				Stride: 2 * 4,
				Pix: []uint8{
					0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b,
					0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b,
					0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b,
					0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b,
					0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b,
					0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b,
					0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b,
				},
			},
		},
		{
			"Fill 4x4 4x4 TopRight Box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			4, 4,
			TopRight,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 4, 4),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
		},
		{
			"Fill 4x4 0x4 Left Box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 3, 3),
				Stride: 4 * 4,
				Pix: []uint8{
					0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f,
					0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f,
					0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f,
					0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3a, 0x3b, 0x3c, 0x3d, 0x3e, 0x3f,
				},
			},
			0, 4,
			Left,
			Box,
			&image.NRGBA{},
		},
		{
			"Fill 0x0 4x4 Right Box",
			&image.NRGBA{},
			4, 4,
			Right,
			Box,
			&image.NRGBA{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Fill(tc.src, tc.w, tc.h, tc.a, tc.f)
			if !compareNRGBA(got, tc.want, 0) {
				t.Fatalf("got result %#v want %#v", got, tc.want)
			}
		})
	}
}

func TestThumbnail(t *testing.T) {
	testCases := []struct {
		name string
		src  image.Image
		w, h int
		f    ResampleFilter
		want *image.NRGBA
	}{
		{
			"Thumbnail 6x2 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 5, 1),
				Stride: 6 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Thumbnail 2x6 1x1 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 1, 5),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0xff,
					0x00, 0xff, 0x00, 0xff, 0x00, 0x00, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
					0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
				},
			},
			1, 1,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 1, 1),
				Stride: 1 * 4,
				Pix:    []uint8{0x55, 0x55, 0x55, 0xc0},
			},
		},
		{
			"Thumbnail 1x3 2x2 box",
			&image.NRGBA{
				Rect:   image.Rect(-1, -1, 0, 2),
				Stride: 1 * 4,
				Pix: []uint8{
					0x00, 0x00, 0x00, 0x00,
					0xff, 0x00, 0x00, 0xff,
					0xff, 0xff, 0xff, 0xff,
				},
			},
			2, 2,
			Box,
			&image.NRGBA{
				Rect:   image.Rect(0, 0, 2, 2),
				Stride: 2 * 4,
				Pix: []uint8{
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
					0xff, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00, 0xff,
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Thumbnail(tc.src, tc.w, tc.h, tc.f)
			if !compareNRGBA(got, tc.want, 0) {
				t.Fatalf("got result %#v want %#v", got, tc.want)
			}
		})
	}
}

func BenchmarkResize(b *testing.B) {
	for _, dir := range []string{"Down", "Up"} {
		for _, filter := range []string{"NearestNeighbor", "Linear", "CatmullRom", "Lanczos"} {
			for _, format := range []string{"JPEG", "PNG"} {
				var size int
				switch dir {
				case "Down":
					size = 100
				case "Up":
					size = 1000
				}

				var f ResampleFilter
				switch filter {
				case "NearestNeighbor":
					f = NearestNeighbor
				case "Linear":
					f = Linear
				case "CatmullRom":
					f = CatmullRom
				case "Lanczos":
					f = Lanczos
				}

				var img image.Image
				switch format {
				case "JPEG":
					img = testdataBranchesJPG
				case "PNG":
					img = testdataBranchesPNG
				}

				b.Run(fmt.Sprintf("%s %s %s", dir, filter, format), func(b *testing.B) {
					b.ReportAllocs()
					for i := 0; i < b.N; i++ {
						Resize(img, size, size, f)
					}
				})
			}
		}
	}
}
