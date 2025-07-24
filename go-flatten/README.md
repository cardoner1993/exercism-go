# Go/Flatten

## Instructions
Take a nested list and return a single flattened list with all values except nil/null.

The challenge is to write a function that accepts an arbitrarily-deep nested list-like structure and returns a flattened structure without any nil/null values.

For example:

input: [1,[2,3,null,4],[null],5]

output: [1,2,3,4,5]


Here's what was needed:

  1. Initialize Go module: go mod init erratum
  2. Rename test file: The test file needed to be named *_test.go (changed tests.go to error_handling_test.go)
  3. Run go test

  The code is an Exercism exercise for error handling in Go. It implements a Use function that:
  - Retries resource opening on transient errors
  - Handles panics from the Frob method
  - Ensures proper cleanup with Defrob and Close calls
  - Returns appropriate errors

  All tests now pass successfully.