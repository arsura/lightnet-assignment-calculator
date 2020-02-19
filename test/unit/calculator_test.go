package unit

import (
	"bytes"
	"encoding/json"
	"github.com/arsura/lightnet-assignment-calculator/controller"
	"github.com/arsura/lightnet-assignment-calculator/router"
	"github.com/arsura/lightnet-assignment-calculator/test/helper"
	"github.com/arsura/lightnet-assignment-calculator/util"
	"math"
	"net/http"
	"testing"
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
		total := controller.Sum(table.a, table.b)
		if !util.FloatNearlyEqual(total, table.result, 0.00001) {
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
		total := controller.Sub(table.a, table.b)
		if !util.FloatNearlyEqual(total, table.result, 0.00001) {
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
		total := controller.Mul(table.a, table.b)
		if !util.FloatNearlyEqual(total, table.result, 0.00001) {
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
		total := controller.Div(table.a, table.b)
		if !util.FloatNearlyEqual(total, table.result, 0.00001) {
			t.Errorf("%f / %f, got: %f, want: %f.", table.a, table.b, total, table.result)
		}
	}
}

func TestHTTPSumCode200(t *testing.T) {
	t.Run("it should return httpCode 200OK and result is 3.0", func(t *testing.T) {

		expectedResult := 3.0
		data := map[string]float64{
			"a": 1.0,
			"b": 2.0,
		}
		bytesRepresentation, _ := json.Marshal(data)

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.sum", bytes.NewBuffer(bytesRepresentation))

		var resp map[string]float64
		json.NewDecoder(w.Body).Decode(&resp)

		if !util.FloatNearlyEqual(resp["result"], expectedResult, 0.00001) {
			t.Errorf("%f + %f, got: %f, want: %f", data["a"], data["b"], resp["result"], expectedResult)
		}

		if status := w.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})
}

func TestHTTPSumCode400(t *testing.T) {
	t.Run("it should return httpCode 400BadRequest", func(t *testing.T) {

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.sum", nil)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("wrong code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

func TestHTTPSubCode200(t *testing.T) {
	t.Run("it should return httpCode 200OK and result is 1.2", func(t *testing.T) {

		expectedResult := 1.2
		data := map[string]float64{
			"a": 3.2,
			"b": 2.0,
		}
		bytesRepresentation, _ := json.Marshal(data)

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.sub", bytes.NewBuffer(bytesRepresentation))

		var resp map[string]float64
		json.NewDecoder(w.Body).Decode(&resp)

		if !util.FloatNearlyEqual(resp["result"], expectedResult, 0.00001) {
			t.Errorf("%f - %f, got: %f, want: %f", data["a"], data["b"], resp["result"], expectedResult)
		}

		if status := w.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})
}

func TestHTTPSubCode400(t *testing.T) {
	t.Run("it should return httpCode 400BadRequest", func(t *testing.T) {

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.sub", nil)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("wrong code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

func TestHTTPMulCode200(t *testing.T) {
	t.Run("it should return httpCode 200OK and result is 0", func(t *testing.T) {

		expectedResult := 0.0
		data := map[string]float64{
			"a": 50.22,
			"b": 0,
		}
		bytesRepresentation, _ := json.Marshal(data)

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.mul", bytes.NewBuffer(bytesRepresentation))

		var resp map[string]float64
		json.NewDecoder(w.Body).Decode(&resp)

		if !util.FloatNearlyEqual(resp["result"], expectedResult, 0.00001) {
			t.Errorf("%f * %f, got: %f, want: %f", data["a"], data["b"], resp["result"], expectedResult)
		}

		if status := w.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})
}

func TestHTTPSMulCode400(t *testing.T) {
	t.Run("it should return httpCode 400BadRequest", func(t *testing.T) {

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.mul", nil)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("wrong code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

func TestHTTPDivCode200(t *testing.T) {
	t.Run("it should return httpCode 200OK and result is 1.0", func(t *testing.T) {

		expectedResult := 1.0
		data := map[string]float64{
			"a": 50.22,
			"b": 50.22,
		}
		bytesRepresentation, _ := json.Marshal(data)

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.div", bytes.NewBuffer(bytesRepresentation))

		var resp map[string]float64
		json.NewDecoder(w.Body).Decode(&resp)

		if !util.FloatNearlyEqual(resp["result"], expectedResult, 0.00001) {
			t.Errorf("%f * %f, got: %f, want: %f", data["a"], data["b"], resp["result"], expectedResult)
		}

		if status := w.Code; status != http.StatusOK {
			t.Errorf("wrong code: got %v want %v", status, http.StatusOK)
		}
	})
}

func TestHTTPDivDivisorIsZeroCode400(t *testing.T) {
	t.Run("it should return httpCode 400BadRequest", func(t *testing.T) {

		data := map[string]float64{
			"a": 50.22,
			"b": 0,
		}
		bytesRepresentation, _ := json.Marshal(data)

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.div", bytes.NewBuffer(bytesRepresentation))

		var resp map[string]float64
		json.NewDecoder(w.Body).Decode(&resp)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("wrong code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}

func TestHTTPSDivCode400(t *testing.T) {
	t.Run("it should return httpCode 400BadRequest", func(t *testing.T) {

		router := router.Router()
		w := helper.PerformRequest(router, http.MethodPost, "/calculator.div", nil)

		if status := w.Code; status != http.StatusBadRequest {
			t.Errorf("wrong code: got %v want %v", status, http.StatusBadRequest)
		}
	})
}
