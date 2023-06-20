# GO (goland) 
`last update jun 20 2023`

## INTRODUCTION
Go or golang is one of the youngest programming languages that is become more and more popular in the cloud engineering world
Was created by google in 2007 and open sourced in 2009

### Why Go was developed?
The evolution of Infrastructure was changed a lot multi-core processors became common and also using cloud infrastructure with hubdreds or thousands of servers with multiple processors. the Existing Programming Languages did not fully take advantage of itso basically you had application that would just execute one task at a time in order so wheter we can use Multi-Threading you can have many of process running in parallel and this makes application fast but also cause some issues.
example: in google docs many users can work on the same document at the same time so when you have two users changing and  adding stuff at the same time to the same document this should work smoothly without one user overriding all the changes that another user is making.
example:issues may happen when things are processedin parallel is a booking system for buying tickets or booking a hotel etc let's say again two users are trying to book the last remaining ticket at the same time this is a concept is called concurrency and needs to e handled by developers in code so they must write code that prevents any conflicts between the tasks when multiple tasks are running in parallel and updating the same data. the code can get pretty complex and handling and preventing the concurrency issues can be pretty hard.

GO was designed to run on multiple cores and built to support concurrency
Concurency in Go is cheap and easy


### Go use cases
what it's best used for is writing appication that need to be very performant and will run on the modern scaled and distributed infrastructure (cloud platform)

### Characteristics of go
simple and readable syntax of a dynamically typed language like python
speed and efficiency of a lower level language like c ++ 
copiled language

### How it compares to other programing languages?

## INSTALLATION & LOCAL DEV ENVIRONMENT
### Local Setup - to write GO code
1. Download Go
	- <a href="https://go.dev/dl/">Download in go.dev</a>
2. Install Go
	- Go distribution actually comes with a Go CLI tool
	- 1.1 Remove any previous Go installation by deleting the <b>/usr/local/go</b> folder
	- `rm -rf /usr/local/go`

@here@you@https://www.youtube.com/watch?v=yyUHQIec83I

3. Install an IDE -Editor for writing Code

## BASIC STRUCTURE OF A GO FILE




# write a simple CLI application (to learn the core concepts and syntax of go)




## resources
Oficial page
 https://go.dev/learn/
book 
Go in action autor "William kenedy"
https://play.google.com/books/reader?id=nDszEAAAQBAJ&hl=es_419
book free
<a href="./bookFree/gobook.png">gobook</a>

