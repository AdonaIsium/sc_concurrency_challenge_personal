# 🧠 Concepts Guide - The Strategy Manual

> *"Understanding concurrency without analogies is like learning StarCraft without ever playing RTS games, or learning Magic without understanding card game mechanics. Possible, but way harder than it needs to be."*

This guide maps every Go concurrency concept to **two frameworks your brain already understands**: Magic: The Gathering and StarCraft: Brood War. Use this as your reference when concepts feel abstract.

---

## 📖 How to Use This Guide

### When to Read This:
- **Before starting** a mission (get the mental framework)
- **During implementation** (when a concept feels unclear)
- **After struggling** (to rebuild understanding)
- **For review** (spaced repetition refreshers)

### How It's Organized:
Each concept has:
1. **What It Is** (Plain English)
2. **MTG Analogy** (How it maps to Magic)
3. **SC:BW Analogy** (How it maps to StarCraft)
4. **Why It Matters** (Real-world use cases)
5. **Common Mistakes** (What beginners do wrong)
6. **Code Pattern** (The Go implementation)

---

## 🎯 Core Primitives

### Goroutines

**What It Is:**
A lightweight thread of execution. Like a thread, but managed by Go's runtime instead of the OS. Goroutines are cheap—you can have thousands or millions.

**🎴 MTG Analogy:**
**Casting a Creature Spell**

When you cast a creature in MTG:
- It resolves and enters the battlefield
- It exists independently (has its own stats, abilities)
- It acts on your turn (controlled execution)
- Multiple creatures can be in play simultaneously
- You can sacrifice or destroy it (cleanup)

```
Go Parallel:
go myFunction() → Cast creature: "MyFunction enters the battlefield"
- Function runs independently
- Multiple goroutines exist simultaneously
- Each has its own execution context
- Can be stopped/cleaned up
```

**⚔️ SC:BW Analogy:**
**Training a Unit**

When you train a Marine in SC:BW:
- You spend resources (memory allocation)
- The unit is produced and exists independently
- It can move and act without constant micro
- You can control multiple units simultaneously
- Units can die or be destroyed (goroutine exits)

```
Go Parallel:
go myFunction() → Train Marine: "Unit ready for duty"
- Goroutine spawns and runs independently
- Multiple goroutines = multiple units in your army
- Each follows its own logic
- Needs proper cleanup when done
```

**Why It Matters:**
Goroutines let you do multiple things at once without complex threading code. Perfect for:
- Handling multiple web requests
- Processing data pipelines
- Real-time systems
- Background tasks

**⚠️ Common Mistakes:**
```go
// ❌ BAD: Goroutine leaks (units never dying)
func leaky() {
    go func() {
        for {
            // This runs forever, even if you don't need it
            doWork()
        }
    }() // No way to stop this!
}

// ✅ GOOD: Proper lifecycle control
func controlled(ctx context.Context) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return // Clean exit when context canceled
            default:
                doWork()
            }
        }
    }()
}
```

**🎯 Go Pattern:**
```go
// Basic goroutine
go func() {
    // Do work concurrently
}()

// With WaitGroup for coordination
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work
}()
wg.Wait() // Wait for completion
```

**🔥 Pro Tip from Bisu:**
> "Every unit in my army has a purpose. If I can't explain why a unit exists, I don't make it. Same with goroutines—if you can't explain why it needs to be concurrent, don't spawn it."

---

### Channels

**What It Is:**
Channels are typed conduits for communication between goroutines. They can be buffered (hold values) or unbuffered (synchronization points).

**🎴 MTG Analogy:**
**The Stack**

In MTG, the stack is how spells and abilities resolve:
- Last in, first out (LIFO) for stack
- Priority passes between players
- Things resolve in order
- You can't proceed until the stack is empty

**Channels = Controlled Stack:**
```
Unbuffered channel = Stack with instant priority pass
- Send blocks until receive (like waiting for opponent to respond)
- Receive blocks until send (like waiting for something to resolve)
- Synchronization point

Buffered channel = Stack with capacity
- Can put N things on the stack before blocking
- Eventually fills up (needs resolution)
```

