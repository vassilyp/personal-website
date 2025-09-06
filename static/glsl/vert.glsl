#version 300 es

in vec2 a_position;

out vec2 v_position;

uniform vec2 u_resolution;
uniform float u_time;

void main() {
    vec2 zeroToOne = a_position * u_resolution;
    vec2 zeroToTwo = zeroToOne * 2.0;
    vec2 clipSpace = zeroToTwo - 1.0;

    v_position = clipSpace;

    gl_Position = vec4(clipSpace, 0.0, 1.0);
}
