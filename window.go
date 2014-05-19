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
	title        string

	eventManagers []*EventManager

	glfwWindow *glfw.Window

	keyStates      map[int]bool
	joyBtnStates   map[int]bool
	joyAxisStates  map[int]float32
	isOpen         bool
	isJoystick     bool
	nativeMousePos bool
	fullscreen     bool
}

var glfwToWindow map[*glfw.Window]*Window

// InitWindows initializes the Window object system. This must be called before
// other window-based operations may take place.
func InitWindows() error {

	if !glfw.Init() {
		return errors.New("initializing GLFW")
	}

	return nil
}

// NewWindow creates a new Window object. The width and height will be
// the size of the window in pixels. The nativeWidth and nativeHeight represent
// the actual width and height of the drawing space if it were not stretched to
// accomidate the window size. This is only important when SetNativeMousePos is
// enabled.
func NewWindow(width, height, nativeWidth, nativeHeight int, fullscreen bool, title string) Window {

	var window Window

	if glfwToWindow == nil {
		glfwToWindow = make(map[*glfw.Window]*Window)
	}

	window.eventManagers = make([]*EventManager, 0)

	window.keyStates = make(map[int]bool)
	window.joyBtnStates = make(map[int]bool)
	window.joyAxisStates = make(map[int]float32)

	window.width = width
	window.height = height
	window.nativeWidth = nativeWidth
	window.nativeHeight = nativeHeight

	window.fullscreen = fullscreen
	window.title = title

	return window
}

// Open opens the Window object.
func (window *Window) Open() error {

	var err error
	if window.fullscreen {
		primaryMonitor, monitorErr := glfw.GetPrimaryMonitor()
		if monitorErr != nil {
			return monitorErr
		}
		window.glfwWindow, err = glfw.CreateWindow(window.width, window.height, window.title, primaryMonitor, nil)
	} else {
		window.glfwWindow, err = glfw.CreateWindow(window.width, window.height, window.title, nil, nil)
	}
	if err != nil {
		return err
	}

	glfwToWindow[window.glfwWindow] = window

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

	window.isOpen = true

	return nil
}

// Close closes the window and stops reading input.
func (window *Window) Close() error {

	if !window.isOpen {
		return errors.New("window has not been opened")
	}

	window.glfwWindow.Destroy()
	window.isOpen = false

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
// window. Will always return false if the Window object has not been opened.
func (window *Window) ShouldClose() bool {

	if !window.isOpen {
		return false
	}

	return window.glfwWindow.ShouldClose()
}

// UpdateDisplay updates the window to display whatever has been drawn to the
// framebuffer.
func (window *Window) UpdateDisplay() error {

	if !window.isOpen {
		return errors.New("window has not been opened")
	}

	window.glfwWindow.SwapBuffers()

	return nil
}

// UpdateEvents updates events.
func (window *Window) UpdateEvents() error {

	if !window.isOpen {
		return errors.New("window has not been opened")
	}

	glfw.PollEvents()

	if len(window.eventManagers) == 0 {
		return nil
	}

	for i, val := range window.keyStates {
		for _, eventManager := range window.eventManagers {
			if val {
				eventManager.RunKeyEvent(Key(i), Hold)
			}
		}
	}

	if window.isJoystick {
		buttons, err := glfw.GetJoystickButtons(glfw.Joystick1)
		if err != nil {
			return err
		}

		for _, eventManager := range window.eventManagers {
			for i, val := range buttons {
				if val == 0 && window.joyBtnStates[i] {
					eventManager.RunJoystickButtonEvent(i, Release)
					window.joyBtnStates[i] = false
				} else if val == 1 && !window.joyBtnStates[i] {
					eventManager.RunJoystickButtonEvent(i, Press)
					window.joyBtnStates[i] = true
				} else if val == 1 && window.joyBtnStates[i] {
					eventManager.RunJoystickButtonEvent(i, Hold)
				}
			}
		}

		axes, err2 := glfw.GetJoystickAxes(glfw.Joystick1)
		if err2 != nil {
			return err2
		}

		for _, eventManager := range window.eventManagers {
			for i, val := range axes {
				eventManager.RunJoystickAxisEvent(i, float64(val))
			}
		}
	}

	return nil
}

// SetEventManagers sets the EventManagers the Window object sends input events
// to. The Window object uses the copy of the slice given to it, so future
// modifications to that slice will affect the Window object, as well.
func (window *Window) SetEventManagers(eventManagers []*EventManager) {

	window.eventManagers = eventManagers
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

	for _, eventManager := range glfwToWindow[window].eventManagers {
		eventManager.RunKeyEvent(Key(glfwKey), Action(action))
	}
}

func mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

	x, y := window.GetCursorPosition()

	var windHeight int
	if glfwToWindow[window].nativeMousePos {
		x *= float64(glfwToWindow[window].nativeWidth) / float64(glfwToWindow[window].width)
		y *= float64(glfwToWindow[window].nativeHeight) / float64(glfwToWindow[window].height)
		windHeight = glfwToWindow[window].nativeHeight
	} else {
		_, windHeight = window.GetSize()
	}

	for _, eventManager := range glfwToWindow[window].eventManagers {
		eventManager.RunMouseButtonEvent(MouseButton(button), Action(action), int(math.Floor(x)), windHeight-int(math.Floor(y)))
	}
}

func mousePositionCallback(window *glfw.Window, x, y float64) {

	var windHeight int
	if glfwToWindow[window].nativeMousePos {
		x *= float64(glfwToWindow[window].nativeWidth) / float64(glfwToWindow[window].width)
		y *= float64(glfwToWindow[window].nativeHeight) / float64(glfwToWindow[window].height)
		windHeight = glfwToWindow[window].nativeHeight
	} else {
		_, windHeight = window.GetSize()
	}

	for _, eventManager := range glfwToWindow[window].eventManagers {
		eventManager.RunMousePositionEvent(int(math.Floor(x)), windHeight-int(math.Floor(y)))
	}
}

func mouseEnterWindowCallback(window *glfw.Window, entered bool) {

	x, y := window.GetCursorPosition()

	var windHeight int
	if glfwToWindow[window].nativeMousePos {
		x *= float64(glfwToWindow[window].nativeWidth) / float64(glfwToWindow[window].width)
		y *= float64(glfwToWindow[window].nativeHeight) / float64(glfwToWindow[window].height)
		windHeight = glfwToWindow[window].nativeHeight
	} else {
		_, windHeight = window.GetSize()
	}

	for _, eventManager := range glfwToWindow[window].eventManagers {
		eventManager.RunMouseEnterWindowEvent(int(math.Floor(x)), windHeight-int(math.Floor(y)), entered)
	}
}

func windowFocusCallback(window *glfw.Window, focused bool) {

	for _, eventManager := range glfwToWindow[window].eventManagers {
		eventManager.RunWindowFocusEvent(focused)
	}
}

func windowResizeCallback(window *glfw.Window, width, height int) {

	glfwToWindow[window].width = width
	glfwToWindow[window].height = height

	for _, eventManager := range glfwToWindow[window].eventManagers {
		eventManager.RunWindowResizeEvent(width, height)
	}

	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}