**Example:**
```
Player A: Cast Lightning Bolt (send to channel)
[Stack now has Lightning Bolt]
Player B: No response (receive from channel)
[Lightning Bolt resolves]

Go equivalent:
ch <- value  // Put spell on stack
result := <-ch // Let it resolve
```

**⚔️ SC:BW Analogy:**
**Command Queue / Waypoints**

In SC:BW, units follow command queues:
- Shift-click adds commands to queue
- Units execute commands in order
- Queue has capacity (like buffer)
- Clear queue = interrupt

**Channels = Command Queue:**
```
Unbuffered channel = Direct command (do NOW)
- Marine waits for command
- Commander waits for Marine to acknowledge
- Synchronous execution

Buffered channel = Command queue
- Can queue multiple commands
- Marine executes them in order
- Commander doesn't wait for each one
```

**Example:**
```
You: "Attack here, move there, hold position" (queue 3 commands)
Marine: Executes first command...
        Executes second command...
        Executes third command...

Go equivalent:
ch := make(chan Command, 3) // Buffer of 3
ch <- Attack
ch <- Move
ch <- Hold
// Marine (goroutine) processes in order
```

**Why It Matters:**
Channels are how goroutines coordinate. Use them for:
- Producer/consumer patterns
- Work distribution
- Event notification
- Pipeline processing

**⚠️ Common Mistakes:**
```go
// ❌ BAD: Sending on closed channel (panic!)
close(ch)
ch <- value // PANIC!

// ❌ BAD: Forgetting to close channel
ch := make(chan int)
go func() {
    for v := range ch { // This loops forever if ch never closes
        process(v)
    }
}()

// ✅ GOOD: Proper channel lifecycle
ch := make(chan int)
go func() {
    defer close(ch) // Always close when done sending
    for i := 0; i < 10; i++ {
        ch <- i
    }
}()

for v := range ch { // Exits when ch closes
    process(v)
}
```

**🎯 Go Pattern:**
```go
// Unbuffered: synchronous communication
ch := make(chan int)
go func() {
    ch <- 42 // Blocks until someone receives
}()
value := <-ch // Receives (and unblocks sender)

// Buffered: asynchronous up to capacity
ch := make(chan int, 10) // Can hold 10 values
ch <- 1 // Doesn't block until buffer is full
ch <- 2
value := <-ch // value = 1

// With select: multiple channel operations
select {
case v := <-ch1:
    // Received from ch1
case ch2 <- value:
    // Sent to ch2
case <-time.After(time.Second):
    // Timeout after 1 second
}
```

**🔥 Pro Tip from Flash:**
> "In a perfect game, every command has purpose and timing. In perfect Go, every channel operation is intentional. Buffered when you need throughput, unbuffered when you need synchronization."

---

### Mutexes (sync.Mutex / sync.RWMutex)

**What It Is:**
Mutual exclusion locks. They protect shared data from concurrent access. Only one goroutine can hold a mutex at a time.

**🎴 MTG Analogy:**
**Turn Priority / Active Player**

In MTG, only one player has priority at a time:
- Active player has priority first
- You can't do things during opponent's priority
- Priority passes in a controlled way
- No simultaneous actions (usually)

**Mutex = Priority Control:**
```
Lock = Taking priority
- Only you can act
- Others must wait
- You do your thing
Unlock = Passing priority
- Now someone else can act

RWMutex = Advanced priority:
- RLock = Looking at the board (multiple players can look)
- Lock = Changing the board (only one player can modify)
```

**Example:**
```
You: *Lock* "I'm declaring attackers, don't touch the board"
[You assign attackers]
You: *Unlock* "OK, your turn to declare blockers"
Opponent: *Lock* "I'm declaring blockers"

Go equivalent:
mu.Lock()
// Modify shared state
mu.Unlock()
```

