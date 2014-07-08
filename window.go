package paunch

import (
	"errors"
	glfw "github.com/go-gl/glfw3"
	"github.com/velovix/gl"
)

type _Window struct {
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

var paunchWindow _Window

// SetWindowSize sets the width and height of the Paunch window. Width and
// height MUST be set before calling Start(). Any number below zero is not a
// valid width or height value.
func SetWindowSize(width, height int) {

	if paunchWindow.glfwWindow == nil {
		paunchWindow.width = width
		paunchWindow.height = height
		if paunchWindow.nativeWidth == 0 && paunchWindow.nativeHeight == 0 {
			paunchWindow.nativeWidth = width
			paunchWindow.nativeHeight = height
		}
	} else {
		paunchWindow.glfwWindow.SetSize(width, height)
	}
}

// SetNativeWindowSize sets the native width and height of the Paunch window.
// The native width and height are values representing what size the window
// must be at for one screen pixel to equal one pixel in your program. If these
// values are not set, Paunch will make these values equal to the size of the
// window at the time Start() is called.
func SetNativeWindowSize(nativeWidth, nativeHeight int) {

	paunchWindow.nativeWidth = nativeWidth
	paunchWindow.nativeHeight = nativeHeight
}

// SetWindowTitle sets the title of the Paunch window, which is generally the
// text that's shown at the top of the window. The default string is empty.
func SetWindowTitle(title string) {

	paunchWindow.title = title
}

// SetWindowFullScreen sets whether or not the Paunch window is full screen.
// The window will not be full screen by default.
func SetWindowFullScreen(fullscreen bool) {

	paunchWindow.fullscreen = fullscreen
}

// SetWindowNativeMousePos changes the behavior of the reported mouse position.
// If enabled, supplied mouse positions are made relative to the native width
// and native height of the Window object. This is useful for applications that
// stretch to larger window sizes, so that mouse position behavior remains the
// same regardless of window size.
func SetWindowNativeMousePos(shouldBeNative bool) {

	paunchWindow.nativeMousePos = shouldBeNative
}

// ShouldClose returns whether or not the user has attempted to close the
// window. Will always return false if the Window object has not been opened.
func ShouldClose() bool {

	if !paunchWindow.isOpen {
		return false
	}

	return paunchWindow.glfwWindow.ShouldClose()
}

// UpdateDisplay updates the window to display whatever has been drawn to the
// framebuffer.
func UpdateDisplay() error {

	if !paunchWindow.isOpen {
		return errors.New("window has not been opened")
	}

	paunchWindow.glfwWindow.SwapBuffers()

	return nil
}

// UpdateEvents updates events.
func UpdateEvents() error {

	if !paunchWindow.isOpen {
		return errors.New("window has not been opened")
	}

	glfw.PollEvents()

	if len(paunchWindow.eventManagers) == 0 {
		return nil
	}

	for i, val := range paunchWindow.keyStates {
		for _, eventManager := range paunchWindow.eventManagers {
			if val {
				eventManager.RunKeyEvent(Key(i), Hold)
			}
		}
	}

	if paunchWindow.isJoystick {
		buttons, err := glfw.GetJoystickButtons(glfw.Joystick1)
		if err != nil {
			return err
		}

		for _, eventManager := range paunchWindow.eventManagers {
			for i, val := range buttons {
				if val == 0 && paunchWindow.joyBtnStates[i] {
					eventManager.RunJoystickButtonEvent(i, Release)
					paunchWindow.joyBtnStates[i] = false
				} else if val == 1 && !paunchWindow.joyBtnStates[i] {
					eventManager.RunJoystickButtonEvent(i, Press)
					paunchWindow.joyBtnStates[i] = true
				} else if val == 1 && paunchWindow.joyBtnStates[i] {
					eventManager.RunJoystickButtonEvent(i, Hold)
				}
			}
		}

		axes, err2 := glfw.GetJoystickAxes(glfw.Joystick1)
		if err2 != nil {
			return err2
		}

		for _, eventManager := range paunchWindow.eventManagers {
			for i, val := range axes {
				eventManager.RunJoystickAxisEvent(i, float64(val))
			}
		}
	}

	return nil
}

func initWindows() error {

	if !glfw.Init() {
		return errors.New("initializing GLFW")
	}

	return nil
}

func newWindow() _Window {

	var window _Window

	window.eventManagers = make([]*EventManager, 0)

	window.keyStates = make(map[int]bool)
	window.joyBtnStates = make(map[int]bool)
	window.joyAxisStates = make(map[int]float32)

	return window
}

func (window *_Window) open() error {

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

	window.glfwWindow.SetKeyCallback(keyboardCallback)
	window.glfwWindow.SetMouseButtonCallback(mouseButtonCallback)
	window.glfwWindow.SetCursorPositionCallback(mousePositionCallback)
	window.glfwWindow.SetCursorEnterCallback(mouseEnterWindowCallback)
	window.glfwWindow.SetFocusCallback(windowFocusCallback)
	window.glfwWindow.SetSizeCallback(windowResizeCallback)
	window.glfwWindow.SetCharacterCallback(windowCharacterCallback)

	if glfw.JoystickPresent(glfw.Joystick1) {
		window.isJoystick = true
	}

	window.glfwWindow.MakeContextCurrent()

	window.isOpen = true

	return nil
}

func (window *_Window) close() error {

	if !window.isOpen {
		return errors.New("window has not been opened")
	}

	window.glfwWindow.Destroy()
	window.isOpen = false

	return nil
}

func keyboardCallback(window *glfw.Window, glfwKey glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if action == glfw.Repeat {
		return
	}

	paunchWindow.keyStates[int(glfwKey)] = (action == glfw.Press)

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunKeyEvent(Key(glfwKey), Action(action))
	}
}

