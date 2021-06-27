package adapter

import "fmt"

type IPlugin interface {
	GetPin() int
}

type TwoPlugin struct {
	Pin int
}

func (t *TwoPlugin) GetPin() int {
	return t.Pin
}

type ThreePlugin struct {
	Pin int
}

func (t *ThreePlugin) GetPin() int {
	return t.Pin
}

type TwoPower struct {
	plugin IPlugin
}

func (t *TwoPower) SetPlug(plug IPlugin) bool {
	if plug.GetPin() != 2 {
		fmt.Println("TwoPower need TwoPlugin", "get", plug.GetPin(), "plugin")
		return false
	}
	t.plugin = plug
	fmt.Println("TwoPower uesd plugin", plug.GetPin())
	return true
}

type Adapter struct{}

func (ad *Adapter) ThreeToTwo(plugin ThreePlugin) IPlugin {
	return &TwoPlugin{Pin: plugin.Pin - 1}
}