**⚔️ SC:BW Analogy:**
**Resource Patch Control**

In SC:BW, only one worker can mine a mineral patch at a time:
- Worker moves to patch (Lock)
- Worker mines (critical section)
- Worker leaves (Unlock)
- Next worker can mine

**Mutex = Resource Control:**
```
Lock = "This mineral patch is MINE"
- No other worker can use it
- Prevents "bouncing" (race condition)
Unlock = "Patch available"
- Next worker can mine

RWMutex = Gas geyser:
- Many workers can look at gas count (RLock)
- Only one can modify (Lock)
```

**Example:**
```
Worker 1: *Lock* Approach mineral patch
          Mine 8 minerals
          *Unlock* Leave patch
Worker 2: *Lock* Now I can mine
          Mine 8 minerals
          *Unlock*

Go equivalent:
mu.Lock()
balance += 8 // Only one goroutine at a time
mu.Unlock()
```

**Why It Matters:**
Mutexes prevent data races—when multiple goroutines access shared memory unsafely. Use them for:
- Protecting shared state
- Atomic operations
- Critical sections

**⚠️ Common Mistakes:**
```go
// ❌ BAD: Forgetting to unlock (deadlock!)
mu.Lock()
if err := doWork(); err != nil {
    return err // DEADLOCK: mu still locked!
}
mu.Unlock()

// ❌ BAD: Locking too much (poor performance)
mu.Lock()
doLongComputation() // Holding lock too long
mu.Unlock()

// ✅ GOOD: defer unlock (can't forget)
mu.Lock()
defer mu.Unlock()
// Do work

// ✅ GOOD: Minimize critical section
temp := doLongComputation() // Outside lock
mu.Lock()
sharedState = temp // Quick operation
mu.Unlock()

// ✅ GOOD: RWMutex for read-heavy workloads
var rwmu sync.RWMutex

// Many readers can proceed
rwmu.RLock()
value := sharedState
rwmu.RUnlock()

// Only one writer
rwmu.Lock()
sharedState = newValue
rwmu.Unlock()
```

**🎯 Go Pattern:**
```go
type SafeCounter struct {
    mu sync.Mutex
    count int
}

func (s *SafeCounter) Inc() {
    s.mu.Lock()
    defer s.mu.Unlock()
    s.count++
}

func (s *SafeCounter) Value() int {
    s.mu.Lock()
    defer s.mu.Unlock()
    return s.count
}
```

**🔥 Pro Tip from Jaedong:**
> "In Zerg, if all your workers cluster on one patch, you have terrible mining efficiency. In Go, if all your goroutines wait on one mutex, you have terrible concurrency. Design to minimize contention."

---

### Context (context.Context)

**What It Is:**
Contexts carry deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines.

**🎴 MTG Analogy:**
**Game State / Turn Timer**

In MTG tournaments, each match has:
- Time limit (deadline)
- Can concede at any time (cancellation)
- Carries game state (values)

**Context = Match Control:**
```
context.WithTimeout = "You have 50 minutes for this match"
context.WithCancel = "Player concedes, match over"
context.WithValue = "This match is best-of-3, game 2"

When context expires/cancels:
- Stop all spells on stack
- Clean up game state
- Move to next match
```

**Example:**
```
Tournament: "Round 5, you have 50 minutes"
[45 minutes pass]
Tournament: "5 minutes left!"
[Time expires]
Tournament: "Time! Finish current turn, then draw"

Go equivalent:
ctx, cancel := context.WithTimeout(parent, 50*time.Minute)
defer cancel()

select {
case <-work:
    // Work completed
case <-ctx.Done():
    // Time expired, clean up
}
```

**⚔️ SC:BW Analogy:**
**Retreat Order / Mission Briefing**

In SC:BW campaigns, you have:
- Mission objectives (values)
- Time limits (deadlines)
- Commander can order retreat (cancellation)

