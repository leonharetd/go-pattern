package adapter

import "testing"

func TestAdapter(t *testing.T) {

	twoP := TwoPlugin{Pin: 2}
    threeP := ThreePlugin{Pin: 3}
	adapter := Adapter{}
	twoPower := TwoPower{}

	twoPower.SetPlug(&twoP)
	twoPower.SetPlug(&threeP)
	twoPower.SetPlug(adapter.ThreeToTwo(threeP))
}