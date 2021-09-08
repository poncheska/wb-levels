package main

import "fmt"

func main() {
	ib := NewImageBuilder()
	d := NewImageDirector(ib)
	d.BuildDefault()
	fmt.Printf("%#v\n", ib.GetResult())
	d.BuildFullHD()
	fmt.Printf("%#v\n", ib.GetResult())

	sb := NewSettingsBuilder()
	d.SetBuilder(sb)
	d.BuildTransparent()
	fmt.Printf("%#v\n", sb.GetResult())
}

type Director struct {
	ib Builder
}

func NewImageDirector(ib Builder) *Director {
	return &Director{ib}
}

func (d *Director) SetBuilder(ib Builder) {
	d.ib = ib
}

func (d *Director) BuildDefault() {
	d.ib.Reset()
	d.ib.SetHeight(500)
	d.ib.SetWidth(500)
	d.ib.SetColorful()
	d.ib.SetVisible()
}

func (d *Director) BuildFullHD() {
	d.ib.Reset()
	d.ib.SetHeight(1080)
	d.ib.SetWidth(1920)
	d.ib.SetColorful()
	d.ib.SetVisible()
}

func (d *Director) BuildTransparent() {
	d.ib.Reset()
	d.ib.SetHeight(500)
	d.ib.SetWidth(500)
	d.ib.SetResizable()
	d.ib.SetVisible()
}

type Builder interface {
	Reset()
	SetHeight(h int)
	SetWidth(h int)
	SetColorful()
	SetResizable()
	SetVisible()
}

type Image struct {
	Height    int
	Width     int
	Colorful  bool
	Resizable bool
	Visible   bool
}

type ImageSettings struct {
	Height    int
	Width     int
	Colorful  bool
	Resizable bool
	Visible   bool
}

type ImageBuilder struct {
	i Image
}

func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{}
}

func (b *ImageBuilder) Reset() {
	b.i = Image{}
}

func (b *ImageBuilder) SetHeight(h int) {
	b.i.Height = h
}

func (b *ImageBuilder) SetWidth(w int) {
	b.i.Width = w
}

func (b *ImageBuilder) SetColorful() {
	b.i.Colorful = true
}

func (b *ImageBuilder) SetResizable() {
	b.i.Resizable = true
}

func (b *ImageBuilder) SetVisible() {
	b.i.Visible = true
}

func (b *ImageBuilder) GetResult() Image {
	return b.i
}

type SettingsBuilder struct {
	s ImageSettings
}

func NewSettingsBuilder() *SettingsBuilder {
	return &SettingsBuilder{}
}

func (b *SettingsBuilder) Reset() {
	b.s = ImageSettings{}
}

func (b *SettingsBuilder) SetHeight(h int) {
	b.s.Height = h
}

func (b *SettingsBuilder) SetWidth(w int) {
	b.s.Width = w
}

func (b *SettingsBuilder) SetColorful() {
	b.s.Colorful = true
}

func (b *SettingsBuilder) SetResizable() {
	b.s.Resizable = true
}

func (b *SettingsBuilder) SetVisible() {
	b.s.Visible = true
}

func (b *SettingsBuilder) GetResult() ImageSettings {
	return b.s
}
