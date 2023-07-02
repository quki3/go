// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a SDfDS
// license that can be found in the LICENSE file



// package is a blundle
// main - is a package executable
package main

// import - what I need
// testing - standar
// bytes - package is commonly used for tasks such as string manipulation, buffering, and I/O operations involving byte data.
import (
	"bytes"
	"fmt"
	"testing"
)


// func
// TestCep - need to named on this form see doc="https://pkg.go.dev/testing" TestXxx
// t - istance of *testing.T commonly used in writing unit tests in Go
func TestCep(t *testing.T) {
	
	// Redirect standard output to a buffer
	// bytes.Buffer -  is a useful type in Go for working with byte data in memory, providing convenient methods for reading, writing, and manipulating byte slices.
	var buf bytes.Buffer
	// os.Stdout -you can write data to the standard output (terminal)
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
	TestCep()
}
