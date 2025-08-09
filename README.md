# Go Learning Plan

This is my plan to learn Go (Golang) from scratch.  
I will tick off one topic each day.

---

## Schedule

- [ ] **Day 1:** Setup & First Go Program  
  - Install Go, set up workspace (`go mod init`, `go run`)
  - Write "Hello, World!"  
  - Print name, age, favorite programming language  

- [ ] **Day 2:** Variables, Constants, Types & I/O  
  - `var`, `const`, `:=`  
  - Basic types: `int`, `float64`, `string`, `bool`  
  - Take user input with `fmt.Scan`  
  - **Practice:** Build a simple calculator  


- [ ] **Day 3:** Control Flow (if, for, switch)  
  - `if-else`, `switch`, single `for` loop  
  - **Practice:** Prime number checker

- [ ] **Day 4:** Functions & Error Handling  
  - Declaring functions, multiple returns  
  - `error` type & `errors.New`  
  - **Practice:** `divide(a, b)` with error on `b == 0`  
<!-- 
- [ ] **Day 5:** Arrays, Slices & Maps  
  - Arrays vs slices, `append()`, `copy()`  
  - Maps (`map[string]int`)  
  - **Practice:** Average of numbers & character frequency counter  

- [ ] **Day 6:** Structs & Methods  
  - Create structs, methods with receivers  
  - Pointer vs value receivers  
  - **Practice:** `Student` struct with `PrintDetails()` method  

- [ ] **Day 7:** Interfaces & Polymorphism  
  - Declaring interfaces  
  - Implementing for different types  
  - **Practice:** `Shape` interface with `Area()` for `Circle` & `Rectangle`  

---

### **Week 2: Go Deep Dive**
- [ ] **Day 8:** Defer, Panic & Recover  
  - Using `defer` for cleanup  
  - `panic` & `recover` for safe error handling  
  - **Practice:** Division function with panic recovery  

- [ ] **Day 9:** JSON & File Handling  
  - `encoding/json` for Marshal/Unmarshal  
  - File reading/writing (`os`, `ioutil`)  
  - **Practice:** Read text file & count words, save struct as JSON  

- [ ] **Day 10:** Concurrency — Goroutines  
  - `go` keyword to run concurrent functions  
  - **Practice:** Goroutines printing numbers & letters  

- [ ] **Day 11:** Concurrency — Channels & Select  
  - Unbuffered & buffered channels  
  - `select` for multiple channel operations  
  - **Practice:** Concurrent sum of an array  

---

### **Week 3: Networking & Projects**
- [ ] **Day 12:** HTTP Servers  
  - `net/http` basics  
  - Handling routes & writing JSON responses  
  - **Practice:** Web server with `/hello` endpoint  

- [ ] **Day 13:** Consuming APIs  
  - `http.Get`, `http.Post`  
  - Parse JSON responses  
  - **Practice:** Fetch GitHub user info & print  

- [ ] **Day 14:** CLI Tool  
  - Build CLI tools with `flag` package  
  - **Practice:** To-Do list CLI app (add/remove tasks)  

- [ ] **Day 15:** Mini Project + Deployment  
  - Build a real project:  
    - **Option 1:** REST API for managing tasks (CRUD, in-memory)  
    - **Option 2:** Concurrent web scraper saving results to file  
  - Format & build binary (`go fmt`, `go build`)  
  - Push project to GitHub  

---
