# Learning Go

I started with Tour of Go, then went on to read _The Go Programming Language_ by Alan A. A. Donovan, Brian W. Kernighan, from the perspective of a Go beginner.


# The Go Programming Language
## 1 - Tutorial
_This chapter is a tour of the basic components of Go._

First we write a hello-world program and then some echo programs in different ways.

TODO come back and do the benchmark after reading Section 11.


## bufio package 

https://golang.org/pkg/bufio/

Aims to make textual input and output efficient and convenient
*Scanner* - _Scanner provides a convenient interface for reading data such as a file of newline-delimited lines of text._


# 1 - Composite types
For efficiency, larger struct types are usually passed to or returned from functions indirectly using a pointer, and this is required if the function must modify its argument, since in a call-by-value language like Go, the called function receives only a copy of an argument, not a reference to the original argument.


```go
pp := &Point{1, 2}
```


