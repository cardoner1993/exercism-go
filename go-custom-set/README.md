# Instructions
Create a custom __set__ type.

Sometimes it is necessary to define a custom data structure of some type, like a set. In this exercise you will define your own set. How it works internally doesn't matter, as long as it behaves like a set of unique elements.

## How to debug
When a test fails, a message is displayed describing what went wrong and for which input. You can also use the fact that console output will be shown too. You can write to the console using:

import "fmt"
fmt.Println("Debug message")
Note: When debugging with the in-browser editor, using e.g. fmt.Print will not add a newline after the output which can provoke a bug in go test --json and result in messed up test output.

# Here's what was needed:

  1. Initialize Go module: go mod init stringset
  2. Rename test file: The test file needed to be named *_test.go (changed tests.go to error_handling_test.go)
  3. Run go test

  The code is an Exercism exercise for error handling in Go. It implements a Use function that:
  - Retries resource opening on transient errors
  - Handles panics from the Frob method
  - Ensures proper cleanup with Defrob and Close calls
  - Returns appropriate errors

  All tests now pass successfully.