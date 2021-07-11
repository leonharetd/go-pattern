package component

import "testing"

func TestComponent(t *testing.T) {

	root := Componsite{Name: "服装"}
	c1 := Componsite{Name: "男装"}
	c2 := Componsite{Name: "女装"}

	leaf1 := Componsite{Name: "衬衣"}
	leaf2 := Componsite{Name: "夹克"}
	leaf3 := Componsite{Name: "裙子"}
	leaf4 := Componsite{Name: "套装"}

	root.addChild(&c1)
	root.addChild(&c2)
	c1.addChild(&leaf1)
	c1.addChild(&leaf2)
	c2.addChild(&leaf3)
	c2.addChild(&leaf4)

	root.printStruct(" ")
}
