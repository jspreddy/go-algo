package src

import (
	_ "embed"
	"fmt"
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth, windowHeight = 800, 600
const FLOAT_SIZE_MULTIPLIER = 4

//go:embed shaders/basic.vert
var basicVertCode []byte

//go:embed shaders/basic.frag
var basicFragCode []byte

func getSizeInBytes(points []float32) int {
	sizeBytes := FLOAT_SIZE_MULTIPLIER * len(points)
	return sizeBytes
}

func initWindow(title string) *glfw.Window {
	if err := glfw.Init(); err != nil {
		log.With(err).Debug("failed to initialize glfw")
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	// glfw.WindowHint(glfw.ContextVersionMajor, 3)
	// glfw.WindowHint(glfw.ContextVersionMinor, 3)
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

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		logString := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(logString))

		log.Error("Shader Compile Error.", "string", logString)

		return 0, fmt.Errorf("failed to compile %v: %v", source, logString)
	}

	return shader, nil
}

func createRenderPipeline(window *glfw.Window) uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Info("OpenGL version", "version", version)

	prog := gl.CreateProgram()
	log.Info("Glx Context", "major", window.GetAttrib(glfw.ContextVersionMajor))
	log.Info("Glx Context", "minor", window.GetAttrib(glfw.ContextVersionMajor))
	log.Info("Glx Context", "profile", window.GetAttrib(glfw.OpenGLProfile))
	log.Info("Glx Context", "compat", window.GetAttrib(glfw.OpenGLForwardCompatible))

	vertShader, _ := compileShader(string(basicVertCode), gl.VERTEX_SHADER)
	fragShader, _ := compileShader(string(basicFragCode), gl.FRAGMENT_SHADER)

	gl.AttachShader(prog, vertShader)
	gl.AttachShader(prog, fragShader)
	gl.LinkProgram(prog)
	return prog
}

// x, y, z, r, g, b,
var (
	triangle = []float32{
		+0, +1, +0, +1, +0, +0,
		+1, -1, +0, +0, +1, +0,
		-1, -1, +0, +0, +0, +1,
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
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, nil)

	// describe color attribute
	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointerWithOffset(1, 3, gl.FLOAT, false, 6*4, 3*4)
}

func drawLoop(window *glfw.Window, programId uint32) {
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

	programId := createRenderPipeline(window)
	sendDataToOpenGL(triangle)

	for !window.ShouldClose() {
		drawLoop(window, programId)
	}
}
