package engi

type Panel struct {
	Point
	r        *Region
	Bg       uint32
	Parent   *Panel
	Children []*Panel
}

type Mouse struct {
	Point
	Left, Right, Click bool
}

func NewPanel(x, y float32, w, h int) *Panel {

	region := NewRegion(Files.Image("color.png"), 0, 0, w, h)
	point := Point{x, y}

	newPanel := &Panel{point, region, 0xFFFFFF, nil, make([]*Panel, 0)}

	panel := NewEntity([]string{"RenderSystem"})
	panelRender := NewRenderComponent(newPanel, Point{1, 1}, "panel")
	panelSpace := SpaceComponent{newPanel.Point, 0, 0}
	panel.AddComponent(&panelRender)
	panel.AddComponent(&panelSpace)

	Wo.AddEntity(panel)

	return newPanel
}

func (p *Panel) SetParent(parent *Panel) {
	p.Parent = parent
	parent.Children = append(parent.Children, p)

	p.Point = Point{0, 0}
}

func (p *Panel) SetTexture(tex *Texture) {
	if tex != p.r.texture {
		p.r.UpdateTexture(tex)
	}
}

/*
func (p *Panel) SetText( text string ) {

	txtFont := NewGridFont(Files.Image("font.png"), 20, 20)
	txt := NewText( text, txtFont )

	txtEnt := NewEntity([]string{"RenderSystem"})
	txtRender := NewRenderComponent(txt, Point{1, 1}, "text")
	txtSpace := SpaceComponent{p.Point, 0, 0}

	txtEnt.AddComponent(&txtRender)
	txtEnt.AddComponent(&txtSpace)

	p.txt = txt
	Wo.AddEntity(txtEnt)
}
*/
func (p *Panel) SizeToContainer() {
	invTexWidth := 1.0 / p.r.width
	invTexHeight := 1.0 / p.r.height
	//u := float32(x) * invTexWidth
	//v := float32(y) * invTexHeight
	p.r.u2 = (float32(p.r.width) + p.Point.X) * invTexWidth
	p.r.v2 = (float32(p.r.height) + p.Point.Y) * invTexHeight
}

func (p *Panel) SizeToContents() {
	p.r.width = p.r.texture.Width()
	p.r.height = p.r.texture.Height()

	p.SizeToContainer()
}

func (p Panel) Width() float32 {
	return p.r.Height()
}

func (p Panel) Height() float32 {
	return p.r.Width()
}

func (p *Panel) SetPos(x, y float32) {
	p.Point = Point{x, y}
}

func (p *Panel) Align(edge int) {

	if p.Parent == nil {
		switch edge {
		case 1:
			p.Point.X = 0
		case 2:
			p.Point.Y = 0
		case 3:
			p.Point.X = Width()
		case 4:
			p.Point.Y = Height()
		}
	} else {
		switch edge {
		case 1:
			p.Point.X = 0
		case 2:
			p.Point.Y = 0
		case 3:
			p.Point.X = p.Parent.r.width
		case 4:
			p.Point.Y = p.Parent.r.height
		}

	}
}

func (p *Panel) AlignOff(edge int, off float32) {

	if p.Parent == nil {
		switch edge {
		case 1:
			p.Point.X = 0 + off
		case 2:
			p.Point.Y = 0 + off
		case 3:
			p.Point.X = Width() - off
		case 4:
			p.Point.Y = Height() - off
		}
	} else {
		switch edge {
		case 1:
			p.Point.X = off
		case 2:
			p.Point.Y = off
		case 3:
			p.Point.X = p.Parent.r.width - off
		case 4:
			p.Point.Y = p.Parent.r.height - off
		}

	}
}

func (p *Panel) Center() {
	var x, y float32

	if p.Parent == nil {
		x = Width()/2.0 - p.r.width/2.0
		y = Height()/2.0 - p.r.height/2.0
	} else {
		x = p.Parent.r.width/2.0 - p.r.width/2.0
		y = p.Parent.r.height/2.0 - p.r.height/2.0
	}
	p.Point = Point{x, y}
}

func (p *Panel) SetBg(color uint32) {
	if p.Bg != color {
		p.Bg = color
	}
}
