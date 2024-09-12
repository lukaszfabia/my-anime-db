package tests

import (
	"api/internal/controller"
	"testing"
)

func TestGetOrDefault(t *testing.T) {
	result := controller.GetOrDefault("", 10)
	if result != 10 {
		t.Errorf("Expected result to be 10, but got %v", result)
	}

	result = controller.GetOrDefault("123", 10)
	if result != 123 {
		t.Errorf("Expected result to be 123, but got %v", result)
	}

	result = controller.GetOrDefault("", 3.14)
	if result != 3.14 {
		t.Errorf("Expected result to be 3.14, but got %v", result)
	}

	result = controller.GetOrDefault("3.14", 2.71)
	if result != 3.14 {
		t.Errorf("Expected result to be 3.14, but got %v", result)
	}

	result = controller.GetOrDefault("", "default")
	if result != "default" {
		t.Errorf("Expected result to be 'default', but got %v", result)
	}

	result = controller.GetOrDefault("hello", "default")
	if result != "hello" {
		t.Errorf("Expected result to be 'hello', but got %v", result)
	}
}
