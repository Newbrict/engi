package engi

// A Event type has no input nor return parameters.
type Event func()

// A ButtonPanel represents a Panel
// which reacts to Mouse input.
type ButtonPanel struct {
	*Panel           // embedded panel
	hovered    bool  // hover flag
	DoHover    Event // function to call when hovered
	DoOffHover Event // functon to call when unhovered
	DoClick    Event // function to call when clicked
}

// NewButtonPanel returns a new panel with a region of width
// w, height h and position 0, 0.  The texture of this
// region is set to a white pixel in order to be colored.
// The new panel is also appended to the children slice of
// the global gui struct.
func NewButtonPanel(w, h int) *ButtonPanel {

	newPanel := NewPanel(w, h)
	removeFromParent(newPanel)

	doHover := func() { newPanel.SetBg(newPanel.BG.Shade(.80)) }
	dooffHover := func() { newPanel.SetBg(newPanel.BG.Shade(1.25)) }
	doClick := func() {}

	newbtnPanel := &ButtonPanel{
		Panel:      newPanel,
		hovered:    false,
		DoHover:    doHover,
		DoOffHover: dooffHover,
		DoClick:    doClick,
	}

	gPnls := &GUI.Children
	*gPnls = append(*gPnls, newbtnPanel)

	return newbtnPanel
}

func (btn *ButtonPanel) OnClick() {
	btn.DoClick()
}

func (btn *ButtonPanel) OnHover() {
	btn.DoHover()
}

func (btn *ButtonPanel) OnOffHover() {
	btn.DoOffHover()
}

func (btn *ButtonPanel) Hovered() bool {
	if (Cursor.X >= btn.X && Cursor.X <= btn.X+btn.width) && (Cursor.Y >= btn.Y && Cursor.Y <= btn.Y+btn.height) {
		return true
	}

	return false
}

// SetParent sets the parent panel and also
// appends to the children slice of parent.
// The panel is also removed from the previous
// parents child slice.  Position is set
// to 0,0 relative to parent.
func (btn *ButtonPanel) SetParent(graph Graphical) {

	removeFromParent(btn)

	btn.Parent = graph
	parent := graph.GetPanel()

	npC := &parent.GetPanel().Children
	*npC = append(*npC, btn)

	btn.Point = parent.Point
}

func (btn *ButtonPanel) Update() {

	if btn.Hovered() {
		if Cursor.Left && !Cursor.Click {
			btn.OnClick()
			Cursor.Click = true
		} else if !Cursor.Left && Cursor.Click {
			Cursor.Click = false
		}

		if !btn.hovered {
			btn.OnHover()
			window.SetCursor(Hand)
			btn.hovered = true
		}

	} else {
		if btn.hovered {
			btn.OnOffHover()
			window.SetCursor(Arrow)
			btn.hovered = false
		}
	}

}

func (btn *ButtonPanel) Draw(batch *Batch) {
	btn.Update()
	batch.Draw(btn, btn.X, btn.Y, 0, 0, 1, 1, 0, btn.BG)
}

func (btn *ButtonPanel) GetPanel() *Panel {
	return btn.Panel
}
