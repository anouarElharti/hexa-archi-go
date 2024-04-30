package api

import (
	"github.com/anouarElharti/hexa-archi-go/internal/ports"
)

type Adapter struct {
	arith ports.ArithmeticPort
}

// RETURN POINTER
func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apiA Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apiA.arith.Addition(a, b)

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apiA Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apiA.arith.Subtraction(a, b)

	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiA Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apiA.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apiA Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apiA.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	return answer, nil
}
