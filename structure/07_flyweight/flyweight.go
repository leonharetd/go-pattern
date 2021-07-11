package flyweight

import "fmt"

type IDeliverCompany interface{
	Hire(name string)
	DeliveTask(name string, packages []string)
	GetDelive(name string) IDeliver
}

//DeliverCompany 快递公司
type DeliverCompany struct {
	Employees map[string]IDeliver
}

func (dc *DeliverCompany) Hire(name string) {
	if dc.Employees == nil {
		dc.Employees = make(map[string]IDeliver)
	}

	if _, ok := dc.Employees[name]; !ok {
        dc.Employees[name] = &Deliver{Name: name}
	}
}

func (dc *DeliverCompany) GetDelive(name string) IDeliver{
	if delive, ok := dc.Employees[name]; ok {
		return delive
	}
	return nil
}

func (dc *DeliverCompany) DeliveTask(name string, packages []string) {
	dc.Employees[name].DeliverPackets(packages)
}

//IDeliver 快递员能做的事情
type IDeliver interface {
	//送货
	DeliverPackets(packets []string)
}

//Deliver 快递员工,员工是一个享元对象
type Deliver struct {
	Name    string   //快递员的名字
	Packets []string //快递员的携带的快递，这一部分是变化的
}

func (d *Deliver) DeliverPackets(packets []string) {
    fmt.Println(d.Name, "delived", packets)
}