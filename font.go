package engi

import (
	"image"
	"image/color"
	"image/draw"
	"io/ioutil"
	"log"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

var (
	dpi = float64(72)
)

type Color struct {
	R, G, B, A uint8
}

// TODO FG and BG color config
type Font struct {
	URL  string
	Size float64
	BG   Color
	FG   Color
	ttf  *truetype.Font
}

func (f *Font) Create() {
	url := f.URL

	// Read and parse the font
	ttfBytes, err := ioutil.ReadFile(url)
	if err != nil {
		log.Println(err)
		return
	}

	ttf, err := freetype.ParseFont(ttfBytes)
	if err != nil {
		log.Println(err)
		return
	}
	f.ttf = ttf
}

func (f *Font) TextDimensions(text string) (int, int, int) {
	fnt := f.ttf
	size := f.Size
	var (
		totalWidth  = fixed.Int26_6(0)
		totalHeight = fixed.Int26_6(size)
		maxYBearing = fixed.Int26_6(0)
	)
	fupe := fixed.Int26_6(fnt.FUnitsPerEm())
	for _, char := range text {
		idx := fnt.Index(char)
		hm := fnt.HMetric(fupe, idx)
		vm := fnt.VMetric(fupe, idx)
		g := truetype.GlyphBuf{}
		err := g.Load(fnt, fupe, idx, font.HintingNone)
		if err != nil {
			log.Println(err)
			return 0, 0, 0
		}
		totalWidth += hm.AdvanceWidth
		yB := (vm.TopSideBearing * fixed.Int26_6(size)) / fupe
		if yB > maxYBearing {
			maxYBearing = yB
		}
	}

	// Scale to actual pixel size
	totalWidth *= fixed.Int26_6(size)
	totalWidth /= fupe

	return int(totalWidth), int(totalHeight), int(maxYBearing)
}

func (f *Font) Render(text string) *Texture {
	width, height, yBearing := f.TextDimensions(text)
	font := f.ttf
	size := f.Size

	// Colors
	fg := image.NewUniform(color.NRGBA{f.FG.R, f.FG.G, f.FG.B, f.FG.A})
	bg := image.NewUniform(color.NRGBA{f.BG.R, f.BG.G, f.BG.B, f.BG.A})

	// Create the font context
	c := freetype.NewContext()

	nrgba := image.NewNRGBA(image.Rect(0, 0, width, height))
	draw.Draw(nrgba, nrgba.Bounds(), bg, image.ZP, draw.Src)

	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(size)
	c.SetClip(nrgba.Bounds())
	c.SetDst(nrgba)
	c.SetSrc(fg)

	// Draw the text.
	pt := freetype.Pt(0, int(yBearing))
	_, err := c.DrawString(text, pt)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Create texture
	imObj := &ImageObject{nrgba}
	return NewTexture(imObj)

}
