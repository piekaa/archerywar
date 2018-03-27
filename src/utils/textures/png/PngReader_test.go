package png

import (
	"testing"
)

func TestRead(t *testing.T) {

	img, err := Read("body")

	if err != nil {
		t.Error(err)
	}




	if img.Bounds().Dx() != 300 {
		t.Errorf("width is incorrect %d:", img.Bounds().Dx())
	}
	if img.Bounds().Dy() != 800 {
		t.Errorf("height is incorrect%d:", img.Bounds().Dy())
	}

}

func TestReadAsBytes(t *testing.T) {

	b, w,h, err := ReadAsBytes("body")

	if err != nil {
		t.Error(err)
	}

	if w != 300 {
		t.Errorf("Width is %s", w)
	}

	if h != 800 {
		t.Errorf("Height is %s", h)
	}

	if len(b) != 800 * 300 * 4 {
		t.Errorf("Wrong number of bytes %d", len(b))
	}

}


func TestColor(t *testing.T) {

	b,_, _, err := ReadAsBytes("test")

	if err != nil {
		t.Error(err)
	}

	if b[16] != 126 {
		t.Errorf("R: %d", b[16] )
	}
	if b[17] != 162 {
		t.Errorf("G: %d", b[17] )
	}
	if b[18] != 50 {
		t.Errorf("B: %d", b[18] )
	}
	if b[19] != 255 {
		t.Errorf("A: %d", b[19] )
	}

	if b[128+16] != 126 {
		t.Errorf("R: %d", b[128+16] )
	}
	if b[128+17] != 162 {
		t.Errorf("G: %d", b[128+17] )
	}
	if b[128+18] != 50 {
		t.Errorf("B: %d", b[128+18] )
	}
	if b[128+19] != 255 {
		t.Errorf("A: %d", b[128+19] )
	}


}