package alteryx_formulas

import "strings"

func (calc *calculator) equal() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 == value2)
}

func (calc *calculator) notEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	calc.pushValue(value1 != value2)
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

func (calc *calculator) numberGreaterEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) >= value2.(float64))
}

func (calc *calculator) numberLessThan() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) < value2.(float64))
}

func (calc *calculator) numberLessEqual() {
	value1 := calc.popValue()
	value2 := calc.popValue()
	if value1 == nil || value2 == nil {
		calc.pushValue(false)
		return
	}
	calc.pushValue(value1.(float64) <= value2.(float64))
}

func (calc *calculator) contains() {
	value := calc.popValue().(string)
	lookFor := calc.popValue().(string)
	caseInsensitive := calc.popValue().(float64)
	if caseInsensitive != 0 {
		value = strings.ToLower(value)
		lookFor = strings.ToLower(lookFor)
	}
	contains := strings.Contains(value, lookFor)
	calc.pushValue(contains)
}

func (calc *calculator) endsWith() {
	value := calc.popValue().(string)
	lookFor := calc.popValue().(string)
	caseInsensitive := calc.popValue().(float64)
	startAt := len(value) - len(lookFor)
	if startAt < 0 {
		calc.pushValue(false)
		return
	}
	if caseInsensitive != 0 {
		value = strings.ToLower(value)
		lookFor = strings.ToLower(lookFor)
	}

	contains := value[startAt:] == lookFor
	calc.pushValue(contains)
}
