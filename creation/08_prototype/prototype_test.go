package prototype

import (
	"fmt"
	"testing"
)

func TestPrototype(t *testing.T) {
	p := CloneLab{Lab: make(map[string]string)}
	p.set("dfd", "dsfs")
	fmt.Println(p.Lab)

	p1 := p.Clone()
	p1.set("dfd", "clone from p")
	fmt.Println(p1.Lab)

}
