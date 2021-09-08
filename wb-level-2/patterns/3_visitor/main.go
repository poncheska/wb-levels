package main

import "fmt"

func main() {
	data := []Data{
		&BytesData{[]byte{0,10,2}},
		&BytesData{[]byte{0,1,23}},
		&StringData{"gsgdsdg"},
		&StringData{"gsgdfsdg"},
		&StringData{"gsfgdsdg"},
		&BytesData{[]byte{10,10,23}},
		&StringData{"gsgdsdgf"},
		&BytesData{[]byte{2,10,23}},
		&StringData{"fgsgdsdg"},
		&StringData{"gsgffdsdg"},
		&StringData{"gsgdsdffg"},
	}
	c := NewDataCounter()
	for _,v := range data{
		v.Accept(c)
	}
	fmt.Println(c.GetResult())
}

type Visitor interface {
	VisitBytesData(d *BytesData)
	VisitStringData(d *StringData)
}

type DataCounter struct {
	c int
}

func NewDataCounter() *DataCounter {
	return &DataCounter{0}
}

func (c *DataCounter) VisitBytesData(d *BytesData){
	c.c += len(d.Bytes)
}

func (c *DataCounter) VisitStringData(d *StringData){
	c.c += len([]byte(d.String))
}

func (c *DataCounter) GetResult() int {
	return c.c
}

type Data interface {
	Print()
	Accept(v Visitor)
}

type BytesData struct {
	Bytes []byte
}

func (d *BytesData) Print() {
	fmt.Println(string(d.Bytes))
}

func (d *BytesData) Accept(v Visitor) {
	v.VisitBytesData(d)
}

type StringData struct {
	String string
}

func (d *StringData) Print() {
	fmt.Println(d.String)
}

func (d *StringData) Accept(v Visitor) {
	v.VisitStringData(d)
}