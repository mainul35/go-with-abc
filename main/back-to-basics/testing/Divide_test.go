package testing

import (
	"log"
	"testing"
)

func TestDivide(t *testing.T) {
	res, err := Divide(10.0, 2.0)
	if err != nil {
		t.Errorf("Threw an error when it shouldn't: %s", err)
	}
	log.Println(res)
}

func TestBadDivide(t *testing.T) {
	res, err := Divide(10.0, 0.0)

	if err == nil {
		t.Errorf("Expected error, but no error thrown")
	}
	log.Println(res)
}
