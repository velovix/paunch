package paunch

import (
	gl "github.com/chsc/gogl/gl33"
	"image"
	"image/png"
	"os"
)

// Renderable is an object that can be drawn on the screen
type Renderable struct {
	mode            int
	size            int
	vertex_buffer   gl.Uint
	texcoord_buffer gl.Uint
	texture         []gl.Uint
	verticies       []float32
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

// NewRenderable returns a new Renderable object based on the specified shape
// type and verticies.
func NewRenderable(mode int, verticies []float32) (Renderable, error) {

	renderable := Renderable{mode, len(verticies), 0, 0, nil, verticies}

	gl.GenBuffers(1, &renderable.vertex_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.vertex_buffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(verticies)*4), gl.Pointer(&verticies[0]), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	return renderable, checkForErrors()
}

// Texture applies a texture from a 32-bit PNG file to a Renderable. The
// Renderable will automatically be drawn with this texture. The texture may
// also be split into multiple smaller textures automatically if clip is > 1.
func (renderable *Renderable) Texture(coords []float32, filename string, clip int) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gl.GenBuffers(1, &renderable.texcoord_buffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(renderable.texcoord_buffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(coords)*4), gl.Pointer(&coords[0]), gl.STREAM_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	data, decodeErr := png.Decode(file)
	if decodeErr != nil {
		return decodeErr
	}

	clipWidth := data.Bounds().Max.X - data.Bounds().Min.X
	clipHeight := (data.Bounds().Max.Y - data.Bounds().Min.Y) / clip

	renderable.texture = make([]gl.Uint, clip)
	gl.GenTextures(gl.Sizei(clip), &renderable.texture[0])

	byteData := imageToBytes(data)
	clips := make([][]byte, clip)
	for i, _ := range clips {
		clips[i] = byteData[i*(len(byteData)/len(clips)) : (i+1)*(len(byteData)/len(clips))]
		gl.BindTexture(gl.TEXTURE_2D, renderable.texture[len(clips)-1-i])
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
		gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
		gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA,
			gl.Sizei(clipWidth), gl.Sizei(clipHeight),
			0, gl.RGBA, gl.UNSIGNED_BYTE,
			gl.Pointer(&clips[i][0]))
	}

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return checkForErrors()
}

// Draw draws the Renderable.
func (renderable *Renderable) Draw(frame int) error {

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertex_buffer)
	gl.VertexAttribPointer(gl.Uint(0), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if renderable.texcoord_buffer != 0 {
		gl.ActiveTexture(gl.TEXTURE0)

		gl.BindBuffer(gl.ARRAY_BUFFER, renderable.texcoord_buffer)
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

	gl.BindBuffer(gl.ARRAY_BUFFER, renderable.vertex_buffer)
	gl.BufferSubData(gl.ARRAY_BUFFER, 0, gl.Sizeiptr(len(renderable.verticies)*4), gl.Pointer(&renderable.verticies[0]))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
}
