package engi

// Constants for each edge of game window.
const (
	LEFT int = iota
	TOP
	RIGHT
	BOTTOM
)

var (
	White Color = Color{255, 255, 255, 255}
	Black Color = Color{0, 0, 0, 255}
	Red   Color = Color{255, 0, 0, 255}
	Green Color = Color{0, 255, 0, 255}
	Blue  Color = Color{0, 0, 255, 255}
)

type Graphical interface {
	Draw(*Batch)
	Update()
	SetParent(Graphical)
	GetPanel() *Panel
}

// A Mouse represents the state of the current input mouse.
type Mouse struct {
	Point                           // position of cursor relative to game window
	Left, Right, Middle, Click bool // flags for input
}

// A Panel represents a region with a parent and children.
type Panel struct {
	Point                   // absolute position
	*Region                 // region to be rendered
	BG          Color       // background color of region
	Visible     bool        // visibility of panel
	Transparent int         // transparency
	Parent      Graphical   // parent panel
	Children    []Graphical // slice containing children panels
}

type Gui struct {
	*Panel
	InputPanels []*InputPanel
}

func (gui *Gui) Render(b *Batch, render *RenderComponent, space *SpaceComponent) {
	drawChildren(gui.Children)
}

func NewGui() *Gui {
	panel := &Panel{
		Point:    Point{},
		Region:   &Region{},
		BG:       Color{},
		Visible:  true,
		Parent:   nil,
		Children: make([]Graphical, 0),
	}

	return &Gui{
		Panel:       panel,
		InputPanels: make([]*InputPanel, 0),
	}
}

// NewPanel returns a new panel with a region of width
// w, height h and position 0, 0.  The texture of this
// region is set to a white pixel in order to be colored.
// The new panel is also appended to the children slice of
// the global gui struct.
func NewPanel(w, h int) *Panel {
	region := NewRegion(nil, 0, 0, w, h)

	newPanel := &Panel{
		Point:    Point{},
		Region:   region,
		BG:       White,
		Visible:  true,
		Parent:   GUI.Panel,
		Children: make([]Graphical, 0),
	}

	gPnls := &GUI.Children
	*gPnls = append(*gPnls, newPanel)

	return newPanel
}

// DrawChildren draws each child by traversing through all the children of each parent.
func drawChildren(children []Graphical) {
	for _, child := range children {
		pnl := child.GetPanel()
		if pnl.Visible {
			child.Draw(Wo.Batch())
			drawChildren(pnl.Children)
		}
	}
}

// panelPosition returns the index of child in its parent slice.
func panelPosition(child Graphical, children []Graphical) int {
	for i, pnl := range children {
		if pnl == child {
			return i
		}
	}
	return -1
}

func removeFromParent(child Graphical) {
	pnl := child.GetPanel()
	pC := &pnl.Parent.GetPanel().Children
	cI := panelPosition(child, *pC)
	*pC = append((*pC)[:cI], (*pC)[cI+1:]...)
}

// SetParent sets the parent panel and also
// appends to the children slice of parent.
// The panel is also removed from the previous
// parents child slice.  Position is set
// to 0,0 relative to parent.
func (pnl *Panel) SetParent(graph Graphical) {
	removeFromParent(pnl)

	pnl.Parent = graph
	parent := graph.GetPanel()

	npC := &parent.GetPanel().Children
	*npC = append(*npC, pnl)

	pnl.Point = parent.Point
}

// SetTexture sets the texture of the panels region
// without modifiying the textures rendered dimensions.
func (pnl *Panel) SetTexture(tex *Texture) {
	pnl.texture = tex

	invTexWidth := 1.0 / float32(tex.Width())
	invTexHeight := 1.0 / float32(tex.Height())

	pnl.u2 = float32(pnl.width) * invTexWidth
	pnl.v2 = float32(pnl.height) * invTexHeight
}

// SetPos sets the position relative to parent.
// All of the panels childrens positions are also
// updated relative to the panel.
func (pnl *Panel) SetPos(x, y float32) {
	parent := pnl.Parent.GetPanel()

	oldPos := pnl.Point
	pnl.Point = Point{parent.X + x, parent.Y + y}

	for _, child := range pnl.Children {
		childPanel := child.GetPanel()

		childPanel.SetPos(childPanel.X-oldPos.X, childPanel.Y-oldPos.Y)
	}
}

func (pnl *Panel) SetSize(width, height float32) {
	pnl.width = width
	pnl.height = height
}

// SizeToContainer sizes the current texture
// to the width and height of the panel.
func (pnl *Panel) SizeToContainer() {
	invTexWidth := 1.0 / float32(pnl.width)
	invTexHeight := 1.0 / float32(pnl.height)

	pnl.u2 = float32(pnl.width) * invTexWidth
	pnl.v2 = float32(pnl.height) * invTexHeight
}

// SizeToContents sizes the panel to the size of its texture.
func (pnl *Panel) SizeToContents() {
	pnl.width = pnl.texture.Width()
	pnl.height = pnl.texture.Height()

	pnl.SizeToContainer()
}

// Align aligns the panel to the edge of its parent.
func (pnl *Panel) Align(edge int) {
	pnl.AlignOff(edge, 0)
}

// AlignOff aligns the panel to the edge of its parent.
func (pnl *Panel) AlignOff(edge int, off float32) {
	x, y := pnl.Pos()
	parent := pnl.Parent.GetPanel()

	switch edge {
	case LEFT:
		x = off
	case TOP:
		y = off
	case RIGHT:
		x = (parent.width - pnl.width) - off
	case BOTTOM:
		y = (parent.height - pnl.height) - off

	}

	pnl.SetPos(x, y)
}

// Center sets the position of the panel to the center of its parent.
func (pnl *Panel) Center() {
	parent := pnl.Parent.GetPanel()

	x := parent.width/2.0 - pnl.width/2.0
	y := parent.height/2.0 - pnl.height/2.0

	pnl.SetPos(x, y)
}

func (pnl *Panel) SetBg(color Color) {
	pnl.BG = color
}

func (pnl *Panel) Width() float32 {
	return pnl.width
}

func (pnl *Panel) Height() float32 {
	return pnl.height
}

// Pos returns the panels position relative to its parent.
func (pnl *Panel) Pos() (x, y float32) {
	parent := pnl.Parent.GetPanel()
	x = pnl.X - parent.X
	y = pnl.Y - parent.Y

	return x, y
}

func (pnl *Panel) Draw(batch *Batch) {
	batch.Draw(pnl, pnl.X, pnl.Y, 0, 0, 1, 1, 0, pnl.BG)
}

func (pnl *Panel) GetPanel() *Panel {
	return pnl
}

func (pnl *Panel) Update() {

}
