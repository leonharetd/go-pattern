package samplefactory

import "fmt"

type Car interface {
	Speaker()
}

type CarA struct {
	Name string
}

func(a *CarA) Speaker() {
	fmt.Printf("%s speaker", a.Name)
}

type CarB struct {
	Name string
}

func(b *CarB) Speaker() {
	fmt.Printf("%s speaker", b.Name)
}

func newCarA(name string) *CarA {
	return &CarA{
		Name: name,
	}
}

func newCarB(name string) *CarB {
	return &CarB{
		Name: name,
	}
}

func CarFactory(name, carType string) Car{
	if carType == "A" {
		return newCarA(name)
	} else if carType == "B" {
		return newCarB(name)
	}
	return nil
}


