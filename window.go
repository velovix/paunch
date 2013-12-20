package paunch

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
)

// Window is an object that manages window creation and user input.
type Window struct {
	Width  int
	Height int

	actorManager *ActorManager

	glfwWindow *glfw.Window
}

var glfwToWindow map[*glfw.Window]*Window

// Open opens a new window ready to be drawn in.
func (window *Window) Open(width int, height int, title string) error {

	if glfwToWindow == nil {
		glfwToWindow = make(map[*glfw.Window]*Window)
	}

	if !glfw.Init() {
		return errors.New("initializing GLFW")
	}

	var err error
	window.glfwWindow, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}

	window.Width = width
	window.Height = height

	window.glfwWindow.SetKeyCallback(keyboardCallback)

	window.glfwWindow.MakeContextCurrent()

	glfwToWindow[window.glfwWindow] = window

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

// SetActorManager sets the Actor Manager the Window object sends input events
// to.
func (window *Window) SetActorManager(actorManager *ActorManager) {

	window.actorManager = actorManager
}

func keyboardCallback(window *glfw.Window, glfwKey glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.keyEvent(int(glfwKey), int(action))
	}
}
