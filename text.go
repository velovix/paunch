package paunch

import (
	"code.google.com/p/freetype-go/freetype"
	"image"
	"image/color"
	"image/draw"
)

const dpi = 72

// Text is an object that represents drawable text. Text is used similarly to
// Renderable objects.
type Text struct {
	message string

	context  *freetype.Context
	fontSize float64

	x float64
	y float64

	renderable Renderable
	fontColor  *image.Uniform
}

// NewText creates a new Text object. The x and y positions represent the left
// and bottom of text without tails.
func NewText(x, y float64, font Font, fontSize float64, message string) (Text, error) {

	var text Text

	text.x = x
	text.y = y - fontSize

	text.fontColor = image.NewUniform(color.RGBA{0, 0, 0, 255})

	text.context = freetype.NewContext()
	text.context.SetDPI(dpi)
	text.context.SetFont(font.font)
	text.context.SetFontSize(fontSize)
	text.context.SetSrc(text.fontColor)
	text.context.SetHinting(freetype.FullHinting)

	text.fontSize = fontSize

	err := text.SetMessage(message)
	if err != nil {
		return text, err
	}

	return text, nil
}

// Draw draws the Text object.
func (text Text) Draw() {

	text.renderable.Draw(0)
}

// SetMessage changes the message displayed by the Text object.
func (text *Text) SetMessage(message string) error {

	text.message = message
	err := text.updateText()
	if err != nil {
		return err
	}

	return nil
}

// GetMessage returns the current message displayed by the Text object.
func (text *Text) GetMessage() string {

	return text.message
}

// SetColor sets the color of the Text object. The default is black
// (0, 0, 0, 255).
func (text *Text) SetColor(r, g, b, a uint8) error {

	text.fontColor = image.NewUniform(color.RGBA{r, g, b, a})
	text.context.SetSrc(text.fontColor)
	err := text.updateText()
	if err != nil {
		return err
	}

	return nil
}

// GetPosition returns the position of the Text object.
func (text *Text) GetPosition() (float64, float64) {

	return text.x, text.y + text.fontSize
}

// Move moves the Text object the specified distance.
func (text *Text) Move(x, y float64) {

	text.x += x
	text.y += y - text.fontSize

	text.renderable.Move(x, y-text.fontSize)
}

// SetPosition sets the Text object's position to the specified point.
func (text *Text) SetPosition(x, y float64) {

	text.x = x
	text.y = y - text.fontSize

	text.renderable.SetPosition(x, y-text.fontSize)
}

func flipRGBA(src *image.RGBA) {

	srcCopy := image.NewRGBA(src.Bounds())
	for y := 0; y < src.Bounds().Max.Y-src.Bounds().Min.Y; y++ {
		for x := 0; x < src.Bounds().Max.X-src.Bounds().Min.X; x++ {
			srcCopy.Set(x, y, src.At(x, y))
		}
	}

	for y := 0; y < src.Bounds().Max.Y-src.Bounds().Min.Y; y++ {
		for x := 0; x < src.Bounds().Max.X-src.Bounds().Min.X; x++ {
			src.Set(x, y, srcCopy.At(x, (src.Bounds().Max.Y-src.Bounds().Min.Y)-1-y))
		}
	}
}

func (text *Text) findTextDimensions() (int, int, error) {

	tempRGBA := image.NewRGBA(image.Rect(0, 0, 1, 1))

	draw.Draw(tempRGBA, tempRGBA.Bounds(), image.Transparent, image.ZP, draw.Src)

	text.context.SetDst(tempRGBA)
	text.context.SetClip(tempRGBA.Bounds())
	pt := freetype.Pt(0, int(text.context.PointToFix32(text.fontSize)>>8))
	endPos, err := text.context.DrawString(text.message, pt)
	if err != nil {
		return 0, 0, err
	}

	return int(endPos.X >> 8), int(text.context.PointToFix32(text.fontSize*2) >> 8), nil
}

func (text *Text) updateText() error {

	width, height, err := text.findTextDimensions()
	if err != nil {
		return err
	}
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(rgba, rgba.Bounds(), image.Transparent, image.ZP, draw.Src)

	text.context.SetDst(rgba)
	text.context.SetClip(rgba.Bounds())
	pt := freetype.Pt(0, int(text.context.PointToFix32(text.fontSize)>>8))
	_, err = text.context.DrawString(text.message, pt)
	if err != nil {
		return err
	}

	flipRGBA(rgba)
	text.renderable, err = NewRenderableFromData(text.x, text.y, float64(width), float64(height), rgba.Pix, 1)
	if err != nil {
		return err
	}

	return nil
}
