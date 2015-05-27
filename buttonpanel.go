package engi

type doClick func()

type ButtonPanel struct {
	Panel
	IsHovered, hovered bool
	DoClick            doClick
}

func NewButtonPanel(x, y float32, w, h int) *ButtonPanel {
	region := NewRegion(Files.Image("color.png"), 0, 0, w, h)
	point := Point{x, y}

	newPanel := &Panel{point, region, 0xFFFFFF, nil, make([]*Panel, 0)}
	newbtnPanel := &ButtonPanel{*newPanel, false, false, func() {}}

	btnPanel := NewEntity([]string{"RenderSystem"})
	btnpanelRender := NewRenderComponent(newbtnPanel, Point{1, 1}, "btnPanel")
	btnpanelSpace := SpaceComponent{newPanel.Point, 0, 0}
	btnPanel.AddComponent(&btnpanelRender)
	btnPanel.AddComponent(&btnpanelSpace)

	Wo.AddEntity(btnPanel)

	return newbtnPanel
}

func (btn *ButtonPanel) OnClick() {
	btn.DoClick()
}

func (btn *ButtonPanel) OnHover() {
	btn.SetBg(btn.Bg + 10)
}
func (btn *ButtonPanel) OffHover() {
	btn.SetBg(btn.Bg - 10)
}
