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

//Director ...
type Director struct {
	ib Builder
}

//NewImageDirector ...
func NewImageDirector(ib Builder) *Director {
	return &Director{ib}
}

//SetBuilder ...
func (d *Director) SetBuilder(ib Builder) {
	d.ib = ib
}

//BuildDefault ...
func (d *Director) BuildDefault() {
	d.ib.Reset()
	d.ib.SetHeight(500)
	d.ib.SetWidth(500)
	d.ib.SetColorful()
	d.ib.SetVisible()
}

//BuildFullHD ...
func (d *Director) BuildFullHD() {
	d.ib.Reset()
	d.ib.SetHeight(1080)
	d.ib.SetWidth(1920)
	d.ib.SetColorful()
	d.ib.SetVisible()
}

//BuildTransparent ...
func (d *Director) BuildTransparent() {
	d.ib.Reset()
	d.ib.SetHeight(500)
	d.ib.SetWidth(500)
	d.ib.SetResizable()
	d.ib.SetVisible()
}

//Builder ...
type Builder interface {
	Reset()
	SetHeight(h int)
	SetWidth(h int)
	SetColorful()
	SetResizable()
	SetVisible()
}

//Image ...
type Image struct {
	Height    int
	Width     int
	Colorful  bool
	Resizable bool
	Visible   bool
}

//ImageSettings ...
type ImageSettings struct {
	Height    int
	Width     int
	Colorful  bool
	Resizable bool
	Visible   bool
}

//ImageBuilder ...
type ImageBuilder struct {
	i Image
}

//NewImageBuilder ...
func NewImageBuilder() *ImageBuilder {
	return &ImageBuilder{}
}

//Reset ...
func (b *ImageBuilder) Reset() {
	b.i = Image{}
}

//SetHeight ...
func (b *ImageBuilder) SetHeight(h int) {
	b.i.Height = h
}

//SetWidth ...
func (b *ImageBuilder) SetWidth(w int) {
	b.i.Width = w
}

//SetColorful ...
func (b *ImageBuilder) SetColorful() {
	b.i.Colorful = true
}

//SetResizable ...
func (b *ImageBuilder) SetResizable() {
	b.i.Resizable = true
}

//SetVisible ...
func (b *ImageBuilder) SetVisible() {
	b.i.Visible = true
}

//GetResult ...
func (b *ImageBuilder) GetResult() Image {
	return b.i
}

//SettingsBuilder ...
type SettingsBuilder struct {
	s ImageSettings
}

//NewSettingsBuilder ...
func NewSettingsBuilder() *SettingsBuilder {
	return &SettingsBuilder{}
}

//Reset ...
func (b *SettingsBuilder) Reset() {
	b.s = ImageSettings{}
}

//SetHeight ...
func (b *SettingsBuilder) SetHeight(h int) {
	b.s.Height = h
}

//SetWidth ...
func (b *SettingsBuilder) SetWidth(w int) {
	b.s.Width = w
}

//SetColorful ...
func (b *SettingsBuilder) SetColorful() {
	b.s.Colorful = true
}

//SetResizable ...
func (b *SettingsBuilder) SetResizable() {
	b.s.Resizable = true
}

//SetVisible ...
func (b *SettingsBuilder) SetVisible() {
	b.s.Visible = true
}

//GetResult ...
func (b *SettingsBuilder) GetResult() ImageSettings {
	return b.s
}
