#version 330

// Input
smooth in vec2 f_texcoord;
uniform sampler2D tex_id;

// Outputs
out vec4 fragColor;

void main() {

	fragColor = texture2D(tex_id, f_texcoord);
}
