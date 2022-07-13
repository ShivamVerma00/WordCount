package main

import (
	"testing"
)

func TestWord(t *testing.T) {
	str := "This is Demo"
	expected := "Succes"

	actual := word(str)

	if actual != expected {
		t.Errorf("Failed...")
	}
}
