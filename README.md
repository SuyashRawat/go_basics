# Go Basics

A curated collection of Golang concepts, examples, interview notes, and backend engineering patterns.

The purpose of this repository is to serve as my personal Go revision guide so that I can quickly revisit concepts before interviews without searching through multiple resources.

Each topic contains:

* Runnable examples
* Short explanations
* Common interview questions
* Important takeaways
* Production and backend engineering concepts

---

## Learning Path

### Fundamentals

* 01_basics
* 02_control_flow
* 03_functions
* 04_arrays
* 05_slices
* 06_maps
* 07_strings_and_runes

### Object-Oriented Concepts

* 08_structs
* 09_methods
* 10_pointers
* 11_interfaces

### Language Features

* 12_error_handling
* 13_packages_and_modules
* 14_generics

### Concurrency

* 15_goroutines
* 16_channels
* 17_sync_primitives
* 18_context
* 19_concurrency_patterns

### Advanced Concepts

* 20_memory_management
* 21_reflection
* 22_design_patterns
* 23_dependency_injection

### Backend Development

* 24_json
* 25_file_handling
* 26_http_server
* 27_middleware
* 28_rest_api

### Databases

* 29_database_sql
* 30_database_transactions
* 31_database_patterns

### Testing

* 32_testing
* 33_benchmarks

### Production Engineering

* 34_logging
* 35_configuration
* 36_graceful_shutdown
* 37_production_patterns

### Interview Preparation

* 38_interview_patterns
* 39_machine_coding_snippets
* 40_common_interview_questions

---

## How To Run Examples

Each concept is organized as:

```text
01_basics/
├── 01_variables/
│   └── main.go
├── 02_constants/
│   └── main.go
```

Run an example:

```bash
go run 01_basics/01_variables/main.go
```

---

## Interview Topics Covered

### Core Go

* Variables
* Constants
* Data Types
* Arrays
* Slices
* Maps
* Structs
* Methods
* Interfaces
* Pointers
* Generics

### Concurrency

* Goroutines
* Channels
* Buffered Channels
* Select
* Mutex
* RWMutex
* Atomic Operations
* Context
* Worker Pools
* Fan-In / Fan-Out
* Pipelines

### Backend

* JSON
* HTTP Servers
* Middleware
* REST APIs
* Database Access
* Transactions

### Production

* Graceful Shutdown
* Logging
* Configuration
* Error Handling
* Dependency Injection

---

## Frequently Asked Interview Questions

* Array vs Slice
* Slice Internals
* Why are slices passed by value?
* Value Receiver vs Pointer Receiver
* Interface Internals
* nil Interface vs nil Pointer
* Goroutine Leaks
* Buffered vs Unbuffered Channels
* Mutex vs RWMutex
* Context Cancellation
* Escape Analysis
* Stack vs Heap
* Panic vs Error
* JSON Marshal vs Unmarshal
* Transaction Handling
* Graceful Shutdown

---

## Goal

After completing this repository, I should be comfortable with:

* Go SDE-1 Interviews
* Backend Engineer Interviews
* Machine Coding Rounds
* Production Go Services
* Concurrency Problems
* System Design Discussions involving Go

---

## References

* Go Documentation
* Effective Go
* Go Blog
* Go Source Code
* Personal Notes & Interview Learnings
