package engi

//resize repostions
//import "fmt"
type Panel struct {
	Point
	r  *Region
	bg uint32
}

func NewPanel(x, y float32, w, h int) *Panel {

	region := NewRegion(Files.Image("box.png"), 0, 0, w, h)
	point := Point{x, y}

	newPanel := &Panel{point, region, 0xCC0000}

	panel := NewEntity([]string{"RenderSystem"})
	panelRender := NewRenderComponent(newPanel, Point{1, 1}, "panel")
	panelSpace := SpaceComponent{newPanel.Point, 0, 0}
	panel.AddComponent(&panelRender)
	panel.AddComponent(&panelSpace)

	Wo.AddEntity(panel)

	return newPanel
}

func (p *Panel) SizeToContainer() {
	invTexWidth := 1.0 / p.r.width
	invTexHeight := 1.0 / p.r.height
	//u := float32(x) * invTexWidth
	//v := float32(y) * invTexHeight
	p.r.u2 = (float32(p.r.width) + p.Point.X) * invTexWidth
	p.r.v2 = (float32(p.r.height) + p.Point.Y) * invTexHeight
}

func (p *Panel) SizeToContents() {
	p.r.width = 48
	p.r.height = 64
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

func (p *Panel) Center() {
	x := Width()/2.0 - p.r.width/2.0
	y := Height()/2.0 - p.r.height/2.0

	p.Point = Point{x, y}
}

func (p *Panel) SetBg(color uint32) {
	p.bg = color
}

//func (p *Panel) Paint( )

/*
type Region struct {
	texture       *Texture
	u, v          float32
	u2, v2        float32
	width, height float32
}
*/
