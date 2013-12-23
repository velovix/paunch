package paunch

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
	"math"
)

// Window is an object that manages window creation and user input.
type Window struct {
	Width  int
	Height int

	actorManager *ActorManager

	glfwWindow *glfw.Window

	keyStates [512]bool
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
	window.glfwWindow.SetMouseButtonCallback(mouseButtonCallback)
	window.glfwWindow.SetCursorPositionCallback(mousePositionCallback)
	window.glfwWindow.SetCursorEnterCallback(mouseEnterWindowCallback)

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

	for i, val := range window.keyStates {
		if val && window.actorManager != nil {
			window.actorManager.keyEvent(i, Hold)
		}
	}

	glfw.PollEvents()
	return nil
}

// SetActorManager sets the Actor Manager the Window object sends input events
// to.
func (window *Window) SetActorManager(actorManager *ActorManager) {

	window.actorManager = actorManager
}

func keyboardCallback(window *glfw.Window, glfwKey glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if action == glfw.Repeat {
		return
	}

	glfwToWindow[window].keyStates[int(glfwKey)] = (action == glfw.Press)

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.keyEvent(int(glfwKey), int(action))
	}
}

func mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

	if glfwToWindow[window].actorManager != nil {
		x, y := window.GetCursorPosition()

		glfwToWindow[window].actorManager.mouseButtonEvent(int(button), int(action), int(math.Floor(x)), int(math.Floor(y)))
	}
}

func mousePositionCallback(window *glfw.Window, x, y float64) {

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.mousePositionEvent(int(math.Floor(x)), int(math.Floor(y)))
	}
}

func mouseEnterWindowCallback(window *glfw.Window, entered bool) {

	if glfwToWindow[window].actorManager != nil {
		x, y := window.GetCursorPosition()

		glfwToWindow[window].actorManager.mouseEnterWindowEvent(int(math.Floor(x)), int(math.Floor(y)), entered)
	}
}

func windowFocusCallback(window *glfw.Window, focused bool) {

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.windowFocusEvent(focused)
	}
}
