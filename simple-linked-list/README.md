# Go/Flatten

## Instructions
You work for a music streaming company.

You've been tasked with creating a playlist feature for your music player application.

Instructions
Write a prototype of the music player application.

For the prototype, each song will simply be represented by a number. Given a range of numbers (the song IDs), create a singly linked list.

Given a singly linked list, you should be able to reverse the list to play the songs in the opposite order.


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