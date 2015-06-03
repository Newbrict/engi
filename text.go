package engi

type Text struct {
	font    *Font
	content string
	texture *Texture
}

func NewText(s string, f *Font) *Text {
	return &Text{
		font:    f,
		content: s,
		texture: f.Texture(s),
	}
}

func (t *Text) Render(b *Batch, render *RenderComponent, space *SpaceComponent) {
	t.texture.Render(b, render, space)
}

func (t *Text) SetText(s string) {
	if s == t.content {
		return
	}

	t.texture = t.font.Texture(s)
}

func (t *Text) Width() float32 {
	w, _, _ := t.font.TextDimensions(t.content)

	return float32(w)
}

func (t *Text) Height() float32 {
	_, h, _ := t.font.TextDimensions(t.content)

	return float32(h)
}
