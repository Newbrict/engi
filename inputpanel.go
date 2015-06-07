package engi

import "time"

// A InputPanel represents a Panel which responds to text input from the keyboard.
type InputPanel struct {
	*TextPanel
	Blinker  *Panel
	Button   *ButtonPanel
	Selected bool
	ticker   *time.Ticker
}

// NewInputPanel returns a new panel with a region of width
// w, height h and position 0, 0.  The texture of this
// region is set to a white pixel in order to be colored.
// The new panel is also appended to the children slice of
// the global gui struct.
func NewInputPanel(w, h int) *InputPanel {

	newtxtPanel := NewTextPanel(w, h)
	newtxtPanel.SetSize(float64(h))

	removeFromParent(newtxtPanel)

	newBlinker := NewPanel(1, h-2)
	newButton := NewButtonPanel(w, h)

	newinputPanel := &InputPanel{
		TextPanel: newtxtPanel,
		Blinker:   newBlinker,
		Button:    newButton,
		Selected:  false,
		ticker:    nil,
	}

	newinputPanel.SetFg(Color{0x00, 0x00, 0x00, 0xFF})
	newinputPanel.SetBg(Color{0xFF, 0xFF, 0xFF, 0xFF})

	newBlinker.SetParent(newinputPanel)
	newBlinker.SetBg(Black)
	newBlinker.BG.A = 0

	newButton.SetParent(newinputPanel)
	newButton.BG.A = 0
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
	if ip.text != "" {
		texWidth := float32(ip.texture.width) - 2
		//ip.width = texWidth
		x, y := ip.Blinker.Pos()
		if x != texWidth {
			ip.Blinker.SetPos(texWidth, y)
		}
	}

	if !ip.Button.Hovered() && ip.ticker != nil && ip.Selected && Cursor.Left {
		ip.Selected = false
		ip.ticker.Stop()
		ip.Blinker.BG.A = 0
	}

}

func (ip *InputPanel) GetPanel() *Panel {
	return ip.Panel
}
