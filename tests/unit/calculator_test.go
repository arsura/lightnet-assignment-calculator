package unit_test

import (
	"github.com/arsura/lightnet-assignment-calculator/controllers"
	"testing"
)

func TestSum1(t *testing.T) {
	result := controllers.Sum(1, 2.2)
	if 3.2 != result {
		t.Errorf("1 + 2.2 = %f; want 3.2", result)
	}
}

func TestSum2(t *testing.T) {
	result := controllers.Sum(45687.12, 50000000.11)
	if 50045687.23 != result {
		t.Errorf("45687.12 + 50000000.11 = %f; want 50045687.23", result)
	}
}

func TestSub1(t *testing.T) {
	result := controllers.Sub(1, 2.2)
	if -1.2 != result {
		t.Errorf("1 - 2.2 = %f; want -1.2", result)
	}
}

func TestSub2(t *testing.T) {
	result := controllers.Sub(50000000.11, 45687.12)
	if 49954312.99 != result {
		t.Errorf("50000000.11 - 45687.12 = %f; want 49954312.99", result)
	}
}

func TestMul1(t *testing.T) {
	result := controllers.Mul(0, 10.32)
	if 0 != result {
		t.Errorf("0 * 10.32 = %f; want 0", result)
	}
}

func TestMul2(t *testing.T) {
	result := controllers.Mul(12345.21, 10523.2)
	if 129911113.872 != result {
		t.Errorf("12345.21 * 10523.2 = %f; want 129911113.872", result)
	}
}

func TestDiv1(t *testing.T) {
	result := controllers.Div(0, 1.05)
	if 0 != result {
		t.Errorf("0 / 1.05 = %f; want 0", result)
	}
}

func TestDiv2(t *testing.T) {
	result := controllers.Div(1.25, 0.25)
	if 5 != result {
		t.Errorf("1.25 / 0.25 = %f; want 5", result)
	}
}
