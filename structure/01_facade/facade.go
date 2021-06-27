package facade

import "fmt"

type Mode interface {
	Test()
}

type ModeA struct {}
func (a *ModeA) Test() {
	fmt.Println("modeA")
}

type ModeB struct {}
func (b *ModeB) Test() {
	fmt.Println("modeB")
}

type ModeC struct {}
func (c *ModeC) Test() {
	fmt.Println("modeC")
}

type Facade struct {}
func (f *Facade) Test() {
	a := ModeA{}
	a.Test()

	b := ModeB{}
	b.Test()

	c := ModeC{}
	c.Test()
}