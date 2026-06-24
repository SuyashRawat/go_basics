# Go Garbage Collector - Comprehensive Notes

## Overview

Go includes an automatic garbage collector (GC) that frees unused memory. The GC is **concurrent**, **tri-color**, and **mark-and-sweep** based. It runs automatically without manual intervention, making memory management safer and simpler.

---

## Key Characteristics

### Concurrent
- GC runs concurrently with the main program (most of the time)
- Minimizes "stop the world" pauses (where all goroutines stop)
- Designed for low-latency applications

### Tri-Color Marking
- Objects are colored to track GC state:
  - **White**: Unvisited objects (potential garbage)
  - **Gray**: Visited but not yet fully processed
  - **Black**: Visited and fully processed (keeps references)
- Ensures all reachable objects are found

### Non-Generational
- Unlike some GCs, Go's GC doesn't distinguish between young and old objects
- Every collection scans the entire heap

---

## How Go's GC Works

### The Mark Phase
1. **Start**: GC pauses goroutines and marks root objects as gray
   - Roots: stack variables, global variables, registers
2. **Mark Worker Goroutines**: Multiple goroutines scan gray objects
   - Move objects from gray to black (mark as reachable)
   - New objects found are marked gray
3. **Write Barriers**: Track modifications during marking
   - When a black object points to white, mark the white object gray
   - Ensures no reachable objects are missed

### The Sweep Phase
1. Runs concurrently with program execution
2. Walks the heap finding white (unreachable) objects
3. Frees memory of white objects
4. Returns memory to the allocator

### Pause Time
- **Initial pause**: ~1-10ms (mark setup)
- **Final pause**: ~1-10ms (mark completion, begin sweep)
- Modern Go versions minimize these pauses through optimization

---

## GC Triggers

The GC runs automatically when:

1. **Heap Growth**: When heap size reaches a target
   - Controlled by `GOGC` environment variable (default: 100)
   - Next GC at `heap_size = previous_heap_size * (1 + GOGC/100)`

2. **Time-Based**: If GC hasn't run in 2 minutes (since Go 1.19)
   - Prevents GC starvation in programs that don't allocate much

3. **Manual**: `runtime.GC()` can force a collection
   - Useful for benchmarking or critical sections

4. **Concurrent Mark Completion**: When marking finishes

---

## Performance Considerations

### GC Overhead
- CPU usage: ~5-15% overhead on typical applications
- Memory usage: ~25% extra RAM (heap buffer for concurrent marking)
- Pause time: Typically <1ms in production (modern Go versions)

### What Affects GC Performance
- **Allocation Rate**: More allocations = more work for GC
- **Heap Size**: Larger heaps take longer to scan
- **Number of Goroutines**: More goroutines = more root scanning
- **Object Size Distribution**: Many small objects are harder to track

### Best Practices for Performance
1. **Reduce Allocations**
   - Reuse buffers and objects when possible
   - Use `sync.Pool` for temporary objects
   - Avoid unnecessary interface{} conversions

2. **Control GC Frequency**
   - Adjust `GOGC` environment variable
   - Higher GOGC = less frequent collections (more memory used)
   - Lower GOGC = more frequent collections (more CPU used)

3. **Monitor GC Behavior**
   ```bash
   GODEBUG=gctrace=1 ./your_program
   ```
   - Prints GC statistics to stderr
   - Shows pause times, heap size, collection frequency

---

## GC Tuning

### Environment Variables

#### GOGC (default: 100)
- Controls the GC trigger threshold
- Percentage increase in heap before next GC
- Example: `GOGC=50` means GC at 50% heap growth

```bash
GOGC=200 ./program    # Less frequent, more memory
GOGC=50  ./program    # More frequent, less memory
```

#### GOMAXPROCS
- Number of OS threads for GC mark workers
- Affects GC speed (more threads = faster marking)
- Usually matches number of CPU cores

#### GOMEMLIMIT (Go 1.19+)
- Sets a hard memory limit for the program
- GC tries to stay below this limit
- Useful for containerized environments

```bash
GOMEMLIMIT=512MiB ./program
```

### Runtime API

