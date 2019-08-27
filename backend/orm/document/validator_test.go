package document

import (
	"testing"
)

func TestGetAllValidator(t *testing.T) {
	validatorList, err := Validator{}.GetAllValidator()
	if err != nil {
		t.Error(err)
	}

	for k, v := range validatorList {
		t.Logf("idx: %v  v: %v \n", k, v)
	}

}

func TestValidatorGetCandidatesTopN(t *testing.T) {

	validators, power, upTimeMap, err := Validator{}.GetCandidatesTopN()

	if err != nil {
		t.Error(err)
	}

	t.Logf("power: %v \n", power)

	t.Logf("validators------------")
	for k, v := range validators {
		t.Logf("k: %v  v: %v \n", k, v)
	}

	t.Log("uptime map-------------")

	for k, v := range upTimeMap {
		t.Logf("k: %v  v: %v \n", k, v)
	}

}

func TestValidator_QueryValidatorMonikerByAddrArr(t *testing.T) {
	valaddr := []string{"iva1r2pq5y674llvd654tr9lm7s68wnumd0pf04v99"}
	data, err := Validator{}.QueryValidatorMonikerByAddrArr(valaddr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
