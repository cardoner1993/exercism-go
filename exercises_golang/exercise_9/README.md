# Exercise 009
Write a program that accepts a sequence of whitespace separated words as input and prints the words after removing all duplicate words and sorting them alphanumerically.

Suppose the following input is supplied to the program:
```
hello world and practice makes perfect and hello world again
```
Then, the output should be:
```
again and hello makes perfect practice world
```
## Steps
Start by: go mod init exercise3
Then add: go get github.com/deckarep/golang-set/v2

## Tests
Included tests are in exercise_3_test.go run it with: go test
Run specific test with: go test -run <test-name>

### NOTE
Basic Test StructureGo has built-in testing support. Test files must end with _test.go and contain functions that start with Test