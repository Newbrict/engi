package engi

import "time"

// A InputPanel represents a Panel which responds to text input from the keyboard.
type InputPanel struct {
	*Panel
	TextPanel *TextPanel
	Blinker   *Panel
	Button    *ButtonPanel
	Selected  bool
	ticker    *time.Ticker
}

// NewInputPanel returns a new panel with a region of width
// w, height h and position 0, 0.  The texture of this
// region is set to a white pixel in order to be colored.
// The new panel is also appended to the children slice of
// the global gui struct.
func NewInputPanel(w, h int) *InputPanel {
	newPanel := NewPanel(w, h)
	removeFromParent(newPanel)

	newtxtPanel := NewTextPanel(w, h)
	newtxtPanel.SetSize(float64(h))

	newBlinker := NewPanel(1, h-2)
	newButton := NewButtonPanel(w, h)

	newinputPanel := &InputPanel{
		Panel:     newPanel,
		TextPanel: newtxtPanel,
		Blinker:   newBlinker,
		Button:    newButton,
		Selected:  false,
		ticker:    nil,
	}

	newtxtPanel.SetParent(newinputPanel)
	newBlinker.SetParent(newinputPanel)
	newButton.SetParent(newinputPanel)

	newtxtPanel.SetFg(Color{0x00, 0x00, 0x00, 0xFF})
	newtxtPanel.SetBg(Color{0x00, 0x00, 0x00, 0x00})
	newtxtPanel.SetPos(0, 5)

	newBlinker.SetBg(Black)
	newBlinker.BG.A = 0

	newButton.BG.A = 0
	newButton.Cursor = IBEAM
	newButton.DoClick = func() {
		if !newinputPanel.Selected {
			ticker := time.NewTicker(time.Millisecond * 700)
			newinputPanel.Selected = true
			newinputPanel.ticker = ticker

			go func() {
				for _ = range ticker.C {
					if newBlinker.BG.A == 0 {
						newBlinker.BG.A = 255
					} else {
						newBlinker.BG.A = 0
					}
				}
			}()
		}
	}

	newButton.DoHover = func() {}
	newButton.DoOffHover = func() {}

	newBlinker.Center()
	newBlinker.AlignOff(LEFT, 2)

	gPnls := &GUI.Children
	*gPnls = append(*gPnls, newinputPanel)

	ipPnls := &GUI.InputPanels
	*ipPnls = append(*ipPnls, newinputPanel)

	return newinputPanel
}

// SetParent sets the parent panel and also
// appends to the children slice of parent.
// Position is set to 0,0 relative to parent.
func (ip *InputPanel) SetParent(graph Graphical) {
	removeFromParent(ip)

	ip.Parent = graph
	parent := graph.GetPanel()

	npC := &parent.GetPanel().Children
	*npC = append(*npC, ip)

	ip.Point = parent.Point
}

func (ip *InputPanel) Draw(batch *Batch) {
	ip.Update()
	batch.Draw(ip, ip.X, ip.Y, 0, 0, 1, 1, 0, ip.BG)
}

func (ip *InputPanel) Update() {
	x, y := ip.Blinker.Pos()
	texWidth := float32(ip.TextPanel.texture.width)

	if x != texWidth && texWidth <= ip.Width() {
		ip.Blinker.SetPos(texWidth, y)
	} else if texWidth >= ip.Width() && x != ip.Width()-2 {
		ip.Blinker.SetPos(ip.Width()-2, y)
	}

	if !ip.Button.Hovering() && ip.ticker != nil && ip.Selected && Cursor.Left {
		ip.Selected = false
		ip.ticker.Stop()
		ip.Blinker.BG.A = 0
	}

}

func (ip *InputPanel) GetPanel() *Panel {
	return ip.Panel
}
