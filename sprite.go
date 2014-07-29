package paunch

import (
	"github.com/velovix/gl"
	"image"
	"image/png"
	"os"
)

// Sprite is a textured object, an object that is displayed on the screen using
// image data.
type Sprite struct {
	texcoordBuffer gl.Uint
	texture        []gl.Uint
	shape          *Shape
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

// NewSprite creates a new Sprite object using the given data, which is
// expected to be in RGBA format. If you use PNG image files, you can use the
// NewSpriteFromImage shortcut function instead.
func NewSprite(x, y, width, height float64, data []byte, clip int) (*Sprite, error) {

	verticies := []float64{
		x, y,
		x + width, y,
		x, y + height,

		x + width, y + height,
		x + width, y,
		x, y + height}

	shape, err := NewShape(gl.TRIANGLES, verticies)
	if err != nil {
		return nil, err
	}

	sprite := &Sprite{texcoordBuffer: 0, texture: nil, shape: shape}

	texCoords := []float32{
		0, 0,
		1, 0,
		0, 1,

		1, 1,
		1, 0,
		0, 1}

	gl.GenBuffers(1, &sprite.texcoordBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, gl.Uint(sprite.texcoordBuffer))
	gl.BufferData(gl.ARRAY_BUFFER, gl.Sizeiptr(len(texCoords)*4), gl.Pointer(&texCoords[0]), gl.STREAM_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	sprite.texture = make([]gl.Uint, clip)
	gl.GenTextures(gl.Sizei(clip), &sprite.texture[0])

	clips := make([][]byte, clip)
	for i := range clips {
		clips[i] = data[i*(len(data)/len(clips)) : (i+1)*(len(data)/len(clips))]
		gl.BindTexture(gl.TEXTURE_2D, sprite.texture[len(clips)-1-i])
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

	return sprite, checkForErrors()
}

// NewSpriteFromImage creates a new Sprite object using the given PNG image
// file.
func NewSpriteFromImage(x, y float64, filename string, clip int) (*Sprite, error) {

	var sprite *Sprite

	file, err := os.Open(filename)
	if err != nil {
		return sprite, err
	}
	defer file.Close()

	data, err := png.Decode(file)
	if err != nil {
		return sprite, err
	}

	width, height, byteData := imageToBytes(data)
	sprite, err = NewSprite(x, y, float64(width), float64(height/clip), byteData, clip)
	if err != nil {
		return sprite, err
	}

	return sprite, nil
}

// NewSpriteFromSprite creates a new Sprite object that uses the same image
// data as the supplied Sprite object. This can serve to save a lot of GPU
// memory when dealing with Sprite objects that use image data.
func NewSpriteFromSprite(sprite *Sprite) (*Sprite, error) {

	shape, err := NewShapeFromShape(sprite.shape)
	if err != nil {
		return nil, err
	}

	newSprite := &Sprite{texcoordBuffer: sprite.texcoordBuffer, texture: sprite.texture,
		shape: shape}

	return newSprite, checkForErrors()
}

// SetScaling sets the scaling factor of the Sprite object. For instance, an x
// and y scale value of two will make the Sprite object twice as large.
func (sprite *Sprite) SetScaling(xScale, yScale float64) {

	sprite.shape.SetScaling(xScale, yScale)
}

// Draw draws the Sprite object.
func (sprite *Sprite) Draw(frame int) error {

	gl.BindBuffer(gl.ARRAY_BUFFER, sprite.shape.vertexBuffer)
	gl.VertexAttribPointer(gl.Uint(0), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
	gl.BindAttribLocation(paunchEffect.program, gl.Uint(0), gl.GLString("position"))
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	if sprite.texcoordBuffer != 0 {
		gl.ActiveTexture(gl.TEXTURE0)

		gl.BindBuffer(gl.ARRAY_BUFFER, sprite.texcoordBuffer)
		gl.VertexAttribPointer(gl.Uint(1), 2, gl.FLOAT, gl.FALSE, 0, gl.Offset(nil, 0))
		gl.BindAttribLocation(paunchEffect.program, gl.Uint(0), gl.GLString("texcoord"))
		gl.BindBuffer(gl.ARRAY_BUFFER, 0)

		gl.BindTexture(gl.TEXTURE_2D, sprite.texture[frame])
		gl.EnableVertexAttribArray(gl.Uint(1))
	}

	gl.EnableVertexAttribArray(gl.Uint(0))
	gl.DrawArrays(gl.TRIANGLES, 0, gl.Sizei(sprite.shape.size))
	gl.DisableVertexAttribArray(gl.Uint(0))
	gl.DisableVertexAttribArray(gl.Uint(1))

	gl.BindTexture(gl.TEXTURE_2D, 0)

	return checkForErrors()
}

// Move moves the Sprite object a specified distance.
func (sprite *Sprite) Move(x, y float64) {

	sprite.shape.Move(x, y)
}

// SetPosition sets the position of the Sprite object relative to the bottom-
// left corner.
func (sprite *Sprite) SetPosition(x, y float64) {

	sprite.shape.SetPosition(x, y)
}

// GetPosition returns the X and Y position of the bottom-left corner of the
// Sprite object.
func (sprite *Sprite) GetPosition() (x, y float64) {

	return sprite.shape.GetPosition()
}
