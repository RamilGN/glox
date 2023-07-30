package main

import (
	"reflect"
	"testing"
)

func assertNil(t *testing.T, val any) {
	t.Helper()

	if val != nil {
		t.Errorf("value: %v is not nil", val)
	}
}

func requireNil(t *testing.T, val any) {
	t.Helper()

	if val != nil {
		t.Fatalf("value: %v is not nil", val)
	}
}

func requireNotNil(t *testing.T, val any) {
	t.Helper()

	if val == nil {
		t.Fatalf("value: %v is nil", val)
	}
}

func assertNotNil(t *testing.T, val any) {
	t.Helper()

	if val == nil {
		t.Errorf("value: %v is nil", val)
	}
}

// TODO: generics.
func assertEqualString(t *testing.T, expected, actual string) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected: %q, actual: %q", expected, actual)
	}
}

// TODO: generics.
func assertEqualSlice(t *testing.T, expected, actual any) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, actual %v", expected, actual)
	}
}

func assertType(t *testing.T, expected, actual any) {
	t.Helper()

	expectedType := reflect.TypeOf(expected)

	actualType := reflect.TypeOf(actual)
	if expectedType != actualType {
		t.Errorf("expected type: %v, actual type: %v", expectedType, actualType)
	}
}

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()

	if expected != actual {
		t.Errorf("expected %+v, actual %+v", expected, actual)
	}
}
