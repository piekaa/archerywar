//vertex shader
#version 330 core 
attribute vec4 position; 
attribute vec2 texCoord;
out vec2 v_TexCoord;
void main() {
	gl_Position = position;
	// gl_Position = tex;
	v_TexCoord = texCoord;
}

//!@#$

//fragment shader
#version 330 core
// layout (location = 0) out vec4 color;

in vec2 v_TexCoord;
uniform sampler2D u_Texture;
uniform vec4 u_Color;

void main() {
	vec4 texColor = texture(u_Texture, v_TexCoord);
	gl_FragColor = texColor;
}