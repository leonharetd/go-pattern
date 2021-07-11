package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	base := ConcreteDecorator{}

	month := MonthDecorator{&base}
	group := GroupDecorator{&month}
	g := MonthDecorator(month) // ?
	fmt.Println(g, month)
	fmt.Println(month.Cacl("张三"))
	fmt.Println(group.Cacl("李四"))
}
