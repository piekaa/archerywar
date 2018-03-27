package shader

import (
	"github.com/piekaa/archerywar/src/utils/asset"
	"strings"
)
//file must end with .shader, if file is basic.shader, name should be "basic"
//
//file must be in assets/shaders/ directory
//
//returns vertex shader, fragment shader, error
func Read(name string) (vs, fs string, err error) {
	src, err:= asset.ReadString("shaders/" + name + ".shader")
	if err != nil {
		return "", "", err
	}
	s := strings.Split(src, "//!@#$")
	return s[0], s[1], nil
}