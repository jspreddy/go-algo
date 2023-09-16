package src

import (
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth, windowHeight = 800, 600

func init() {
	// GLFW event handling must run on the main OS thread
	// runtime.LockOSThread()
}

func initWindow(title string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		log.With(err).Debug("failed to initialize glfw")
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	// glfw.WindowHint(glfw.ContextVersionMajor, 4)
	// glfw.WindowHint(glfw.ContextVersionMinor, 6)
	// glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	// glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Test Window", nil, nil)
	if err != nil {
		log.With(err).Error("Encountered an error while trying to create window.")
		panic(err)
	}
	window.MakeContextCurrent()

	return window
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Debug("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog

}

var (
	triangle = []float32{
		0, 0.5, 0, // top
		-0.5, -0.5, 0, // left
		0.5, -0.5, 0, // right
	}
)

func draw(vao uint32, window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	gl.BindVertexArray(vao)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(triangle)/3))

	glfw.PollEvents()
	window.SwapBuffers()
}

// makeVao initializes and returns a vertex array from the points provided.
func makeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func UI() {
	runtime.LockOSThread()

	window := initWindow("Test Window")
	defer glfw.Terminate()

	program := initOpenGL()

	vao := makeVao(triangle)
	for !window.ShouldClose() {
		draw(vao, window, program)
	}
}
