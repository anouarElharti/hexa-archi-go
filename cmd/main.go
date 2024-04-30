package main

import (
	"github.com/anouarElharti/hexa-archi-go/internal/adapters/core/arithmetic"
	"github.com/anouarElharti/hexa-archi-go/internal/ports"
)

func main() {
	// ports
	var core ports.ArithmeticPort
	core = arithmetic.NewAdapter()

}