**Context = Command Structure:**
```
context.WithTimeout = "Complete objective in 20 minutes"
context.WithCancel = "RETREAT! All units fall back!"
context.WithValue = "Mission: Destroy enemy base"

When context cancels:
- All units stop current action
- Execute retreat order
- Return to base
```

**Example:**
```
Commander: "Squad Alpha, you have 10 minutes to clear area"
[Squad engages enemy]
Commander: "All units, RETREAT NOW!"
[Units disengage and return]

Go equivalent:
ctx, cancel := context.WithTimeout(parent, 10*time.Minute)
go squadAlpha(ctx) // Squad monitors ctx.Done()

// Later...
cancel() // Trigger retreat
```

**Why It Matters:**
Contexts provide graceful cancellation and deadline propagation. Use them for:
- API request timeouts
- Graceful shutdown
- Canceling goroutine trees
- Request-scoped values

**⚠️ Common Mistakes:**
```go
// ❌ BAD: Not checking context
func work(ctx context.Context) {
    for {
        doWork() // Never checks if we should stop!
    }
}

// ❌ BAD: Storing context in struct (anti-pattern)
type Worker struct {
    ctx context.Context // Don't do this!
}

// ✅ GOOD: Check context regularly
func work(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return // Canceled, exit cleanly
        default:
            doWork()
        }
    }
}

// ✅ GOOD: Pass context as first parameter
func worker(ctx context.Context, data string) {
    // Use ctx, don't store it
}
```

**🎯 Go Pattern:**
```go
// With timeout
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel() // Always defer cancel

// With cancellation
ctx, cancel := context.WithCancel(context.Background())
go worker(ctx)
// Later...
cancel() // Cancel all workers

// Checking cancellation
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return // Exit on cancellation
        case work := <-workChan:
            process(work)
        }
    }
}
```

**🔥 Pro Tip from Flash:**
> "A good player knows when to retreat. A bad player commits all units to a lost battle. Context is your retreat order—respect it."

---

### WaitGroups (sync.WaitGroup)

**What It Is:**
WaitGroups wait for a collection of goroutines to finish. Add when spawning, Done when finishing, Wait until all complete.

**🎴 MTG Analogy:**
**Stack Resolution**

In MTG, the stack must fully resolve before the game continues:
- Multiple spells/abilities go on stack (Add)
- Each resolves one at a time (Done)
- Can't move to next phase until stack is empty (Wait)

**WaitGroup = Stack Completion:**
```
wg.Add(1) = "Put spell on stack"
wg.Done() = "Spell resolves"
wg.Wait() = "Wait for stack to empty"

You can't proceed until all spells resolve.
```

**Example:**
```
Cast 3 spells (wg.Add(3))
Spell 1 resolves (wg.Done())
Spell 2 resolves (wg.Done())
Spell 3 resolves (wg.Done())
Stack is empty (wg.Wait() returns)
Next phase begins

Go equivalent:
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        castSpell()
    }()
}
wg.Wait() // Stack clear, continue
```

**⚔️ SC:BW Analogy:**
**Production Queue Completion**

In SC:BW, you queue unit production:
- Queue 5 Marines (Add)
- Each Marine completes (Done)
- Wait until all Marines done (Wait)

**WaitGroup = Production Tracker:**
```
wg.Add(5) = "Queue 5 Marines"
wg.Done() = "Marine #1 complete"
wg.Done() = "Marine #2 complete"
...
wg.Wait() = "All Marines ready, proceed with attack"
```

**Example:**
```
Barracks: Queuing 5 Marines (wg.Add(5))
[Marine 1 completes] (wg.Done())
[Marine 2 completes] (wg.Done())
[Marine 3 completes] (wg.Done())
[Marine 4 completes] (wg.Done())
[Marine 5 completes] (wg.Done())
Commander: All units ready! (wg.Wait() returns)

Go equivalent:
var wg sync.WaitGroup
for i := 0; i < 5; i++ {
    wg.Add(1)
    go trainMarine(&wg)
}
wg.Wait() // All Marines ready
launchAttack()
```

