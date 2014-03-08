package paunch

import (
	"errors"
	gl "github.com/chsc/gogl/gl33"
	glfw "github.com/go-gl/glfw3"
	"math"
)

// Window is an object that manages window creation and user input.
type Window struct {
	width        int
	height       int
	nativeWidth  int
	nativeHeight int

	actorManager *ActorManager

	glfwWindow *glfw.Window

	keyStates      map[int]bool
	joyBtnStates   map[int]bool
	joyAxisStates  map[int]float32
	isJoystick     bool
	nativeMousePos bool
}

var glfwToWindow map[*glfw.Window]*Window

// Open opens a new window ready to be drawn in. The width and height will be
// the size of the window in pixels. The nativeWidth and nativeHeight represent
// the actual width and height of the drawing space if it were not stretched to
// accomidate the window size. This is only important when SetNativeMousePos is
// enabled.
func (window *Window) Open(width, height, nativeWidth, nativeHeight int, title string) error {

	if glfwToWindow == nil {
		glfwToWindow = make(map[*glfw.Window]*Window)
	}

	window.keyStates = make(map[int]bool)
	window.joyBtnStates = make(map[int]bool)
	window.joyAxisStates = make(map[int]float32)

	if !glfw.Init() {
		return errors.New("initializing GLFW")
	}

	var err error
	window.glfwWindow, err = glfw.CreateWindow(width, height, title, nil, nil)
	if err != nil {
		return err
	}

	window.width = width
	window.height = height
	window.nativeWidth = nativeWidth
	window.nativeHeight = nativeHeight

	window.glfwWindow.SetKeyCallback(keyboardCallback)
	window.glfwWindow.SetMouseButtonCallback(mouseButtonCallback)
	window.glfwWindow.SetCursorPositionCallback(mousePositionCallback)
	window.glfwWindow.SetCursorEnterCallback(mouseEnterWindowCallback)
	window.glfwWindow.SetFocusCallback(windowFocusCallback)
	window.glfwWindow.SetSizeCallback(windowResizeCallback)

	if glfw.JoystickPresent(glfw.Joystick1) {
		window.isJoystick = true
	}

	window.glfwWindow.MakeContextCurrent()

	glfwToWindow[window.glfwWindow] = window

	return nil
}

// Destroy closes the window and stops reading input.
func (window *Window) Destroy() error {

	glfw.Terminate()
	return nil
}

// GetSize returns the current width and height of the Window object.
func (window *Window) GetSize() (width, height int) {

	return window.width, window.height
}

// SetSize sets the width and height of the Window object and triggers a window
// resize event.
func (window *Window) SetSize(width, height int) {

	window.glfwWindow.SetSize(width, height)
}

// GetNativeSize returns the native width and height of the Window object.
func (window *Window) GetNativeSize() (nativeWidth, nativeHeight int) {

	return window.nativeWidth, window.nativeHeight
}

// SetNativeSize sets the native width and height of the Window object to the
// specified values.
func (window *Window) SetNativeSize(nativeWidth, nativeHeight int) {

	window.nativeWidth = nativeWidth
	window.nativeHeight = nativeHeight
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

	if window.actorManager == nil {
		return nil
	}

	for i, val := range window.keyStates {
		if val && window.actorManager != nil {
			window.actorManager.RunKeyEvent(Key(i), Hold)
		}
	}

	if window.isJoystick {
		buttons, err := glfw.GetJoystickButtons(glfw.Joystick1)
		if err != nil {
			panic(err)
		}

		for i, val := range buttons {
			if val == 0 && window.joyBtnStates[i] {
				window.actorManager.RunJoystickButtonEvent(i, Release)
				window.joyBtnStates[i] = false
			} else if val == 1 && !window.joyBtnStates[i] {
				window.actorManager.RunJoystickButtonEvent(i, Press)
				window.joyBtnStates[i] = true
			} else if val == 1 && window.joyBtnStates[i] {
				window.actorManager.RunJoystickButtonEvent(i, Hold)
			}
		}

		axes, err2 := glfw.GetJoystickAxes(glfw.Joystick1)
		if err2 != nil {
			panic(err2)
		}

		for i, val := range axes {
			window.actorManager.RunJoystickAxisEvent(i, float64(val))
		}
	}

	return nil
}

// SetActorManager sets the Actor Manager the Window object sends input events
// to.
func (window *Window) SetActorManager(actorManager *ActorManager) {

	window.actorManager = actorManager
}

// SetNativeMousePos changes the behavior of the reported mouse position. If
// enabled, supplied mouse positions are made relative to the native width and
// native height of the Window object. This is useful for applications that
// stretch to larger window sizes, so that mouse position behavior remains the
// same regardless of window size.
func (window *Window) SetNativeMousePos(shouldBeNative bool) {

	window.nativeMousePos = shouldBeNative
}

func keyboardCallback(window *glfw.Window, glfwKey glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if action == glfw.Repeat {
		return
	}

	glfwToWindow[window].keyStates[int(glfwKey)] = (action == glfw.Press)

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.RunKeyEvent(Key(glfwKey), Action(action))
	}
}

func mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

	if glfwToWindow[window].actorManager != nil {
		x, y := window.GetCursorPosition()

		var windHeight int
		if glfwToWindow[window].nativeMousePos {
			x *= float64(glfwToWindow[window].nativeWidth) / float64(glfwToWindow[window].width)
			y *= float64(glfwToWindow[window].nativeHeight) / float64(glfwToWindow[window].height)
			windHeight = glfwToWindow[window].nativeHeight
		} else {
			_, windHeight = window.GetSize()
		}
		glfwToWindow[window].actorManager.RunMouseButtonEvent(MouseButton(button), Action(action), int(math.Floor(x)), windHeight-int(math.Floor(y)))
	}
}

func mousePositionCallback(window *glfw.Window, x, y float64) {

	if glfwToWindow[window].actorManager != nil {
		var windHeight int
		if glfwToWindow[window].nativeMousePos {
			x *= float64(glfwToWindow[window].nativeWidth) / float64(glfwToWindow[window].width)
			y *= float64(glfwToWindow[window].nativeHeight) / float64(glfwToWindow[window].height)
			windHeight = glfwToWindow[window].nativeHeight
		} else {
			_, windHeight = window.GetSize()
		}
		glfwToWindow[window].actorManager.RunMousePositionEvent(int(math.Floor(x)), windHeight-int(math.Floor(y)))
	}
}

func mouseEnterWindowCallback(window *glfw.Window, entered bool) {

	if glfwToWindow[window].actorManager != nil {
		x, y := window.GetCursorPosition()

		var windHeight int
		if glfwToWindow[window].nativeMousePos {
			x *= float64(glfwToWindow[window].nativeWidth) / float64(glfwToWindow[window].width)
			y *= float64(glfwToWindow[window].nativeHeight) / float64(glfwToWindow[window].height)
			windHeight = glfwToWindow[window].nativeHeight
		} else {
			_, windHeight = window.GetSize()
		}
		glfwToWindow[window].actorManager.RunMouseEnterWindowEvent(int(math.Floor(x)), windHeight-int(math.Floor(y)), entered)
	}
}

func windowFocusCallback(window *glfw.Window, focused bool) {

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.RunWindowFocusEvent(focused)
	}
}

func windowResizeCallback(window *glfw.Window, width, height int) {

	glfwToWindow[window].width = width
	glfwToWindow[window].height = height

	if glfwToWindow[window].actorManager != nil {
		glfwToWindow[window].actorManager.RunWindowResizeEvent(width, height)
	}

	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}
