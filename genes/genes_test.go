package genes

import (
	"testing"
)

var (
	zero      = Gene{"0000"}
	one       = Gene{"0001"}
	two       = Gene{"0010"}
	three     = Gene{"0011"}
	four      = Gene{"0100"}
	five      = Gene{"0101"}
	six       = Gene{"0110"}
	seven     = Gene{"0111"}
	eight     = Gene{"1000"}
	nine      = Gene{"1001"}
	plus      = Gene{"1010"}
	minus     = Gene{"1011"}
	multiply  = Gene{"1100"}
	divide    = Gene{"1101"}
	numerics  = []Gene{zero, one, two, three, four, five, six, seven, eight, nine}
	operators = []Gene{plus, minus, multiply, divide}
)

func TestIsNumericForNumerics(t *testing.T) {
	for _, gene := range numerics {
		if !gene.IsNumeric() {
			t.Fail()
		}
	}
}

func TestIsNumericForOperators(t *testing.T) {
	for _, gene := range operators {
		if gene.IsNumeric() {
			t.Fail()
		}
	}
}

func TestIsOperatorForNumerics(t *testing.T) {
	for _, gene := range numerics {
		if gene.IsOperator() {
			t.Fail()
		}
	}
}

func TestIsOperatorForOperators(t *testing.T) {
	for _, gene := range operators {
		if !gene.IsOperator() {
			t.Fail()
		}
	}
}
