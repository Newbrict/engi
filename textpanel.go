package engi

type TextPanel struct {
	*Panel
	font *Font
	text string
}

// NewTextPanel returns a new text panel with a region of width
// w, height h and position 0, 0.  The texture of this
// region is set to a white pixel in order to be colored.
// The new panel is also appended to the children slice of
// the global gui struct.
func NewTextPanel(w, h int) *TextPanel {

	// should be replaced by newfont
	fg := Color{0x00, 0x00, 0x00, 0xFF}
	bg := Color{0x00, 0x00, 0x00, 0x00}
	size := 15.00
	url := "data/fonts/tahoma.ttf"

	newFont := &Font{url, size, bg, fg, nil}
	newFont.Create()

	newPanel := NewPanel(w, h)
	removeFromParent(newPanel)

	newtxtPanel := &TextPanel{
		Panel: newPanel,
		font:  newFont,
		text:  "",
	}

	gPnls := &GUI.Children
	*gPnls = append(*gPnls, newtxtPanel)

	return newtxtPanel
}

func (tp *TextPanel) RenderText() {
	if tp.text != "" {
		tex := tp.font.Render(tp.text)
		tp.SetTexture(tex)
	}
}

// SetParent sets the parent panel and also
// appends to the children slice of parent.
// The panel is also removed from the previous
// parents child slice.  Position is set
// to 0,0 relative to parent.
func (tp *TextPanel) SetParent(graph Graphical) {

	removeFromParent(tp)

	tp.Parent = graph
	parent := graph.GetPanel()

	npC := &parent.GetPanel().Children
	*npC = append(*npC, tp)

	tp.Point = parent.Point
}

func (tp *TextPanel) SetFont(font string) {
	defer tp.RenderText()

	tp.font.URL = "data/fonts/" + font
	tp.font.Create()
}

func (tp *TextPanel) SetText(txt string) {
	defer tp.RenderText()

	tp.text = txt
}

func (tp *TextPanel) SetFg(fg Color) {
	defer tp.RenderText()

	tp.font.FG = fg
}

func (tp *TextPanel) SetBg(bg Color) {
	defer tp.RenderText()

	tp.font.BG = bg
}

func (tp *TextPanel) SetSize(size float64) {
	defer tp.RenderText()

	tp.font.Size = size
}

func (tp *TextPanel) GetPanel() *Panel {
	return tp.Panel
}

func (tp *TextPanel) Draw(batch *Batch) {
	tp.Update()
	batch.Draw(tp, tp.X, tp.Y, 0, 0, 1, 1, 0, tp.BG)
}

func (tp *TextPanel) Update() {

}