func mouseButtonCallback(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {

	x, y := window.GetCursorPosition()

	var windHeight int
	if paunchWindow.nativeMousePos {
		x *= float64(paunchWindow.nativeWidth) / float64(paunchWindow.width)
		y *= float64(paunchWindow.nativeHeight) / float64(paunchWindow.height)
		windHeight = paunchWindow.nativeHeight
	} else {
		_, windHeight = window.GetSize()
	}

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunMouseButtonEvent(MouseButton(button), Action(action), x, float64(windHeight)-y)
	}
}

func mousePositionCallback(window *glfw.Window, x, y float64) {

	var windHeight int
	if paunchWindow.nativeMousePos {
		x *= float64(paunchWindow.nativeWidth) / float64(paunchWindow.width)
		y *= float64(paunchWindow.nativeHeight) / float64(paunchWindow.height)
		windHeight = paunchWindow.nativeHeight
	} else {
		_, windHeight = window.GetSize()
	}

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunMousePositionEvent(x, float64(windHeight)-y)
	}
}

func mouseEnterWindowCallback(window *glfw.Window, entered bool) {

	x, y := window.GetCursorPosition()

	var windHeight int
	if paunchWindow.nativeMousePos {
		x *= float64(paunchWindow.nativeWidth) / float64(paunchWindow.width)
		y *= float64(paunchWindow.nativeHeight) / float64(paunchWindow.height)
		windHeight = paunchWindow.nativeHeight
	} else {
		_, windHeight = window.GetSize()
	}

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunMouseEnterWindowEvent(x, float64(windHeight)-y, entered)
	}
}

func windowFocusCallback(window *glfw.Window, focused bool) {

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunWindowFocusEvent(focused)
	}
}

func windowResizeCallback(window *glfw.Window, width, height int) {

	paunchWindow.width = width
	paunchWindow.height = height

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunWindowResizeEvent(width, height)
	}

	gl.Viewport(0, 0, gl.Sizei(width), gl.Sizei(height))
}

func windowCharacterCallback(window *glfw.Window, character uint) {

	for _, eventManager := range paunchWindow.eventManagers {
		eventManager.RunCharacterEvent(rune(character))
	}
}
