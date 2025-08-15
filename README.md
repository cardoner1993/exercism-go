# 🚀 Go Exercism Solutions

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Exercism Track](https://img.shields.io/badge/Exercism-Go%20Track-009639?style=flat&logo=exercism)](https://exercism.org/tracks/go)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Tests](https://img.shields.io/badge/Tests-Passing-brightgreen.svg)](#running-tests)

A comprehensive collection of [Exercism](https://exercism.org/) Go programming exercises with detailed solutions, comprehensive tests, and learning resources.

## 📁 Exercise Structure

This repository contains the following exercises:

### Exercism Go Track Exercises
- **go-binarysearch** - Binary search algorithm implementation
- **go-custom-set** - Custom set data structure
- **go-double-linked-list** - Double linked list implementation
- **go-error-handling** - Error handling patterns in Go
- **go-flatten** - Array flattening exercise
- **go-palindrome** - Palindrome detection
- **go-tally-matches** - Tournament results tallying
- **simple-linked-list** - Simple linked list implementation

### Additional Exercises
- **go_truck** - Truck management system
- **exercises_golang/exercise_3** - Custom exercise 3
- **exercises_golang/exercise_9** - Custom exercise 9

## 📋 Table of Contents

- [Exercise Structure](#-exercise-structure)
- [About](#-about)
- [Getting Started](#-getting-started)
- [Running Tests](#-running-tests)
- [Learning Resources](#-learning-resources)
- [Goals](#-goals)

## 🎯 About

This repository contains my solutions to Go exercises from the Exercism programming platform. Each solution includes:

- ✅ **Complete implementation** following Go best practices
- 🧪 **Comprehensive test suites** with edge cases
- 📚 **Detailed explanations** and learning notes
- 🔧 **Multiple solution approaches** where applicable
- 📈 **Performance benchmarks** for algorithmic problems

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/go-exercism-solutions.git
   cd go-exercism-solutions
   ```

2. **Initialize Go module (if not already done):**
   ```bash
   go mod init go-exercism-solutions
   go mod tidy
   ```

3. **Verify installation:**
   ```bash
   go test ./...
   ```

## 🧪 Running Tests

### Run All Tests
```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests with coverage
go test -cover ./...
```

### Run Specific Exercise Tests
```bash
# Run with benchmarks
go test -bench=. ./go-binarysearch

# Run with race detection
go test -race ./go-double-linked-list
```

### Generate Coverage Report
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

## 📖 Learning Resources

### Go-Specific Resources
- 📘 [Official Go Documentation](https://golang.org/doc/)
- 📗 [Effective Go](https://golang.org/doc/effective_go.html)
- 📙 [Go by Example](https://gobyexample.com/)
- 📕 [The Go Programming Language Book](https://www.gopl.io/)

### Exercism Resources
- 🎯 [Go Track on Exercism](https://exercism.org/tracks/go)
- 💡 [Go Track Concepts](https://exercism.org/tracks/go/concepts)
- 👥 [Exercism Community Solutions](https://exercism.org/tracks/go/exercises)

### Testing in Go
- 🧪 [Go Testing Package](https://golang.org/pkg/testing/)
- 📊 [Benchmarking in Go](https://golang.org/pkg/testing/#hdr-Benchmarks)
- 🔍 [Test Coverage in Go](https://golang.org/doc/tutorial/add-a-test)


## 🎯 Goals

- [ ] Complete all beginner exercises
- [ ] Achieve 100% test coverage
- [ ] Add performance benchmarks for all solutions
- [ ] Create detailed explanations for complex algorithms
- [ ] Add alternative solution approaches
- [ ] Implement CI/CD pipeline

## 📝 Notes

### Test Strategy
- Test happy path scenarios
- Test edge cases and boundary conditions
- Test error conditions
- Use table-driven tests where appropriate
- Add benchmarks for performance-critical code

## 🙏 Acknowledgments

- [Exercism](https://exercism.org/) for providing excellent coding exercises
- The Go community for continuous learning resources
- Contributors who help improve these solutions

---

⭐ **Star this repository** if you find it helpful!

🐛 **Found a bug?** [Open an issue](https://github.com/yourusername/go-exercism-solutions/issues)

💡 **Have a suggestion?** [Start a discussion](https://github.com/yourusername/go-exercism-solutions/discussions)