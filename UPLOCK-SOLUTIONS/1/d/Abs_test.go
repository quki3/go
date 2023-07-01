package main

import (
    "math"
    "testing"
)

func TestXxx(t *testing.T) {
    num := -5.6
    expected := 5.6

    result := Xxx(num)

    if result != expected {
        t.Errorf("Expected %f, but got %f", expected, result)
    }
}
