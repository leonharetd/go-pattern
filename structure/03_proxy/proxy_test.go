package proxy

import (
	"fmt"
	"testing"
)

func TestOrder(t *testing.T) {
    order := OrderProxy{realOrder: Order{
		orderName: "book",
		orderUser: "Bob",
		productName: "book",
	}}

	fmt.Printf("%v\n", order.realOrder)
	order.setOrderName("book2", "Bob")
	fmt.Printf("%v\n", order.realOrder)

	order.setOrderName("book2", "Alice")
}