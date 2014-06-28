#version 320

// Input
smooth in vec2 f_texcoord;
uniform sampler2D tex_id;
uniform int shape_mode;
uniform vec4 shape_color;

// Outputs
out vec4 fragColor;

void main() {

	if(shape_mode == 0) {
		fragColor = texture2D(tex_id, f_texcoord);
	} else {
		fragColor = shape_color;
	}
}
