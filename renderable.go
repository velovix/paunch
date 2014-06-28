package paunch

import (
	gl "github.com/chsc/gogl/gl32"
	"image"
	"image/png"
	"os"
)

// Renderable is an object that can be drawn on the screen
type Renderable struct {
	mode           int
	size           int
	vertexBuffer   gl.Uint
	texcoordBuffer gl.Uint
	texture        []gl.Uint
	verticies      []float32
}

func imageToBytes(img image.Image) (int, int, []byte) {

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

	return img.Bounds().Max.X, img.Bounds().Max.Y, flippedBytes
}

// NewRenderableFromShape creates a new Renderable object using the given shape
// type and verticies. The shape is formed in the same way OpenGL shapes are
// made.
func NewRenderableFromShape(shape Shape, verticies []float64) (*Renderable, error) {

	verticies32 := make([]float32, len(verticies))
	for i, val := range verticies {
		verticies32[i] = float32(val)
	}

	renderable := &Renderable{mode: int(shape), size: len(verticies), vertexBuffer: 0, texcoordBuffer: 0, texture: nil, verticies: verticies32}

	gl.GenBuffers(1, &renderable.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.vertexBuffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies32[0]), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return renderable, checkForErrors()
}

// NewRenderableFromData creates a new Renderable object using the given data,
// which is expected to be in RGBA format.
func NewRenderableFromData(x, y, width, height float64, data []byte, clip int) (*Renderable, error) {

	var renderable *Renderable

	verticies := []float32{
		float32(x), float32(y),
		float32(x + width), float32(y),
		float32(x), float32(y + height),

		float32(x + width), float32(y + height),
		float32(x + width), float32(y),
		float32(x), float32(y + height)}

	renderable = &Renderable{int(ShapeTriangles), len(verticies), 0, 0, nil, verticies}

	gl.GenBuffers(1, &renderable.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.vertexBuffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	texCoords := []float32{
		0, 0,
		1, 0,
		0, 1,

		1, 1,
		1, 0,
		0, 1}

	gl.GenBuffers(1, &renderable.texcoordBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.texcoordBuffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(texCoords)*4), gl.Pointer(&texCoords[0]), gl.STREAM_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	renderable.texture = make([]gl.Uint, clip)
	gl.GenTextures(gl.Sizei(clip), &renderable.texture[0])

	clips := make([][]byte, clip)
	for i := range clips {
		clips[i] = data[i*(len(data)/len(clips)) : (i+1)*(len(data)/len(clips))]
		gl.BindTexture(gl.TEXTURE_2D, renderable.texture[len(clips)-1-i])
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
			gl.Sizei(width), gl.Sizei(height),
			0, gl.RGBA, gl.UNSIGNED_BYTE,
			gl.Pointer(&clips[i][0]))
	}

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return renderable, checkForErrors()
}

// NewRenderableFromImage creates a new Renderable object using the given PNG
// image file.
func NewRenderableFromImage(x, y float64, filename string, clip int) (*Renderable, error) {

	var renderable *Renderable

	file, err := os.Open(filename)
	if err != nil {
		return renderable, err
	}
	defer file.Close()

	data, err := png.Decode(file)
	if err != nil {
		return renderable, err
	}

	width, height, byteData := imageToBytes(data)
	renderable, err = NewRenderableFromData(x, y, float64(width), float64(height/clip), byteData, clip)
	if err != nil {
		return renderable, err
	}

	return renderable, nil
}

// NewRenderableFromRenderable creates a new Renderable object that uses the
// same image data as the supplied Renderable object. This can serve to save a
// lot of GPU memory when dealing with Renderable objects that use image data.
// Renderable objects made from shapes will not benefit from being created
// this way from a performance point of view.
func NewRenderableFromRenderable(renderable *Renderable) *Renderable {

	newRenderable := &Renderable{mode: renderable.mode, size: renderable.size, texcoordBuffer: renderable.texcoordBuffer,
		texture: renderable.texture, verticies: make([]float32, len(renderable.verticies))}

	copy(newRenderable.verticies, renderable.verticies)

	gl.GenBuffers(1, &newRenderable.vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(newRenderable.vertexBuffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(newRenderable.verticies)*4), gl.Pointer(&newRenderable.verticies[0]), gl.DYNAMIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return newRenderable
}

func (renderable *Renderable) SetScaling(xScale, yScale float64) {

	verticies := make([]float32, len(renderable.verticies))

	xTransform := renderable.verticies[0] - (renderable.verticies[0] * float32(xScale))
	yTransform := renderable.verticies[1] - (renderable.verticies[1] * float32(yScale))

	for i := range verticies {
		if i%2 == 0 {
			verticies[i] = renderable.verticies[i] * float32(xScale)
			verticies[i] += xTransform
		} else {
			verticies[i] = renderable.verticies[i] * float32(yScale)
			verticies[i] += yTransform
		}
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertexBuffer)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// Draw draws the Renderable.
func (renderable *Renderable) Draw(frame int) error {

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertexBuffer)
	gl.VertexAttribPointer(gl.Uint(0), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if renderable.texcoordBuffer != 0 {
		gl.ActiveTexture(gl.TEXTURE0)

		gl.BindBuffer(gl.ARRAY_BUFFER, renderable.texcoordBuffer)
		gl.VertexAttribPointer(gl.Uint(1), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)

		gl.BindTexture(gl.TEXTURE_2D, renderable.texture[frame])
		gl.EnableVertexAttribArray(gl.Uint(1))
	}

	gl.EnableVertexAttribArray(gl.Uint(0))
	gl.DrawArrays(gl.Enum(renderable.mode), 0, gl.Sizei(renderable.size))
	gl.DisableVertexAttribArray(gl.Uint(0))
	gl.DisableVertexAttribArray(gl.Uint(1))

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return checkForErrors()
}

// Move moves the Renderable a specified distance.
func (renderable *Renderable) Move(x, y float64) {

	for i := 0; i < len(renderable.verticies); i += 2 {
		renderable.verticies[i] += float32(x)
		renderable.verticies[i+1] += float32(y)
	}

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertexBuffer)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, gl.Sizeiptr(len(renderable.verticies)*4), gl.Pointer(&renderable.verticies[0]))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}

// SetPosition sets the position of the Renderable object with the first
// defined vertex as the start point. For Renderable objects made with the
// NewRenderableSurface function, the start point is the bottom left.
func (renderable *Renderable) SetPosition(x, y float64) {

	xDisp := x - float64(renderable.verticies[0])
	yDisp := y - float64(renderable.verticies[1])

	renderable.Move(xDisp, yDisp)
}

// GetPosition returns the X and Y position of the first specified vertex of
// the Renderable object. If the Renderable object was created using
// NewRenderableSurface, the lower left vertex will be returned.
func (renderable *Renderable) GetPosition() (x, y float64) {

	return float64(renderable.verticies[0]), float64(renderable.verticies[1])
}
