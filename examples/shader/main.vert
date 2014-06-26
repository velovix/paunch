#version 330

// Inputs
layout(location = 0) in vec4 position;
layout(location = 1) in vec2 texcoord;

// Outputs
smooth out vec2 f_texcoord;

uniform vec2 screen_size;

vec4 adjustScreenPixels(vec4 vertex) {

	if(screen_size.x != 0 && screen_size.y != 0) {
		vertex.xy /= screen_size.xy;
		vertex.xy *= vec2(2, 2);
		vertex.xy -= vec2(1, 1);
	}

	return vertex;
}

void main() {

	vec4 vertex = position;

	vertex = adjustScreenPixels(vertex);

	f_texcoord = texcoord;
	gl_Position = vertex;
}
