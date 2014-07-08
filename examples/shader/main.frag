#version 120

// Input
varying vec2 f_texcoord;
uniform sampler2D tex_id;
uniform int shape_mode;
uniform vec4 shape_color;

void main() {

	if(shape_mode == 0) {
		gl_FragColor = texture2D(tex_id, gl_TexCoord[0].st);
	} else {
		gl_FragColor = shape_color;
	}
}
