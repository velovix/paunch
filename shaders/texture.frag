#version 330

// Input
in vec2 f_texcoord;
uniform sampler2D tex_id;

// Outputs
out vec4 outputColor;

void main() {

	outputColor = texture2D(tex_id, f_texcoord);
}
