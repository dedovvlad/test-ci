package main

import "testing"

func TestHello(t *testing.T) {
	ph := PrintHello()
	dataTest := "Hello"

	if ph != dataTest {
		t.Errorf("Error: %s != %s", ph, dataTest)
	}
}
