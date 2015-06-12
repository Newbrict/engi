// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package engi

type Responder interface {
	Render()
	Resize(width, height int)
	Preload()
	Setup()
	Close()
	Update(dt float32)
	Mouse(x, y float32, action Action, btn MouseButton)
	Scroll(amount float32)
	Key(key Key, action Action, mod Modifier)
	Type(char rune)
	AddEntity(e *Entity)
	Batch() *Batch
	New()
}

type Game struct{}

func (g *Game) Preload()                                           {}
func (g *Game) Setup()                                             {}
func (g *Game) Close()                                             {}
func (g *Game) Update(dt float32)                                  {}
func (g *Game) Render()                                            {}
func (g *Game) Resize(w, h int)                                    {}
func (g *Game) Mouse(x, y float32, action Action, btn MouseButton) {}
func (g *Game) Scroll(amount float32)                              {}
func (g *Game) Key(key Key, action Action, mod Modifier) {
	if key == Escape {
		Exit()
	} else if key == Backspace && action == PRESS {
		for _, pnl := range GUI.InputPanels {
			if pnl.Selected {
				size := len(pnl.TextPanel.text)
				if size > 0 {
					pnl.TextPanel.SetText(pnl.TextPanel.text[:size-1])
				} else if size == 0 {
					pnl.TextPanel.SetText("")
				}
				break
			}

		}
	}
}
func (g *Game) Type(char rune) {
	for _, pnl := range GUI.InputPanels {
		if pnl.Selected {
			pnl.TextPanel.SetText(pnl.TextPanel.text + string(char))
			break
		}
	}
}
