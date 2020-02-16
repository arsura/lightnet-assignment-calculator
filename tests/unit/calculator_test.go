package unit

import (
	"github.com/arsura/lightnet-assignment-calculator/controllers"
	"github.com/arsura/lightnet-assignment-calculator/tests/helper"
	"testing"
	"math"
)

func TestSum(t *testing.T) {
	tables := []struct {
		a      float64
		b      float64
		result float64
	}{
		{1, 2.2, 3.2},
		{2, 2, 4},
		{5, 2, 7},
		{1, -1, 0},
		{45687.12, 50000000.11, 50045687.23},
	}

	for _, table := range tables {
		total := controllers.Sum(table.a, table.b)
		if !helper.FloatNearlyEqual(total, table.result, 0.00001) {
			t.Errorf("%f + %f, got: %f, want: %f.", table.a, table.b, total, table.result)
		}
	}
}

func TestSub(t *testing.T) {
	tables := []struct {
		a      float64
		b      float64
		result float64
	}{
		{1, 2.2, -1.2},
		{2, 2, 0},
		{-2, -2, 0},
		{50000000.11, 45687.12, 49954312.99},
	}

	for _, table := range tables {
		total := controllers.Sub(table.a, table.b)
		if !helper.FloatNearlyEqual(total, table.result, 0.00001) {
			t.Errorf("%f - %f, got: %f, want: %f.", table.a, table.b, total, table.result)
		}
	}
}

func TestMul(t *testing.T) {
	tables := []struct {
		a      float64
		b      float64
		result float64
	}{
		{1, 52.2, 52.2},
		{-2, -2, 4},
		{0, 2, 0},
		{-2, 0, 0},
		{500.11, 456.12, 228110.1732},
	}

	for _, table := range tables {
		total := controllers.Mul(table.a, table.b)
		if !helper.FloatNearlyEqual(total, table.result, 0.00001) {
			t.Errorf("%f * %f, got: %f, want: %f.", table.a, table.b, total, table.result)
		}
	}
}

func TestDiv(t *testing.T) {
	tables := []struct {
		a      float64
		b      float64
		result float64
	}{
		{52.2, 52.2, 1},
		{-2, 0, math.Inf(-1)},
		{0, 2, 0},
		{-2, 2, -1},
		{5235.10, 456.12, 11.4774621},
	}

	for _, table := range tables {
		total := controllers.Div(table.a, table.b)
		if !helper.FloatNearlyEqual(total, table.result, 0.00001) {
			t.Errorf("%f / %f, got: %f, want: %f.", table.a, table.b, total, table.result)
		}
	}
}
