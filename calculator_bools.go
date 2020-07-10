package alteryx_formulas

func (calc *calculator) equal() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 == value2)
}

func (calc *calculator) numberGreaterThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) > value2.(float64))
}
