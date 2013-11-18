package paunch

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
)

// Window is an object that manages window creation and user input.
type Window struct {
	Width      int
	Height     int
	glfwWindow *glfw.Window
}

// Open opens a new window ready to be drawn in.
func (window *Window) Open(width int, height int, title string) error {

	if !glfw.Init() {
		return errors.New("Could not initialize GLFW")
	}

	var err error
	window.glfwWindow, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}

	window.Width = width
	window.Height = height

	window.glfwWindow.MakeContextCurrent()

	return nil
}

// Destroy closes the window and stops reading input.
func (window *Window) Destroy() error {

	glfw.Terminate()
	return nil
}

// ShouldClose returns whether or not the user has attempted to close the
// window.
func (window *Window) ShouldClose() bool {

	return window.glfwWindow.ShouldClose()
}

// UpdateDisplay updates the window to display whatever has been drawn to the
// framebuffer.
func (window *Window) UpdateDisplay() error {

	window.glfwWindow.SwapBuffers()
	return nil
}

// UpdateEvents updates events.
func (window *Window) UpdateEvents() error {

	glfw.PollEvents()
	return nil
}
