# Example repo of how to test exec.Command (and other system functions) in Golang

Accompanies this article: [Unit-Testing-Exec-Command-In-Golang](https://jamiethompson.me/posts/Unit-Testing-Exec-Command-In-Golang/)

## Overview
Package `funshell` contains the `ShellCommand` function which is being tested.  
`ShellCommand` just kicks off the shell command `echo 'fun!'`.  
The `main` function in `main.go` just calls this, catches errors and then prints out anything that was output to stdout.  
Package `funshell_test` runs some simple unit tests against `ShellCommand`.  

## Running
Run the program with `go run main.go`.  
Run the tests with `go test ./..`.