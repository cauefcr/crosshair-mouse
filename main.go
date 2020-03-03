package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.6-compatibility/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func main() {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	pm := glfw.GetPrimaryMonitor()
	vm := pm.GetVideoMode()
	// glfw.WindowHint(glfw.RedBits, vm.RedBits)
	// glfw.WindowHint(glfw.GreenBits, vm.GreenBits)
	// glfw.WindowHint(glfw.BlueBits, vm.BlueBits)
	glfw.WindowHint(glfw.TransparentFramebuffer, glfw.True)
	glfw.WindowHint(glfw.Decorated, glfw.False)

	// if it's the same size of the window it doesn't work (at least on windows), idk why
	window, err := glfw.CreateWindow(vm.Width-1, vm.Height-1, "much wow", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	window.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)

	if err := gl.Init(); err != nil {
		panic(err)
	}
	prog := gl.CreateProgram()
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
		gl.UseProgram(prog)
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
