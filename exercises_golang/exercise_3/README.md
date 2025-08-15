# 3. Exercise 003 : Create a map with numbers squared

With a given integral number n, write a program to generate a map that contains (i, i*i) such that is an integral number between 1 and n (both included), and then the program should print the map with representation of the value

Suppose the following input is supplied to the program: 8

Then, the output should be: map[1:1 2:4 3:9 4:16 5:25 6:36 7:49 8:64]

## Steps
Start by: go mod init exercise3

## Tests
Included tests are in exercise_3_test.go run it with: go test
Run specific test with: go test -run <test-name>

### NOTE
Basic Test StructureGo has built-in testing support. Test files must end with _test.go and contain functions that start with Test