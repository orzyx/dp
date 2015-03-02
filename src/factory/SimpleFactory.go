package factory

type Operation interface {
	GetResult() float64
	SetNumA(float64)
	SetNumB(float64)
}

type BaseOperation struct {
	numberA float64
	numberB float64
}

func (operation *BaseOperation) SetNumA(numA float64) {
	operation.numberA = numA
}

func (operation *BaseOperation) SetNumB(numB float64) {
	operation.numberB = numB
}

type OperationAdd struct {
	BaseOperation
}

func (this *OperationAdd) GetResult() float64 {
	return this.numberA + this.numberB
}

type OperationSub struct {
	BaseOperation
}

func (this *OperationSub) GetResult() float64 {
	return this.numberA - this.numberB
}

type OperationMul struct {
	BaseOperation
}

func (this *OperationMul) GetResult() float64 {
	return this.numberA * this.numberB
}

type OperationDiv struct {
	BaseOperation
}

func (this *OperationDiv) GetResult() float64 {
	if this.numberB == 0 {
		panic("divide by zero!!")
	}

	return this.numberA / this.numberB
}

type OperationFactory struct {
}

func (this *OperationFactory) CreateOperation(operator string) (op Operation) {
	switch operator {
	case "+":
		op = new(OperationAdd)
	case "-":
		op = new(OperationSub)
	case "/":
		op = new(OperationDiv)
	case "*":
		op = new(OperationMul)
	default:
		panic("运算符号错误！")
	}
	return
}
