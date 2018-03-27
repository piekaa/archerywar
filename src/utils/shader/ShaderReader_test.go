package shader

import (
	"testing"
	"strings"
)





func TestRead(t *testing.T) {

	vertexShader := `//vertex shader
#version 330 core
layout (location = 0) in vec4 position;

void main() {
	gl_Position = position;
}`


fragmentShader := `//fragment shader
#version 330 core
layout (location = 0) out vec4 color;

void main() {
	color = vec4(1.0, 1.0, 1.0, 1.0);
}`

	vs, fs, err := Read("test")

	if err != nil {
		t.Error(err)
	}

	vs =strings.TrimSpace(vs)
	fs =strings.TrimSpace(vs)
	vertexShader =strings.TrimSpace(vs)
	fragmentShader =strings.TrimSpace(vs)


	if vs != vertexShader {
		t.Errorf("%s\n!=\n%s", vs, vertexShader)
	}

	if fs != fragmentShader {
		t.Errorf("%s\n!=\n%s", fs, fragmentShader)
	}
}