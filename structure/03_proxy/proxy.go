package proxy

type Order struct {
    productName string
	orderName string
	orderUser string
}

func (o *Order) getProductName() string {
	return o.productName
}

func (o *Order) setProductName(productName string) {
	o.productName = productName
}

func (o *Order) setOrderName(orderName string) {
	o.orderName = orderName
}

func (o *Order) getOrderUser() string {
	return o.orderUser
}


type OrderProxy struct {
    realOrder Order
}

func (o *OrderProxy) getProductName() string {
	return o.realOrder.productName
}

func (o *OrderProxy) setProductName(productName string, orderUser string) {
	if o.realOrder.orderUser != orderUser {
		panic("productName not allowed")
	}
	o.realOrder.productName = productName
}

func (o *OrderProxy) setOrderName(orderName string, orderUser string) {
	if o.realOrder.orderUser != orderUser {
		panic("productName not allowed")
	}
	o.realOrder.orderName = orderName
}

func (o *OrderProxy) getOrderUser() string {
	return o.realOrder.orderUser
}