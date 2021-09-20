package main

import "fmt"

func main() {
	data := []Data{
		&BytesData{[]byte{0, 10, 2}},
		&BytesData{[]byte{0, 1, 23}},
		&StringData{"gsgdsdg"},
		&StringData{"gsgdfsdg"},
		&StringData{"gsfgdsdg"},
		&BytesData{[]byte{10, 10, 23}},
		&StringData{"gsgdsdgf"},
		&BytesData{[]byte{2, 10, 23}},
		&StringData{"fgsgdsdg"},
		&StringData{"gsgffdsdg"},
		&StringData{"gsgdsdffg"},
	}
	c := NewDataCounter()
	for _, v := range data {
		v.Accept(c)
	}
	fmt.Println(c.GetResult())
}

//Visitor ...
type Visitor interface {
	VisitBytesData(d *BytesData)
	VisitStringData(d *StringData)
}

//DataCounter ...
type DataCounter struct {
	c int
}

//NewDataCounter ...
func NewDataCounter() *DataCounter {
	return &DataCounter{0}
}

//VisitBytesData ...
func (c *DataCounter) VisitBytesData(d *BytesData) {
	c.c += len(d.Bytes)
}

//VisitStringData ...
func (c *DataCounter) VisitStringData(d *StringData) {
	c.c += len([]byte(d.String))
}

//GetResult ...
func (c *DataCounter) GetResult() int {
	return c.c
}

//Data ...
type Data interface {
	Print()
	Accept(v Visitor)
}

//BytesData ...
type BytesData struct {
	Bytes []byte
}

//Print ...
func (d *BytesData) Print() {
	fmt.Println(string(d.Bytes))
}

//Accept ...
func (d *BytesData) Accept(v Visitor) {
	v.VisitBytesData(d)
}

//StringData ...
type StringData struct {
	String string
}

//Print ...
func (d *StringData) Print() {
	fmt.Println(d.String)
}

//Accept ...
func (d *StringData) Accept(v Visitor) {
	v.VisitStringData(d)
}
