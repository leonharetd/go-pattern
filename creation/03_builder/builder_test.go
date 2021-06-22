package builder

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	c := NewBuildContract(InsuranceContract{}).
		SetContract("Cont").
		SetPersonName("Bob").
		SetCompanyName("A").
		SetBeginDate(123).SetEndDate(456).Build()
	c.SomeOptions()

	c1 := NewBuildContract(InsuranceContract{}).
		SetContract("Cont").
		SetPersonName("Alice").
		SetPersonName("Bob").
		SetCompanyName("A").
		SetBeginDate(123).SetEndDate(456).Build()
	c1.SomeOptions()
}
