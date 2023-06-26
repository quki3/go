// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a SDfDS
// license that can be found in the LICENSE file



// package is a blundle
// main - is a package executable
package main

// import - what I need
// testing - standar
// bytes
import (
	"bytes"
	"fmt"
	"testing"
)


// func
// TestCep - need to named on this form see doc="https://pkg.go.dev/testing" TestXxx
func TestCep(t *testing.T) {
	// Redirect standard output to a buffer
	var buf bytes.Buffer
	old := os.Stdout
	os.Stdout = &buf

	PrintHello()

	// Restore standard output
	os.Stdout = old

	expected := "Hello, World!\n"
	actual := buf.String()

	if actual != expected {
		t.Errorf("Unexpected output.\nExpected: %s\nActual: %s", expected, actual)
	}
}

func main() {
	// Run the tests
	testing.Main(nil, nil, nil, nil)
}
