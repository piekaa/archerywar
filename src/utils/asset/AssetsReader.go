package asset

import (
	"golang.org/x/mobile/asset"
	"github.com/piekaa/archerywar/src/collections/dynamic"
)

func ReadString(name string) (string, error) {
	bytes, err :=  ReadBytes(name)
	if err != nil {
		return  "", err
	}
	return string(bytes), nil
}


func ReadBytes(name string) ([]byte, error) {
	file, err  := asset.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dArr := dynamic.New()
	//todo change that to greater size
	bufferSize := 10
	for {
		buffer := make([]byte, bufferSize)
		n, err := file.Read(buffer)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		if n == 0 {
			break
		}
		dArr.AddN(buffer, n)
	}
	if err != nil {
		return nil, err
	}
	return dArr.Bytes, nil
}
