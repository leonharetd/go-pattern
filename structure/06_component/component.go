package component

import "fmt"

type Component interface {
	addChild(child Component)
	removeChild(child Component)
	getChild(index int) Component
	printStruct(pre string)
	setParent(child Component)
}

type Leaf struct {
	Name string
	parent Component
}

// addChild add an child
func (l *Leaf) addChild(child Component) {
    fmt.Println("Leaf do not add Child")
}

func (l *Leaf) removeChild(child Component) {
	fmt.Println("Leaf do not remove Child")
}

func (l *Leaf) getChild(index int) Component {
	return nil
}

func (l *Leaf) printStruct(pre string) {
	fmt.Println(pre, "-", l.Name)
}

func (l *Leaf) setParent(parent Component) {
	l.parent = parent
}

type Componsite struct {
	Name string
	parent Component
	childComponsites []Component
}

func (c *Componsite) addChild(child Component) {
    c.childComponsites = append(c.childComponsites, child)
	child.setParent(child)
}

func (c *Componsite) removeChild(child Component) {
	fmt.Println("componsite do not remove child")
}

func (c *Componsite) getChild(index int) Component {
	if index >= len(c.childComponsites) {
		fmt.Println(index, "child not exist")
		return nil
	}
	return c.childComponsites[index]
}

func (c *Componsite) setParent(parent Component) {
	c.parent = parent
}

func (c *Componsite) printStruct(pre string) {
	fmt.Println(pre, "-", c.Name)
	for _, comp := range c.childComponsites {
		comp.printStruct(pre+" ")
	}
}