```go
// Force a garbage collection
runtime.GC()

// Get GC statistics
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Printf("Alloc: %v MB\n", m.Alloc / 1024 / 1024)
fmt.Printf("TotalAlloc: %v MB\n", m.TotalAlloc / 1024 / 1024)
fmt.Printf("Sys: %v MB\n", m.Sys / 1024 / 1024)
fmt.Printf("NumGC: %v\n", m.NumGC)

// Debug GC with per-goroutine traces
runtime.SetGCPercent(50)  // Adjust GC percentage
```

---

## Memory Profiling

### pprof Package
```go
import _ "net/http/pprof"

// In main():
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()

// Then visit:
// http://localhost:6060/debug/pprof/heap
```

### Analyzing Heap Dumps
```bash
go tool pprof http://localhost:6060/debug/pprof/heap
go tool pprof http://localhost:6060/debug/pprof/allocs
```

### Memory Stats Example
```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    
    fmt.Printf("Alloc = %v MB", m.Alloc / 1024 / 1024)
    fmt.Printf("\tTotalAlloc = %v MB", m.TotalAlloc / 1024 / 1024)
    fmt.Printf("\tSys = %v MB", m.Sys / 1024 / 1024)
    fmt.Printf("\tNumGC = %v\n", m.NumGC)
}
```

---

## Common Misconceptions

### ❌ Go leaks memory automatically
**False**: The GC is automatic and handles most memory cleanup. Manual leaks occur when:
- Holding references to large objects unnecessarily
- Circular references (Go handles these fine)
- Large maps that grow unbounded

### ❌ GC pauses are unpredictable
**Mostly False**: Modern Go has predictable pause times (<1ms typical). Pauses depend on heap size and allocation patterns, not randomness.

### ❌ You must manually free memory
**False**: This is handled automatically. Calling `delete()` on maps or `x = nil` rarely helps.

### ❌ Goroutines prevent garbage collection
**False**: GC treats goroutine stacks as roots, but blocked goroutines don't prevent collection of their referenced objects.

---

## Go Version Improvements

### Go 1.18+
- Smaller GC pauses through improved concurrency
- Better handling of very large heaps

### Go 1.19+
- `GOMEMLIMIT` environment variable
- Time-based GC trigger (2-minute guarantee)

### Go 1.20+
- Further pause reduction optimizations
- Improved allocation patterns

---

## When to Optimize GC

### Profile First
- Use `pprof` to identify actual bottlenecks
- GC might not be your problem

### Focus On
1. **Allocation Rate**: Often the biggest factor
2. **Object Lifecycle**: Keep objects alive only as long as needed
3. **Heap Size**: Smaller heaps are faster to scan

### Don't Over-Optimize
- Go's GC is highly tuned for most workloads
- Premature optimization often makes code harder to maintain
- Measure before and after changes

---

## Example: GC Behavior Visualization

```go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    var m runtime.MemStats
    
    runtime.ReadMemStats(&m)
    fmt.Printf("Before allocation: %v MB\n", m.Alloc/1024/1024)
    
    // Allocate lots of memory
    var slice []int
    for i := 0; i < 10000000; i++ {
        slice = append(slice, i)
    }
    
    runtime.ReadMemStats(&m)
    fmt.Printf("After allocation: %v MB\n", m.Alloc/1024/1024)
    
    // Force GC
    runtime.GC()
    
    runtime.ReadMemStats(&m)
    fmt.Printf("After GC: %v MB\n", m.Alloc/1024/1024)
    
    // Release slice
    slice = nil
    runtime.GC()
    
    runtime.ReadMemStats(&m)
    fmt.Printf("After releasing and GC: %v MB\n", m.Alloc/1024/1024)
}
```

---

## Summary Checklist

- ✅ GC is concurrent and automatic
- ✅ Pause times are typically <1ms in modern Go
- ✅ Tri-color marking ensures correctness
- ✅ Profile before optimizing
- ✅ Focus on reducing allocations
- ✅ Use `GODEBUG=gctrace=1` for monitoring
- ✅ Consider `GOGC` and `GOMEMLIMIT` tuning for specific workloads
- ✅ For high-performance needs, use `sync.Pool` for object reuse