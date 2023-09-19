package src

import (
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth, windowHeight = 800, 600

func getSizeInBytes(points []float32) int {
	FLOAT_SIZE_MULTIPLIER := 4
	sizeBytes := FLOAT_SIZE_MULTIPLIER * len(points)
	return sizeBytes
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

func createRenderPipeline() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Info("OpenGL version", "version", version)

	prog := gl.CreateProgram()

	// TODO: create vert shader; compile; add to prog;
	// vertShader := gl.CreateShader(gl.VERTEX_SHADER)
	// gl.ShaderSource(vertShader, 1, &vertShaderString, len(vertShaderString))

	// TODO: create frag shader; compile; add to prog;
	// fragShader := gl.CreateShader(gl.FRAGMENT_SHADER)

	gl.LinkProgram(prog)
	return prog
}

var (
	triangle = []float32{
		+0, +1, +0,
		+1, -1, +0,
		-1, -1, +0,
	}
)

func sendDataToOpenGL(points []float32) {
	sizeBytes := getSizeInBytes(points)

	var vboId uint32

	// make space
	gl.GenBuffers(1, &vboId)

	// assign that space to ARRAY_BUFFER
	gl.BindBuffer(gl.ARRAY_BUFFER, vboId)

	// fill the ARRAY_BUFFER with data, which in turn points to vbo above because we bound ARRAY_BUFFER to vbo
	gl.BufferData(gl.ARRAY_BUFFER, sizeBytes, gl.Ptr(points), gl.STATIC_DRAW)

	// enable and tell opengl that the first attribute in array is the vertex and define the vertex
	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}

func draw(window *glfw.Window, programId uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(programId)

	gl.DrawArrays(gl.TRIANGLES, 0, 3)

	// should probably use a callback for window size instead of re-calc every frame.
	w, h := window.GetSize()
	gl.Viewport(0, 0, int32(w), int32(h))

	glfw.PollEvents()
	window.SwapBuffers()
}

func UI() {
	// step 1: lock gl runtime to this thread
	runtime.LockOSThread()
	// step 2: initialize a window
	window := initWindow("Test Window")
	// defer the glfw to be cleaned up at the end
	defer glfw.Terminate()

	programId := createRenderPipeline()
	sendDataToOpenGL(triangle)

	for !window.ShouldClose() {
		draw(window, programId)
	}
}
