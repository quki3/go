# go (goland) 
`last update jul 14 2023`
## welcome
you are still here? great so les't go!

## documentation
wiki,oficial web, <a href="https://cs.opensource.google/go/go">code open source</a>.
## Instalation 
<a href="https://go.dev/dl/">go.dev</a>
## challenges
### 1. 1/ get started
   - /a  Get started with Hello, World. With the follow <a href="https://go.dev/doc/tutorial/getting-started#code">doc</a> create a programm that print "hello world" in the terminal then make the testing for this.
   - /b  Call code in an external package. With the follow <a href="https://go.dev/doc/tutorial/getting-started#code">doc</a>
   - /c Create a Go module with the follow <a href="https://go.dev/doc/tutorial/create-module">documentaion</a> create two modules.
   - /d Test, with this documentation <a href="https://pkg.go.dev/testing#hdr-Examples">testing</a> make a funcion that must retur a positive num use Abs function provided by the math package. and test the number that this func return you must make two file num.go and num_test.go
   - e/ Start a module that other can use.
   - f/ make collection of distinguishable entities such that for any entity it is determined without ambiguity whether it belongs to the collection or not ( set, objects).
   - g/ how you can determine if x belong to X or does not belong to x write a programme that do this for you.
   - h/ prints its command-line arguments on a single line use os.Args[]
   - i/ Modify the os.Arg[] program to print the index and value of each of its arguments,one per line.
   - j/ reads input lines, counts their occurrences, and prints lines that appear more than once.
   - k/ read from the standard input or handle a list of file names,using os.Open to open each one
   - l/ introduce the function ReadFile in the program k (from the io/ioutil package), which reads the entire contents of a named file, and strings.Split, which splits a string into a slice of substrings.
   - m/ demonstrates basic usage of Go’s standard image packages, which we’ll use to create a sequence of bit-mapped images and then encode the sequence as a GIF animation.
   - n/ called fetch that fetches the content of each specified URL and prints it as uninter preted text.
   - o/ use io.Copy(dst, src) reads from src and writes to dst. Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
   - p/ Fetching url concurrently in parallel and reports their times and sizes.
   - q/ find a web site that produces a large amount of data. Investigate caching byrunning fetchall twice in succession to see whether the reported time changes much. do you get the same content each time? Modify fetchall to print its output to a file so it can be examined.
   - r/ Try fetchall with longer argument lists, such as samples from the top million web sites available at alexa.com. How does the program behave if a web site just doesn’t respond?
   - s
   - t
