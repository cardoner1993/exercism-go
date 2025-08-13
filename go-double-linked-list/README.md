# Go/Flatten

## Instructions
Your team has decided to use a doubly linked list to represent each train route in the schedule. Each station along the train's route will be represented by a node in the linked list.

You don't need to worry about arrival and departure times at the stations. Each station will simply be represented by a number.

Routes can be extended, adding stations to the beginning or end of a route. They can also be shortened by removing stations from the beginning or the end of a route.

Sometimes a station gets closed down, and in that case the station needs to be removed from the route, even if it is not at the beginning or end of the route.

The size of a route is measured not by how far the train travels, but by how many stations it stops at.

**Note**
The linked list is a fundamental data structure in computer science, often used in the implementation of other data structures. As the name suggests, it is a list of nodes that are linked together. It is a list of "nodes", where each node links to its neighbor or neighbors. In a singly linked list each node links only to the node that follows it. In a doubly linked list each node links to both the node that comes before, as well as the node that comes after.

If you want to dig deeper into linked lists, check out this article that explains it using nice drawings.

You will write an implementation of a doubly linked list. Implement a Node to hold a value and pointers to the next and previous nodes. Then implement a List which holds references to the first and last node and offers functions for adding and removing items.

Your Node should have the following fields and methods:

Value: the node's value (we will use interface{}).
Next() *Node: pointer to the next node.
Prev() *Node: pointer to the previous node.
You should have a function NewList() that creates and returns a List:

NewList(args ...interface{}) *List: creates a new linked list preserving the order of the values.
Your List should have the following methods:

First() *Node: returns a pointer to the first node (head).
Last() *Node: returns a pointer to the last node (tail).
Push(v interface{}): insert value at the back of the list.
Pop() (interface{}, error): remove value from the back of the list.
Unshift(v interface{}) : insert value at the front of the list.
Shift() (interface{}, error): remove value from the front of the list.
Reverse(): reverse the linked list.


Here's what was needed:

  1. Initialize Go module: go mod init flatten
  2. Rename test file: The test file needed to be named *_test.go (changed tests.go to error_handling_test.go)
  3. Run go test

  The code is an Exercism exercise for error handling in Go. It implements a Use function that:
  - Retries resource opening on transient errors
  - Handles panics from the Frob method
  - Ensures proper cleanup with Defrob and Close calls
  - Returns appropriate errors

  All tests now pass successfully.