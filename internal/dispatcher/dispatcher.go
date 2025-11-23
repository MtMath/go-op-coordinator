package dispatcher

import (
	"fmt"

	"notask/op-coordinator/internal/clients"
)

type Dispatcher struct {
    Add *clients.AddClient
    Sub *clients.SubClient
    Mul *clients.MulClient
    Div *clients.DivClient
}

func NewDispatcher(addAddr, subAddr, mulAddr, divAddr string) *Dispatcher {
    return &Dispatcher{
        Add: clients.NewAddClient(addAddr),
        Sub: clients.NewSubClient(subAddr),
        Mul: clients.NewMulClient(mulAddr),
        Div: clients.NewDivClient(divAddr),
    }
}

func (d *Dispatcher) Dispatch(op string, a, b float64) (float64, error) {
	switch op {

	case "+":
		return d.Add.Compute(a, b)

	case "-":
		return d.Sub.Compute(a, b)

	case "*":
		return d.Mul.Compute(a, b)

	case "/":
		return d.Div.Compute(a, b)
	}

	return 0, fmt.Errorf("unknown operator in dispatcher: %s", op)
}
