package paunch

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
)

type Window struct {
	glfwWindow *glfw.Window
}

func (window *Window) Open(width int, height int, title string) error {

	if !glfw.Init() {
		return errors.New("Could not initialize GLFW")
	}

	var err error
	window.glfwWindow, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}

	window.glfwWindow.MakeContextCurrent()

	return nil
}

func (window *Window) Destroy() error {

	glfw.Terminate()
	return nil
}

func (window *Window) ShouldClose() bool {

	return window.glfwWindow.ShouldClose()
}

func (window *Window) UpdateDisplay() error {

	window.glfwWindow.SwapBuffers()
	return nil
}

func (window *Window) UpdateEvents() error {

	glfw.PollEvents()
	return nil
}
