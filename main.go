package main

import (
	sf "./src/factory"
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var opFactory sf.OperationFactory
	op := opFactory.CreateOperation("/")
	op.SetNumA(3.0)
	op.SetNumB(0.0)
	fmt.Println(op.GetResult())
}
