#version 330

// Inputs
layout(location = 0) in vec4 position;
layout(location = 1) in vec2 texcoord;

// Outputs
smooth out vec2 f_texcoord;

void main() {

	vec4 vertex = position;

	f_texcoord = texcoord;
	gl_Position = vertex;
}
