package alteryx_formulas_test

import (
	"math"
	"testing"
)
import f "github.com/tlarsen7572/alteryx_formulas"

func TestIfNull(t *testing.T) {
	result, errs := f.Calculate(`IF NULL() THEN 1 ELSE 2 ENDIF`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 2.0 {
		t.Fatalf(`expected 2.0 but got %v`, result)
	}
}

func TestAbsNull(t *testing.T) {
	result, errs := f.Calculate(`abs(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAcosNull(t *testing.T) {
	result, errs := f.Calculate(`acos(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAsinNull(t *testing.T) {
	result, errs := f.Calculate(`asin(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAtanNull(t *testing.T) {
	result, errs := f.Calculate(`atan(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAtan2Null(t *testing.T) {
	result, errs := f.Calculate(`atan2(1, NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if expected := math.Atan2(1, 0); result != expected {
		t.Fatalf(`expected %v but got %v`, expected, result)
	}

	result, errs = f.Calculate(`atan2(NULL(), 1)`, nil)
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestAverageNull(t *testing.T) {
	result, errs := f.Calculate(`average(20, 30, 55, NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 26.25 {
		t.Fatalf(`expected 26.25 but got %v`, result)
	}
}

func TestCeilNull(t *testing.T) {
	result, errs := f.Calculate(`ceil(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestCharFromIntNull(t *testing.T) {
	result, errs := f.Calculate(`charfromint(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestCharToIntNull(t *testing.T) {
	result, errs := f.Calculate(`chartoint(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestContainsNull(t *testing.T) {
	result, errs := f.Calculate(`contains(NULL(), '1')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, errs = f.Calculate(`contains('1', NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, errs = f.Calculate(`contains('123', '1', NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestCosNull(t *testing.T) {
	result, errs := f.Calculate(`cos(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestCoshNull(t *testing.T) {
	result, errs := f.Calculate(`cosh(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestCountWordsNull(t *testing.T) {
	result, errs := f.Calculate(`countwords(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 0.0 {
		t.Fatalf(`expected 0 but got %v`, result)
	}
}

func TestDistanceNull(t *testing.T) {
	result, errs := f.Calculate(`DISTANCE(null(), NULL(), NULL(), NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 0.0 {
		t.Fatalf(`expected 0 but got %v`, result)
	}
}

func TestEndsWithNull(t *testing.T) {
	result, errs := f.Calculate(`endsWith(NULL(), '1')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, errs = f.Calculate(`endsWith('1', NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, errs = f.Calculate(`endsWith('123', '3', NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestExpNull(t *testing.T) {
	result, errs := f.Calculate(`exp(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestFindStringNull(t *testing.T) {
	result, errs := f.Calculate(`findstring(NULL(),'a')`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != false {
		t.Fatalf(`expected false but got %v`, result)
	}

	result, errs = f.Calculate(`findstring('a',null())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != true {
		t.Fatalf(`expected true but got %v`, result)
	}
}

func TestFloorNull(t *testing.T) {
	result, errs := f.Calculate(`floor(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestGetWordNull(t *testing.T) {
	result, errs := f.Calculate(`getword('abc def', NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `abc` {
		t.Fatalf(`expected abc but got %v`, result)
	}

	result, errs = f.Calculate(`getword(NULL(), 0)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestHexToNumberNull(t *testing.T) {
	result, errs := f.Calculate(`hexToNumber(NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}
}

func TestIifNull(t *testing.T) {
	result, errs := f.Calculate(`iif(null(),10,20)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 20.0 {
		t.Fatalf(`expected 20 but got %v`, result)
	}
}

func TestLeftNull(t *testing.T) {
	result, errs := f.Calculate(`left(NULL(),1)`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != nil {
		t.Fatalf(`expected nil but got %v`, result)
	}

	result, errs = f.Calculate(`left('abc',NULL())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != `` {
		t.Fatalf(`expected blank string but got %v`, result)
	}
}

func TestLengthNull(t *testing.T) {
	result, errs := f.Calculate(`length(null())`, nil)
	if len(errs) > 0 {
		t.Fatalf(`expected no errors but got: %v`, errs)
	}
	if result != 0.0 {
		t.Fatalf(`expected 0 but got %v`, result)
	}
}
