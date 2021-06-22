package builder

import "fmt"

type InsuranceContract struct {
	Contract    string
	PersonName  string
	CompanyName string
	BeginDate   int64
	EndDate     int64
}

func (c InsuranceContract) SomeOptions(){
	fmt.Printf("%s %s %s", c.Contract, c.PersonName, c.CompanyName)
}

type BuildContract struct {
	contract InsuranceContract
}

func NewBuildContract(contract InsuranceContract) *BuildContract {
	return &BuildContract{contract: contract}
}
func (b *BuildContract) SetContract(c string) *BuildContract{
	b.contract.Contract = c
	return b
}

func (b *BuildContract) SetPersonName(personName string) *BuildContract{
	if b.contract.PersonName != "" {
        panic("personName is only one")
	}
	b.contract.PersonName = personName
	return b

}

func (b *BuildContract) SetCompanyName(cm string) *BuildContract{
	b.contract.CompanyName = cm
	return b
}

func (b *BuildContract) SetBeginDate(date int64) *BuildContract{
	b.contract.BeginDate = date
	return b
}

func (b *BuildContract) SetEndDate(date int64) *BuildContract{
	b.contract.EndDate = date
	return b
}

func (b *BuildContract) Build() InsuranceContract{
	return b.contract
}