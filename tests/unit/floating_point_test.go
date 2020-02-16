package unit

import (
	"github.com/arsura/lightnet-assignment-calculator/tests/helper"
	"testing"
)

func TestFloatNearlyEqual(t *testing.T) {
	tables := []struct {
		a      float64
		b      float64
		result bool
	}{
		{2, 2, true},
		{-2, 2, false},
		{2.00001, 2.00010, false},
		{0.000011, 0.000012, false},
		{0.000000120, 0.000000119, false},
	}

	for _, table := range tables {
		isEqual := helper.FloatNearlyEqual(table.a, table.b, 0.00001)
		if isEqual != table.result {
			t.Errorf("%f == %f, got: %v, want: %v.", table.a, table.b, isEqual, table.result)
		}
	}
}
