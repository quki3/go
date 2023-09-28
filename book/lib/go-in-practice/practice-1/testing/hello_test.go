package main

import "testing"

func TestName(t *testing.T) {			// function starting with t are run
	name := getName()

	if name != "World!" {						// report err
		t.Error("Response from getName is unexpected value")
	}
}
