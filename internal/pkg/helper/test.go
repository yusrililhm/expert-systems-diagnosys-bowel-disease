package helper

import (
	"testing"
)

func Equals(t *testing.T, expected any, actual any) {
	if expected != actual {
		t.Fatalf("Not equal. Expected %s actual %s", expected, actual)
	}
}

func Nil(t *testing.T, data any) {
	if data == nil {
		t.Fatal("Expected nil, but got not nil")
	}
}

func NotNil(t *testing.T, data any) {
	if data != nil {
		t.Fatal("Expected not nil, but got nil")
	}
}
