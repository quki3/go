package main

import (
    "math"
    "testing"
)

func TestAbs(t *testing.T) {
    num := -5.6
    expected := 5.6

    result := abs(num)

    if result != expected {
        t.Errorf("Expected %f, but got %f", expected, result)
    }
}
