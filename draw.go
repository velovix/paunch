package paunch

import (
	"errors"
	"fmt"
	gl "github.com/chsc/gogl/gl33"
	"image"
	"image/png"
	"os"
	"runtime"
)

const (
	POINTS         = gl.POINTS
	LINE_STRIP     = gl.LINE_STRIP
	LINE_LOOP      = gl.LINE_LOOP
	LINES          = gl.LINES
	TRIANGLE_STRIP = gl.TRIANGLE_STRIP
	TRIANGLE_FAN   = gl.TRIANGLE_FAN
	TRIANGLES      = gl.TRIANGLES
)

// Renderable is an object that can be drawn on the screen
type Renderable struct {
	mode            int
	size            int
	vertex_buffer   gl.Uint
	texcoord_buffer gl.Uint
	texture         gl.Uint
}

func checkForErrors() error {

	var errList []gl.Enum
	for err := gl.GetError(); err != gl.NO_ERROR; {
		errList = append(errList, err)
	}

	if len(errList) == 0 {
		return nil
	} else {
		return errors.New(fmt.Sprintln("OpenGL Error(s): ", errList))
	}
}

func imageToBytes(img image.Image) []byte {

	width := img.Bounds().Max.X - img.Bounds().Min.X
	height := img.Bounds().Max.Y - img.Bounds().Min.Y
	var bytes []byte

	for i := 0; i < width*height; i++ {
		r, g, b, a := img.At(i%width, i/width).RGBA()
		bytes = append(bytes, byte(r))
		bytes = append(bytes, byte(g))
		bytes = append(bytes, byte(b))
		bytes = append(bytes, byte(a))
	}

	flippedBytes := make([]byte, width*height*4)
	for i := 0; i < height*width*4; i += width * 4 {
		pixrow := bytes[i : i+width*4]
		for j, val := range pixrow {
			flippedBytes[((width*(height-1)*4)-i)+j] = val
		}
	}

	return flippedBytes
}

// InitDraw sets up the drawing session for use.
func InitDraw(window Window) error {

	runtime.LockOSThread()

	if err := gl.Init(); err != nil {
		return errors.New("Error initializing OpenGL")
	}

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	gl.Viewport(0, 0, gl.Sizei(window.Width), gl.Sizei(window.Height))

	return checkForErrors()
}

// NewRenderable returns a new Renderable object based on the specified shape
// type and verticies.
func NewRenderable(mode int, verticies []float32) (Renderable, error) {

	renderable := Renderable{mode, len(verticies), 0, 0, 0}

	gl.GenBuffers(1, &renderable.vertex_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.vertex_buffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return renderable, checkForErrors()
}

// .Texture applies a texture from a 32-bit PNG file to a Renderable. The
// Renderable will automatically be drawn with this texture.
func (renderable *Renderable) Texture(coords []float32, filename string) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gl.GenBuffers(1, &renderable.texcoord_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.texcoord_buffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(coords)*4), gl.Pointer(&coords[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	data, decodeErr := png.Decode(file)
	if decodeErr != nil {
		return decodeErr
	}

	byteData := imageToBytes(data)

	gl.GenTextures(1, &renderable.texture)
	gl.BindTexture(gl.TEXTURE_2D, renderable.texture)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
		gl.Sizei(data.Bounds().Max.X-data.Bounds().Min.X),
		gl.Sizei(data.Bounds().Max.Y-data.Bounds().Min.Y),
		0, gl.RGBA, gl.UNSIGNED_BYTE,
		gl.Pointer(&byteData[0]))
	gl.BindTexture(gl.TEXTURE_2D, 0)

	return checkForErrors()
}

// DrawRenderable draws a Renderable
func DrawRenderable(renderable Renderable) {

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertex_buffer)
	gl.VertexAttribPointer(gl.Uint(0), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if renderable.texcoord_buffer != 0 {
		gl.ActiveTexture(gl.TEXTURE0)

		gl.BindBuffer(gl.ARRAY_BUFFER, renderable.texcoord_buffer)
		gl.VertexAttribPointer(gl.Uint(1), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)

		gl.BindTexture(gl.TEXTURE_2D, renderable.texture)
		gl.EnableVertexAttribArray(gl.Uint(1))
	}

	gl.EnableVertexAttribArray(gl.Uint(0))
	gl.DrawArrays(gl.Enum(renderable.mode), 0, gl.Sizei(renderable.size))
	gl.DisableVertexAttribArray(gl.Uint(0))
	gl.DisableVertexAttribArray(gl.Uint(1))

	gl.BindTexture(gl.TEXTURE_2D, 0)
}

// Clear clears the pixels on screen. This should probably be called before
// every new frame.
func Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT)
}
