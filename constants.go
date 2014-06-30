package paunch

import (
	gl "github.com/chsc/gogl/gl32"
	glfw "github.com/go-gl/glfw3"
	al "github.com/vova616/go-openal/openal"
)

const tolerance = 0.01

// Axis corresponds to an axis value.
type Axis int

// Axis IDs
const (
	_ Axis = iota
	X
	Y
)

// Direction corresponds to a directional value.
type Direction int

// Direction IDs
const (
	_ Direction = iota
	Up
	Down
	Left
	Right
)

// Aliases for Direction IDs
const (
	Top    = Up
	Bottom = Down
)

// Action corresponds to a key or button action.
type Action int

// Action IDs
const (
	Press   = Action(glfw.Press)
	Release = Action(glfw.Release)
	Hold    = Action(glfw.Repeat)
)

// MouseButton corresponds to a mouse button.
type MouseButton int

// Mouse button IDs
const (
	MouseButton1      = MouseButton(glfw.MouseButton1)
	MouseButton2      = MouseButton(glfw.MouseButton2)
	MouseButton3      = MouseButton(glfw.MouseButton3)
	MouseButton4      = MouseButton(glfw.MouseButton4)
	MouseButton5      = MouseButton(glfw.MouseButton5)
	MouseButton6      = MouseButton(glfw.MouseButton6)
	MouseButton7      = MouseButton(glfw.MouseButton7)
	MouseButton8      = MouseButton(glfw.MouseButton8)
	MouseButtonLast   = MouseButton(glfw.MouseButtonLast)
	MouseButtonLeft   = MouseButton(glfw.MouseButtonLeft)
	MouseButtonRight  = MouseButton(glfw.MouseButtonRight)
	MouseButtonMiddle = MouseButton(glfw.MouseButtonMiddle)
)

// Key corresponds to a key on the keyboard.
type Key int

