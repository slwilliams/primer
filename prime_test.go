package main

import (
	"testing"
)

func TestDefaultPrimeCalc(t *testing.T) {
	r := getPrime(defaultMaxNumber)

	if defaultMaxNumber != r.Max {
		t.Fatalf("Max value not equal to max constant (%d) %d ", defaultMaxNumber, r.Max)
	}
}

func TestArgPrimeCalc(t *testing.T) {
	for _, n := range []int{90000, 97} {
		r := getPrime(n)

		if n != r.Max {
			t.Fatalf("Max value not equal to max constant. Expected: %d; Got: %d", n, r.Max)
		}

		if r.Value > r.Max {
			t.Fatalf("Calculated prime higher than the max %d; Got: %d", r.Max, r.Value)
		}
	}
}
