package decorator

type Component interface {
	Cacl(user string) float64
}

type ConcreteDecorator struct {}
func (c *ConcreteDecorator) Cacl(user string) float64 {
	return 0.0
}


type MonthDecorator struct {
	component Component
}

func (m *MonthDecorator) Cacl(user string) float64 {
	c := m.component.Cacl(user)
	c = c + 10000
	return c
}

type GroupDecorator struct {
	component Component
}

func (g *GroupDecorator) Cacl(user string) float64 {
	c := g.component.Cacl(user)
	c = c + 8000
	return c
}



