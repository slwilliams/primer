package main

import (
	"testing"
)

func TestUUID(t *testing.T) {
	testID := getUUIDv4()
	if testID == "" || len(testID) < 10 {
		t.Fatal("Invalid ID generated")
	}
}
