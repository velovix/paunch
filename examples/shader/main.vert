#version 120

// Inputs
attribute vec4 position;
attribute vec2 texcoord;

// Outputs
varying vec2 f_texcoord;

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

	gl_TexCoord[0].st = texcoord;
	f_texcoord = texcoord;
	gl_Position = vertex;
}
