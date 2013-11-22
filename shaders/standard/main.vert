#version 330

// Inputs
layout(location = 0) in vec4 position;

void main() {

	vec4 vertex = position;

	gl_Position = vertex;
}
