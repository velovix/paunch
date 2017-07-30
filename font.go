package paunch

import (
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"io/ioutil"
)

// Font represents a font file. Most common font files are supported.
type Font struct {
	font *truetype.Font
}

// NewFont creates a new Font object based on the font file provided.
func NewFont(fontFile string) (*Font, error) {

	font := &Font{}

	fontData, err := ioutil.ReadFile(fontFile)
	if err != nil {
		return font, err
	}

	font.font, err = freetype.ParseFont(fontData)
	if err != nil {
		return font, err
	}

	return font, nil
}
