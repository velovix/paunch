#version 330

// Inputs
layout(location = 0) in vec4 position;
layout(location = 1) in vec2 texcoord;

// Output
smooth out vec2 f_texcoord;

void main() {

	gl_Position = position;
	f_texcoord = texcoord;
}
