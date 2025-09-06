#version 300 es

precision highp float;

in vec2 v_position;
uniform vec2 u_resolution;
uniform float u_time;

out vec4 outColor;

float random(in vec2 st) {
    return fract(sin(dot(st.xy, vec2(12.9898, 78.233))) * 43758.5453123);
}

// Based on Morgan McGuire @morgan3d
// https://www.shadertoy.com/view/4dS3Wd
float noise(in vec2 st) {
    vec2 i = floor(st);
    vec2 f = fract(st);

    // Four corners in 2D of a tile
    float a = random(i);
    float b = random(i + vec2(1.0, 0.0));
    float c = random(i + vec2(0.0, 1.0));
    float d = random(i + vec2(1.0, 1.0));

    vec2 u = f * f * (3.0 - 2.0 * f);

    return mix(a, b, u.x) +
        (c - a) * u.y * (1.0 - u.x) +
        (d - b) * u.x * u.y;
}

// Based on the book of shaders: https://thebookofshaders.com/13/
#define OCTAVES 8
float fbm(in vec2 st) {
    // Initial values
    float value = 0.0;
    float amplitude = 0.2;

    // Loop of octaves
    for (int i = 0; i < OCTAVES; i++) {
        value += amplitude * noise(st);
        st *= 2.;
        amplitude *= .5;
    }
    return value;
}

void main() {
    vec2 st = v_position;

    // Base color: dark blue
    vec3 color = vec3(0.2627, 0.349, 0.4627);

    // Make the clouds go up with time
    st.y += u_time * 0.05;

    // Create clouds by adding fractal brownian motion noise
    color += fbm(st * 5.);

    // Create a "sub-layer" of clouds underneath
    color -= fbm(st - vec2(u_time * 0.15));
    color += fbm(st + vec2(u_time * 0.01));

    outColor = vec4(color, 1.);
}