**Why It Matters:**
WaitGroups ensure all spawned work completes before continuing. Use them for:
- Waiting for parallel tasks
- Clean shutdown (wait for all goroutines)
- Batch processing

**⚠️ Common Mistakes:**
```go
// ❌ BAD: Add inside goroutine (race condition!)
go func() {
    wg.Add(1) // Might happen after Wait()
    defer wg.Done()
    work()
}()

// ❌ BAD: Forgetting Done (hangs forever)
wg.Add(1)
go func() {
    work()
    // Missing wg.Done()!
}()
wg.Wait() // Hangs forever

// ✅ GOOD: Add before spawning
wg.Add(1)
go func() {
    defer wg.Done() // Can't forget with defer
    work()
}()
wg.Wait()

// ✅ GOOD: Add in loop before goroutines
for i := 0; i < n; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        work()
    }()
}
wg.Wait()
```

**🎯 Go Pattern:**
```go
var wg sync.WaitGroup

// Spawn workers
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        worker(id)
    }(i)
}

// Wait for all to complete
wg.Wait()
fmt.Println("All workers finished")
```

**🔥 Pro Tip from Bisu:**
> "You don't attack until all units are in position. You don't proceed until all goroutines finish. WaitGroup is your formation check."

---

## 🎨 Design Patterns

### Worker Pool

**What It Is:**
A fixed number of goroutines that process jobs from a queue. Prevents spawning unlimited goroutines.

**🎴 MTG Analogy:**
**Mana Pool**

You have limited mana per turn:
- Fixed number of lands (workers)
- Play spells from hand (jobs from queue)
- Can't cast more spells than mana available
- Mana refreshes next turn (workers pick next job)

**Worker Pool = Mana Management:**
```
5 lands (5 worker goroutines)
10 spells in hand (10 jobs in queue)

Turn 1: Cast 5 spells (use all workers)
Turn 2: Cast next 5 spells (workers process next batch)

You're limited by mana (workers), not number of spells (jobs).
```

**⚔️ SC:BW Analogy:**
**Worker Mining**

You have N workers mining:
- Fixed number of workers (worker goroutines)
- Unlimited mineral patches (job queue)
- Each worker mines continuously
- Adding more workers increases throughput

**Worker Pool = Mining Efficiency:**
```
8 workers (8 goroutine pool)
Minerals = jobs
Each worker: Pick patch, mine, return, repeat

More workers = faster mining
But too many = inefficient (workers wait for patches)
```

**Why It Matters:**
Prevents resource exhaustion from unlimited goroutines. Use for:
- Rate limiting
- Controlled parallelism
- Server request handling

**🎯 Go Pattern:**
```go
func workerPool(jobs <-chan Job, results chan<- Result) {
    const numWorkers = 5
    var wg sync.WaitGroup

    // Spawn fixed workers
    for i := 0; i < numWorkers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs { // Process until channel closes
                result := process(job)
                results <- result
            }
        }(i)
    }

    wg.Wait()
    close(results)
}
```

---

### Fan-Out / Fan-In

**What It Is:**
- **Fan-out**: One source distributes work to multiple workers
- **Fan-in**: Multiple sources merge into one destination

**🎴 MTG Analogy:**
**Fan-Out = Proliferate**
- One trigger, affects multiple targets
- "When X happens, do Y to all creatures"

**Fan-In = Lifelink**
- Multiple creatures gain life
- All feed into one life total

**⚔️ SC:BW Analogy:**
**Fan-Out = Split Army**
- One command: "Split into squads"
- Multiple units execute independently

**Fan-In = Rally Point**
- Multiple buildings produce units
- All units converge to one rally point

**🎯 Go Pattern:**
```go
// Fan-Out
func fanOut(input <-chan int, outputs []chan int) {
    for val := range input {
        // Send to all outputs
        for _, out := range outputs {
            out <- val
        }
    }
}

// Fan-In
func fanIn(inputs []<-chan int, output chan<- int) {
    var wg sync.WaitGroup
    for _, input := range inputs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
            for val := range ch {
                output <- val
            }
        }(input)
    }
    wg.Wait()
    close(output)
}
```

