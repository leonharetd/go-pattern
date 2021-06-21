package new

import (
	"fmt"
	"testing"
)

type Car interface {
    Speaker()
}

type WulingHongguang struct {
	Name string
}

func (wh *WulingHongguang) Speaker() {
	fmt.Printf("%s speaker\n", wh.Name)
}

func newWulingHongguang() *WulingHongguang {
	return &WulingHongguang{
		Name: "五菱宏光",
	}
}

func TestNew(t *testing.T) {
    car := newWulingHongguang()
	car.Speaker()
}