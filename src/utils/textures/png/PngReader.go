package png

import (
	"image/png"
	"github.com/piekaa/archerywar/src/utils/asset"
	"image"
	"bytes"
)

//reads png file from assets/textures directory, file name must end with .png but don't put that in name argument
func Read(name string) (image.Image, error) {

	b, err := asset.ReadBytes("textures/"+name+".png")

	if err != nil {
		return nil, err
	}

	img, err := png.Decode( bytes.NewReader(b) )

	if err != nil {
		return nil, err
	}


	return img, err
}


//reads png file from assets/textures directory, file name must end with .png but don't put that in name argument
//
//Reads as byte array RGBA 0,255
//Returns bytes, width, height, error
func ReadAsBytes(name string) (b []byte, w, h int, err error) {

	img, err := Read(name)

	if err != nil {
		return nil,0,0, err
	}

	w = img.Bounds().Dx()
	h = img.Bounds().Dy()

	ib := make([]byte, w*h*4)

	for j := 0 ; j < h ;j ++ {
		for i := 0 ;i < w ; i++ {
			r,g,b,a := img.At(i,h-1-j).RGBA()
			ib[ j*w*4 + i*4 + 0] = byte(r)
			ib[ j*w*4 + i*4 + 1] = byte(g)
			ib[ j*w*4 + i*4 + 2] = byte(b)
			ib[ j*w*4 + i*4 + 3] = byte(a)
		}
	}
	return ib,w,h, nil

}