---

### Pipeline

**What It Is:**
Chain of processing stages, each stage does one thing, output of one stage is input to next.

**🎴 MTG Analogy:**
**Creature Enters Battlefield → Triggers → Abilities**
```
Cast creature → Enters battlefield → Trigger ability → Resolve ability → Continue
Each stage processes and passes result to next.
```

**⚔️ SC:BW Analogy:**
**Production Chain**
```
Gather minerals → Build Barracks → Train Marines → Form Squad → Attack
Each stage completes before next begins.
```

**🎯 Go Pattern:**
```go
func pipeline() {
    // Stage 1: Generate numbers
    numbers := make(chan int)
    go func() {
        defer close(numbers)
        for i := 0; i < 10; i++ {
            numbers <- i
        }
    }()

    // Stage 2: Square numbers
    squares := make(chan int)
    go func() {
        defer close(squares)
        for n := range numbers {
            squares <- n * n
        }
    }()

    // Stage 3: Print results
    for s := range squares {
        fmt.Println(s)
    }
}
```

---

## 🚨 Anti-Patterns & Mistakes

### Data Races

**What It Is:**
Multiple goroutines accessing the same memory location, at least one writing, without synchronization.

**🎴 MTG Analogy:**
**Multiple Players Modifying Board Simultaneously**
```
Player A: "I'm setting my life to 15"
Player B: "I'm setting your life to 10" [at same time]
Result: Undefined! (One overwrites the other)
```

**⚔️ SC:BW Analogy:**
**Two Players Controlling Same Unit**
```
Player A: "Marine, attack left!"
Player B: "Marine, attack right!" [simultaneously]
Result: Marine confused, behavior undefined
```

**How to Detect:**
```bash
go run -race your_program.go
```

**How to Fix:**
- Use channels for communication
- Use mutexes for shared state
- Use atomic operations for simple counters

---

### Deadlocks

**What It Is:**
All goroutines are waiting for each other, none can proceed.

**🎴 MTG Analogy:**
**Both Players Waiting for Priority**
```
Player A: "I pass priority"
Player B: "I pass priority"
Player A: "No, you go first"
Player B: "No, YOU go first"
[Game hangs forever]
```

**⚔️ SC:BW Analogy:**
**Units Blocking Each Other**
```
Marine 1: "Can't move, Marine 2 is in my way"
Marine 2: "Can't move, Marine 1 is in my way"
[Both units stuck forever]
```

**How to Prevent:**
- Always acquire locks in the same order
- Use timeouts (context.WithTimeout)
- Don't hold locks while waiting on channels

---

## 🎓 Learning Path

### Beginner (Boot Camp)
1. Goroutines - understanding concurrency
2. Channels - basic communication
3. Mutexes - protecting shared state
4. WaitGroups - coordination

### Intermediate (Missions 1-2)
5. Contexts - cancellation and timeouts
6. Worker Pools - controlled parallelism
7. Select statements - multiple channel operations
8. Buffered channels - decoupling

### Advanced (Missions 3-4)
9. Fan-out/Fan-in patterns
10. Pipelines
11. Error handling in concurrent code
12. Performance tuning

---

## 🔥 Final Wisdom

**From Flash:**
> "Mastering macro means seeing the whole game state. Mastering concurrency means seeing the whole system state. Think in patterns, not individual goroutines."

**From Jaedong:**
> "Zerg players adapt builds mid-game. Great Go programmers adapt patterns mid-development. Don't force a pattern—choose the right tool for the problem."

**From Jon Finkel:**
> "The stack is the heart of Magic. Channels and mutexes are the heart of Go. Master these, and everything else becomes obvious."

---

**Use this guide as your strategic reference. When stuck, come back and reread the relevant concept. Understanding beats memorization.**

GG HF! 🎯
