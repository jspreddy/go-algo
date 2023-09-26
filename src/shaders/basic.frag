#version 330 core

in vec4 vertexColor;
// out vec4 FragColor;
uniform vec2 u_resolution;  // Canvas size (width,height)
uniform vec2 u_mouse;       // mouse position in screen pixels
uniform float u_time;       // Time in seconds since load

const float PI = 3.141592653589793;

void main() {
    // gl_FragColor = vec4(vertexColor.xyz, 1.0);
    float rad = (PI * u_time);
    gl_FragColor = vec4(
        abs(sin(rad/2)),
        abs(sin(2*rad/3)),
        abs(sin(3*rad/5)),
        1.0
    );
}
