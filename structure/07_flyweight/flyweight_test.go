package flyweight

import (
	"fmt"
	"testing"
)

func TestFlyWeight(t *testing.T) {
    dc := DeliverCompany{}

	dc.Hire("bob")
	dc.Hire("Alice")

	deliver1 := dc.GetDelive("bob")
	deliver2 := dc.GetDelive("bob")

	deliver1.DeliverPackets([]string{"box1", "box2"})

	fmt.Println(deliver1 == deliver2)

}
