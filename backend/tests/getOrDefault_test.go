package tests

import (
	"api/pkg/tools"
	"testing"
)

func TestGetOrDefault(t *testing.T) {
	// Test case 1: GetOrDefault with string
	gotInt := tools.GetOrDefaultNumber("1232", 0).(int)
	wantInt := 1232
	if gotInt != wantInt {
		t.Errorf("got %v, wanted %v", gotInt, wantInt)
	}

	gotFloat := tools.GetOrDefaultNumber("1232.1", 0.0).(float64)
	wantFloat := 1232.1

	if gotFloat != wantFloat {
		t.Errorf("got %v, wanted %v", gotFloat, wantFloat)
	}
}
