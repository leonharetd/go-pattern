package objectpool

import "testing"

func TestObjectPool(t *testing.T) {
	pool := NewProxyPool(3)
	p1 := <- pool
	p1.Show()

	p2 := <- pool
	p2.Show()

	p3 := <- pool
	p3.Show()

}