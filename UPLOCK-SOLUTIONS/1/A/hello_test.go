// package - refers to a container or bundle that organizes related code, files, and resources together. It is a way of structuring and organizing code within a software project
// main - when you develop executable programs, you will use the package “main” for making the package as an executable program. The package “main” tells the Go compiler that the package should compile as an executable program instead of a shared library. The main() function in the package “main” will be the entry point of our executable program. When you build shared libraries, you will not have any main package and main function in the package.
package main 

// import - It's The way you can leverage funtionalites and APIs from external package m your own code
import ( 
  // os package , which is part of standard Library of go
  "os"
  // testing package , which is part of standard Library of go
  "testing"
  )

// func - The keyword used to define a function.
// init - It's a Func Interface of The package testing Init registers testing flags. These flags are automatically registered by the "go test" command before running test functions, so Init is only needed when calling functions such as Benchmark without using "go test".
func init ( ) {
  	os.Setenv("LC_ALL", "en")
}

// TestMain - TestXxx(*testing.T)where Xxx does not start with a lowercase letter.
// func - is the Keyword used to define a function
// t *testing.T - t *testing.T is a parameter type used in testing functions to receive a testing-related object of type *testing.T.you typically include t *testing.T as a parameter to access the testing-related functionalities.
// if - It's a conditional statement used to executed a block of code
// := - we can defined a variable is The sort form of var 
// != - operador used to check inequality between TWo values
// %q - format Verb odds "" double quotes around the string value
func TestMain(t *testing.T) { 
	hello := "Hello, world."
	if out := main(); out != hello {
		t.Errorf("main() = %q, want %q", out, hello)
	}
}