// Keyboard key IDs
const (
	KeyUnknown      = Key(glfw.KeyUnknown)
	KeySpace        = Key(glfw.KeySpace)
	KeyApostrophe   = Key(glfw.KeyApostrophe)
	KeyComma        = Key(glfw.KeyComma)
	KeyMinus        = Key(glfw.KeyMinus)
	KeyPeriod       = Key(glfw.KeyPeriod)
	KeySlash        = Key(glfw.KeySlash)
	Key0            = Key(glfw.Key0)
	Key1            = Key(glfw.Key1)
	Key2            = Key(glfw.Key2)
	Key3            = Key(glfw.Key3)
	Key4            = Key(glfw.Key4)
	Key5            = Key(glfw.Key5)
	Key6            = Key(glfw.Key6)
	Key7            = Key(glfw.Key7)
	Key8            = Key(glfw.Key8)
	Key9            = Key(glfw.Key9)
	KeySemicolon    = Key(glfw.KeySemicolon)
	KeyEqual        = Key(glfw.KeyEqual)
	KeyA            = Key(glfw.KeyA)
	KeyB            = Key(glfw.KeyB)
	KeyC            = Key(glfw.KeyC)
	KeyD            = Key(glfw.KeyD)
	KeyE            = Key(glfw.KeyE)
	KeyF            = Key(glfw.KeyF)
	KeyG            = Key(glfw.KeyG)
	KeyH            = Key(glfw.KeyH)
	KeyI            = Key(glfw.KeyI)
	KeyJ            = Key(glfw.KeyJ)
	KeyK            = Key(glfw.KeyK)
	KeyL            = Key(glfw.KeyL)
	KeyM            = Key(glfw.KeyM)
	KeyN            = Key(glfw.KeyN)
	KeyO            = Key(glfw.KeyO)
	KeyP            = Key(glfw.KeyP)
	KeyQ            = Key(glfw.KeyQ)
	KeyR            = Key(glfw.KeyR)
	KeyS            = Key(glfw.KeyS)
	KeyT            = Key(glfw.KeyT)
	KeyU            = Key(glfw.KeyU)
	KeyV            = Key(glfw.KeyV)
	KeyW            = Key(glfw.KeyW)
	KeyX            = Key(glfw.KeyX)
	KeyY            = Key(glfw.KeyY)
	KeyZ            = Key(glfw.KeyZ)
	KeyLeftBracket  = Key(glfw.KeyLeftBracket)
	KeyBackslash    = Key(glfw.KeyBackslash)
	KeyRightBracket = Key(glfw.KeyRightBracket)
	KeyGraveAccent  = Key(glfw.KeyGraveAccent)
	KeyWorld1       = Key(glfw.KeyWorld1)
	KeyWorld2       = Key(glfw.KeyWorld2)
	KeyEscape       = Key(glfw.KeyEscape)
	KeyEnter        = Key(glfw.KeyEnter)
	KeyTab          = Key(glfw.KeyTab)
	KeyBackspace    = Key(glfw.KeyBackspace)
	KeyInsert       = Key(glfw.KeyInsert)
	KeyDelete       = Key(glfw.KeyDelete)
	KeyRight        = Key(glfw.KeyRight)
	KeyLeft         = Key(glfw.KeyLeft)
	KeyDown         = Key(glfw.KeyDown)
	KeyUp           = Key(glfw.KeyUp)
	KeyPageUp       = Key(glfw.KeyPageUp)
	KeyPageDown     = Key(glfw.KeyPageDown)
	KeyHome         = Key(glfw.KeyHome)
	KeyEnd          = Key(glfw.KeyEnd)
	KeyCapsLock     = Key(glfw.KeyCapsLock)
	KeyScrollLock   = Key(glfw.KeyScrollLock)
	KeyNumLock      = Key(glfw.KeyNumLock)
	KeyPrKeyScreen  = Key(glfw.KeyPrintScreen)
	KeyPause        = Key(glfw.KeyPause)
	KeyF1           = Key(glfw.KeyF1)
	KeyF2           = Key(glfw.KeyF2)
	KeyF3           = Key(glfw.KeyF3)
	KeyF4           = Key(glfw.KeyF4)
	KeyF5           = Key(glfw.KeyF5)
	KeyF6           = Key(glfw.KeyF6)
	KeyF7           = Key(glfw.KeyF7)
	KeyF8           = Key(glfw.KeyF8)
	KeyF9           = Key(glfw.KeyF9)
	KeyF10          = Key(glfw.KeyF10)
	KeyF11          = Key(glfw.KeyF11)
	KeyF12          = Key(glfw.KeyF12)
	KeyF13          = Key(glfw.KeyF13)
	KeyF14          = Key(glfw.KeyF14)
	KeyF15          = Key(glfw.KeyF15)
	KeyF16          = Key(glfw.KeyF16)
	KeyF17          = Key(glfw.KeyF17)
	KeyF18          = Key(glfw.KeyF18)
	KeyF19          = Key(glfw.KeyF19)
	KeyF20          = Key(glfw.KeyF20)
	KeyF21          = Key(glfw.KeyF21)
	KeyF22          = Key(glfw.KeyF22)
	KeyF23          = Key(glfw.KeyF23)
	KeyF24          = Key(glfw.KeyF24)
	KeyF25          = Key(glfw.KeyF25)
	KeyKp0          = Key(glfw.KeyKp0)
	KeyKp1          = Key(glfw.KeyKp1)
	KeyKp2          = Key(glfw.KeyKp2)
	KeyKp3          = Key(glfw.KeyKp3)
	KeyKp4          = Key(glfw.KeyKp4)
	KeyKp5          = Key(glfw.KeyKp5)
	KeyKp6          = Key(glfw.KeyKp6)
	KeyKp7          = Key(glfw.KeyKp7)
	KeyKp8          = Key(glfw.KeyKp8)
	KeyKp9          = Key(glfw.KeyKp9)
	KeyKpDecimal    = Key(glfw.KeyKpDecimal)
	KeyKpDivide     = Key(glfw.KeyKpDivide)
	KeyKpMultiply   = Key(glfw.KeyKpMultiply)
	KeyKpSubtract   = Key(glfw.KeyKpSubtract)
	KeyKpAdd        = Key(glfw.KeyKpAdd)
	KeyKpEnter      = Key(glfw.KeyKpEnter)
	KeyKpEqual      = Key(glfw.KeyKpEqual)
	KeyLeftShift    = Key(glfw.KeyLeftShift)
	KeyLeftControl  = Key(glfw.KeyLeftControl)
	KeyLeftAlt      = Key(glfw.KeyLeftAlt)
	KeyLeftSuper    = Key(glfw.KeyLeftSuper)
	KeyRightShift   = Key(glfw.KeyRightShift)
	KeyRightControl = Key(glfw.KeyRightControl)
	KeyRightAlt     = Key(glfw.KeyRightAlt)
	KeyRightSuper   = Key(glfw.KeyRightSuper)
	KeyMenu         = Key(glfw.KeyMenu)
	KeyLast         = Key(glfw.KeyLast)
)

// SoundState is a value that can be retured by Sound.GetPlaying().
type SoundState int

// Sound state IDs
const (
	Initial = SoundState(al.Initial)
	Playing = SoundState(al.Playing)
	Paused  = SoundState(al.Paused)
	Stopped = SoundState(al.Stopped)
)

// ShaderType is a value that represents a shader type.
type ShaderType int

// Shader type IDs
const (
	Vertex    = ShaderType(gl.VERTEX_SHADER)
	Fragment  = ShaderType(gl.FRAGMENT_SHADER)
	notShader = -1
)

// Shape is a value that represents a shape type.
type Shape int

// Shape IDs
const (
	ShapePoints        = Shape(gl.POINTS)
	ShapeLines         = Shape(gl.LINES)
	ShapeLineStrip     = Shape(gl.LINE_STRIP)
	ShapeLineLoop      = Shape(gl.LINE_LOOP)
	ShapeTriangles     = Shape(gl.TRIANGLES)
	ShapeTriangleStrip = Shape(gl.TRIANGLE_STRIP)
	ShapeTriangleFan   = Shape(gl.TRIANGLE_FAN)
	ShapePolygon       = Shape(gl.POLYGON)
)

type Version int

// OpenGL Version Management
const (
	_ Version = iota
	VersionNew
	VersionOld
	VersionAutomatic
)
