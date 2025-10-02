# Learning Tracker - Spaced Repetition System

> **Purpose**: Track concepts and skills you're learning, schedule reviews for optimal retention.
>
> **Based on**: Spaced repetition research - review at increasing intervals for long-term memory.
>
> **Target Length**: ~100-200 lines (archive mastered concepts to keep it manageable)

---

## 📊 Review Schedule Overview

**Spaced Repetition Intervals**:
- **Day 1**: Learn concept (initial exposure)
- **Day 2**: First review (24 hours later)
- **Day 4**: Second review (2 days later)
- **Day 8**: Third review (4 days later)
- **Day 15**: Fourth review (7 days later)
- **Day 30**: Fifth review (15 days later)
- **Day 60**: Mastery check (30 days later)

**Intervals adjust based on recall difficulty:**
- ✅ **Easy recall**: Use longer interval
- 😐 **Medium recall**: Stick to schedule
- ❌ **Hard recall**: Shorten interval, review sooner

---

## 🎯 Active Learning (In Progress)

> Concepts currently being learned and reinforced

### Concept: Goroutines

**Category**: Core Concurrency Primitive

**Status**: 🌱 Learning (Boot Camp)

**Learned On**: [To be filled when you start]

**💰 Point Value**: See table below for goroutine-related implementations

**What It Is**:
Lightweight threads of execution managed by the Go runtime. Goroutines are cheap to create (you can spawn thousands), making concurrent programming practical and elegant.

**Why It Matters**:
Goroutines are the foundation of all concurrency in Go. You can't build concurrent systems without understanding how they work, when to use them, and how to manage their lifecycle.

**How We Use It**:
```go
// Basic goroutine spawning
go myFunction()

// With WaitGroup for coordination
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do work concurrently
}()
wg.Wait()

// In this project: each Unit runs as a goroutine
go unit.run(ctx) // Unit lives independently
```

**Key Insight**:
Goroutines are like casting creature spells in MTG or training units in SC:BW—once spawned, they exist independently. But unlike game units, goroutines need explicit lifecycle management (no automatic cleanup!).

**Connected To**:
- OS threads (but much lighter weight)
- JavaScript async functions (but with true parallelism)
- MTG: Casting creatures (they resolve and exist on battlefield)
- SC:BW: Training units (they're produced and act independently)

**Review Schedule**:
- ⏳ Day 1: Learn concept (during Boot Camp)
- ⏳ Day 2: First review
- ⏳ Day 4: Second review
- ⏳ Day 8: Third review
- ⏳ Day 15: Fourth review
- ⏳ Day 30: Fifth review
- ⏳ Day 60: Mastery check

**Practice Exercises**:
- [ ] Spawn 1000 goroutines and observe memory usage with pprof
- [ ] Create a goroutine leak intentionally, then detect and fix it
- [ ] Implement proper cleanup using context cancellation
- [ ] Build a worker pool with fixed number of goroutines
- [ ] Race two goroutines and understand the non-determinism

**Real-World Applications** (Count: 0):
- (Will track as you implement in types.go, units.go, etc.)

**Point Opportunities**:
| Implementation | Base Points | Category |
|----------------|-------------|----------|
| Unit.run() goroutine loop | 60 pts | Advanced (lifecycle) |
| Worker pool implementation | 50 pts | Advanced (coordination) |
| Concurrent event processing | 30 pts | Advanced |
| Context-based shutdown | 35 pts | Advanced (cleanup) |

**Resources**:
- CONCEPTS_GUIDE.md (Goroutines section)
- Effective Go: Goroutines
- Go blog: Share Memory By Communicating

---

### Concept: Channels

**Category**: Core Concurrency Primitive

**Status**: 🌱 Learning (Boot Camp)

**Learned On**: [To be filled when you start]

**💰 Point Value**: See table below for channel-related implementations

**What It Is**:
Typed conduits for communication between goroutines. Channels can be unbuffered (synchronization point) or buffered (asynchronous queue up to capacity).

**Why It Matters**:
Go's philosophy: "Don't communicate by sharing memory; share memory by communicating." Channels are THE way goroutines coordinate and exchange data safely.

**How We Use It**:
```go
// Unbuffered: synchronous handoff
ch := make(chan int)
go func() {
    ch <- 42 // Blocks until someone receives
}()
value := <-ch // Receives (unblocks sender)

// Buffered: asynchronous up to capacity
eventCh := make(chan Event, 100)
eventCh <- event // Doesn't block until buffer full

// With select: multi-channel operations
select {
case msg := <-ch1:
    // Handle ch1
case ch2 <- value:
    // Send to ch2
case <-time.After(time.Second):
    // Timeout
}
```

**Key Insight**:
Channels are like the MTG stack—things resolve in order, timing matters. Unbuffered channels = instant priority pass (synchronous). Buffered = stack with capacity (asynchronous).

**Connected To**:
- Queues (but with built-in synchronization)
- MTG stack (ordered resolution, LIFO for stack, FIFO for channels)
- SC:BW command queue (units execute orders in sequence)
- Unix pipes (data flows through pipeline)

**Review Schedule**:
- ⏳ Day 1: Learn concept
- ⏳ Day 2: First review
- ⏳ Day 4: Second review
- ⏳ Day 8: Third review
- ⏳ Day 15: Fourth review
- ⏳ Day 30: Fifth review
- ⏳ Day 60: Mastery check

**Practice Exercises**:
- [ ] Build a 3-stage pipeline using channels
- [ ] Experiment with different buffer sizes and observe blocking behavior
- [ ] Create a deadlock with channels, understand why, then fix it
- [ ] Implement fan-out/fan-in pattern
- [ ] Use select to handle multiple channels with timeout

**Real-World Applications** (Count: 0):
- (Will track uses in battle simulator, command system, etc.)

**Point Opportunities**:
| Implementation | Base Points | Category |
|----------------|-------------|----------|
| Command handling with channels | 15 pts | Intermediate |
| Channel pipeline (3+ stages) | 40 pts | Advanced (composition) |
| Select with timeout/cancel | 30 pts | Advanced |
| Event pub/sub system | 55 pts | Advanced (architecture) |
| Priority queue with channels | 45 pts | Advanced (algorithm) |

**Resources**:
- CONCEPTS_GUIDE.md (Channels section)
- Go blog: Go Concurrency Patterns
- Effective Go: Channels

---

### Concept: Mutexes (sync.Mutex / sync.RWMutex)

**Category**: Core Synchronization Primitive

**Status**: 🌱 Learning (Boot Camp)

**Learned On**: [To be filled when you start]

**💰 Point Value**: See table below for mutex-related implementations

**What It Is**:
Mutual exclusion locks that protect shared data from concurrent access. Only one goroutine can hold a mutex at a time (Lock). RWMutex allows multiple readers OR one writer.

**Why It Matters**:
Prevents data races—the most common and dangerous bug in concurrent programs. Mutexes are essential when you need to protect shared state.

**How We Use It**:
```go
type SafeCounter struct {
    mu    sync.RWMutex
    count map[string]int
}

func (s *SafeCounter) Inc(key string) {
    s.mu.Lock()         // Exclusive access
    defer s.mu.Unlock() // Always unlock
    s.count[key]++
}

func (s *SafeCounter) Value(key string) int {
    s.mu.RLock()        // Shared read access
    defer s.mu.RUnlock()
    return s.count[key]
}
```

**Key Insight**:
Mutexes protect DATA, not code. Only hold the lock while accessing shared state. It's like MTG priority—only one player can modify the board at a time, but multiple can look.

**Connected To**:
- Semaphores (but simpler, binary lock)
- MTG priority (only active player modifies game state)
- SC:BW mineral patch (only one worker mines at once)
- Database locks (row-level locking)

**Review Schedule**:
- ⏳ Day 1: Learn concept
- ⏳ Day 2: First review
- ⏳ Day 4: Second review
- ⏳ Day 8: Third review
- ⏳ Day 15: Fourth review
- ⏳ Day 30: Fifth review
- ⏳ Day 60: Mastery check

**Practice Exercises**:
- [ ] Create a data race intentionally, detect with `-race`, fix with mutex
- [ ] Compare sync.Mutex vs sync.RWMutex performance with benchmarks
- [ ] Create a deadlock through improper lock ordering, then fix it
- [ ] Minimize critical section to improve performance
- [ ] Build a thread-safe cache with RWMutex

**Real-World Applications** (Count: 0):
- (Will track uses in resource manager, unit state, etc.)

**Point Opportunities**:
| Implementation | Base Points | Category |
|----------------|-------------|----------|
| Unit state getters (RWMutex) | 15 pts | Intermediate (concurrency) |
| Unit state setters (RWMutex) | 15 pts | Intermediate (concurrency) |
| Thread-safe cache | 30 pts | Advanced |
| Resource pool with locking | 165 pts | Complex System (Mission 2) |

**Resources**:
- CONCEPTS_GUIDE.md (Mutexes section)
- Go blog: The Go Memory Model
- Effective Go: Concurrency

---

---

## 📅 Review Due Today

> Concepts that need review this session (auto-populated based on schedule)

### Due for Review: [Concept Name]
**Last Reviewed**: [Date] - [Days ago]
**Review Type**: [Third review / Fourth review / etc.]
**Review Method**: [How to review - explain back, apply to new problem, teach to Claude, etc.]

---

## 🎓 Skills Development

> Broader skills being developed through deliberate practice

### Skill: [Name of Skill]

**Category**: [e.g., Debugging, Code Review, Testing, Architecture Design, etc.]

**Current Level**: Beginner / Intermediate / Advanced / Expert

**Goal Level**: [Target level you want to reach]

**Why This Matters**:
[Your motivation for developing this skill]

**Practice Plan**:
1. [Specific way to practice - e.g., "Debug one issue per session without AI help first"]
2. [Another practice approach]
3. [Progressive challenge]

**Progress Indicators**:
- [ ] [Milestone 1 - e.g., "Can debug simple issues independently"]
- [ ] [Milestone 2 - e.g., "Can debug complex issues with minimal help"]
- [ ] [Milestone 3 - e.g., "Can teach debugging strategies to others"]

**Evidence of Growth**:
- [Date]: [Example of skill application]
- [Date]: [Another example showing progress]

**Next Practice Opportunity**:
[Specific upcoming task where you'll practice this]

---

### Skill: [Another Skill]

[Same structure]

---

## 💡 "Aha Moments" Log

> Capture insights and breakthroughs for reinforcement

### [YYYY-MM-DD]: [Concept/Topic]

**The Realization**:
[What clicked? What did you suddenly understand?]

**What Led to It**:
[What triggered this insight? Question, example, debugging, etc.]

**Why It Matters**:
[How this changes your understanding or approach]

**How to Remember**:
[Mental model, analogy, or memory aid]

---

## 🔄 Review Session Notes

> Track how reviews go to adjust intervals

### Review Session: [YYYY-MM-DD]

**Concepts Reviewed**:

1. **[Concept Name]**
   - Recall Difficulty: Easy / Medium / Hard
   - Could Explain: Yes / Partially / No
   - Could Apply: Yes / Needed Help / No
   - Next Review: [Adjusted date based on performance]
   - Notes: [Any observations]

2. **[Another Concept]**
   - [Same structure]

**Review Methods Used**:
- [Method 1 - e.g., "Explained concept to Claude"]
- [Method 2 - e.g., "Applied to new code problem"]

**Insights from Review**:
[What did reviewing reveal about your understanding?]

---

## 📈 Mastery Progress

### Concepts Approaching Mastery

> 4+ successful reviews, applied 5+ times in real code

**[Concept Name]**:
- Reviews: 5/6 completed
- Applications: 7 real-world uses
- Next: Mastery check on [date]
- Confidence: 8/10

---

### Recently Mastered (Last 30 Days)

**[Concept Name]** - Mastered on [YYYY-MM-DD]
- Total reviews: 6
- Real applications: 8
- Key insight: [Main takeaway]
- Archived to: `/archive/mastered-concepts/`

---

## 🎯 Learning Goals & Milestones

### Active Learning Goals

#### Goal: [Specific Learning Goal]

**Target**: [What you want to achieve]

**Timeline**: [When you want to achieve it by]

**Why**: [Your motivation]

**Success Criteria**:
- [ ] [Measurable criterion 1]
- [ ] [Measurable criterion 2]
- [ ] [Measurable criterion 3]

**Concepts Needed**:
- [Concept 1] - Status: ✅ Mastered / 🌿 Learning / ⏳ Not Started
- [Concept 2] - Status: [...]

**Progress**: [Brief status update]

**Next Milestone**: [What's the next step toward this goal]

---

### Completed Learning Goals

**[Goal Name]** - Achieved [YYYY-MM-DD]
- Duration: [How long it took]
- Key learnings: [Main takeaways]
- Evidence: [How you demonstrated achievement]

---

## 🧪 Practice Exercises

> Deliberate practice opportunities to build specific skills

### Exercise: [Exercise Name]

**Skill Practiced**: [What skill this develops]

**Difficulty**: Beginner / Intermediate / Advanced

**Description**:
[What to do]

**Learning Focus**:
[What this exercise teaches]

**Completed**: [ ] Not Started / [✅] Done on [date]

**Outcome**:
[What you learned or struggled with]

**Variations to Try**:
- [Harder version]
- [Different approach]

---

## 📝 Concept Quick Reference

> Fast lookup for concepts you're learning (summary view)

| Concept | Category | Status | Next Review | Applications |
|---------|----------|--------|-------------|--------------|
| [Name] | [Category] | 🌿 | YYYY-MM-DD | 3 |
| [Name] | [Category] | 🌳 | YYYY-MM-DD | 7 |
| [Name] | [Category] | ✅ | Complete | 12 |

---

## 🔍 Gap Analysis

> Identify knowledge gaps and areas for growth

### Known Gaps:
- **[Topic/Concept]**: [Why this is a gap] - Priority: High/Medium/Low
  - Plan to address: [How and when]

### Areas of Confusion:
- **[Concept]**: [What's unclear]
  - Questions to answer: [...]
  - Resources to explore: [...]

### Wishlist (Future Learning):
- [Concept to learn later]
- [Skill to develop in future]

---

## 📚 Learning Resources

> Helpful resources organized by topic

### [Topic/Concept Category]
- **Documentation**: [Link]
- **Tutorial**: [Link]
- **Video**: [Link]
- **Article**: [Link]
- **Book/Chapter**: [Reference]

### [Another Topic]
[Same structure]

---

## 🎓 Teaching Moments

> Concepts you've successfully taught back to Claude (best learning indicator!)

### [YYYY-MM-DD]: Taught [Concept]

**How You Explained It**:
[Your explanation - shows your understanding]

**Gaps Claude Identified**:
[What you missed or got wrong]

**Improved Understanding**:
[How your understanding evolved through teaching]

---

## Template Usage Guide

### Daily Workflow

**Session Start (2 minutes)**:
1. Check "Review Due Today" section
2. Claude asks you to recall due concepts
3. Update recall difficulty based on performance
4. Adjust next review date if needed

**During Session (ongoing)**:
5. When learning new concept, add to "Active Learning"
6. Set up review schedule
7. When you have an "aha moment", log it
8. When you apply a concept, note it in "Real-World Applications"

**Session End (3 minutes)**:
9. Update review schedules based on today's reviews
10. Add any new concepts learned
11. Update skill progress if practiced
12. Move mastered concepts to archive if ready

### Adding a New Concept

**When you learn something new:**

```markdown
### Concept: [Name]
**Category**: [Type]
**Status**: 🌱 Learning
**Learned On**: 2024-03-15
**What It Is**: [Simple explanation]
**Why It Matters**: [Importance]
**How We Use It**: [Code example]
**Key Insight**: [Aha moment]
**Connected To**: [Related concepts]

**Review Schedule**:
- ✅ Day 1 (2024-03-15): Learned - Session #15
- 📅 Day 2 (2024-03-16): First review - **DUE NEXT SESSION**
- ⏳ Day 4 (2024-03-18): Second review
- ⏳ Day 8 (2024-03-22): Third review
- ⏳ Day 15 (2024-03-29): Fourth review
- ⏳ Day 30 (2024-04-13): Fifth review
- ⏳ Day 60 (2024-05-13): Mastery check

**Practice Exercises**:
- [ ] [Specific exercise]

**Real-World Applications** (Count: 0):
[Will add as you use it]
```

### Conducting a Review

**Review Process:**

1. **Claude asks recall question** (without looking at notes):
   - "Can you explain how [concept] works?"
   - "What's the key insight about [concept]?"
   - "How would you apply [concept] to solve [new problem]?"

2. **You attempt to recall** from memory

3. **Rate your recall**:
   - ✅ **Easy**: Remembered everything, explained clearly → Use longer interval
   - 😐 **Medium**: Remembered most, needed minor prompts → Stick to schedule
   - ❌ **Hard**: Struggled to recall, needed significant help → Shorten interval

4. **Update the tracker**:
   ```markdown
   - ✅ Day 8 (2024-03-22): Third review - Recall: Medium
   - 📅 Day 15 (2024-03-29): Fourth review - **DUE NEXT SESSION**
   ```

5. **Apply the concept**: If possible, use it in current work

### Review Interval Adjustments

**Based on recall difficulty:**

| Recall | Current Interval | Next Interval |
|--------|------------------|---------------|
| Easy | 2 days | 4 days |
| Medium | 2 days | 3 days |
| Hard | 2 days | 1 day |
| Easy | 4 days | 7 days |
| Medium | 4 days | 5 days |
| Hard | 4 days | 2 days |

**Pattern**: Easy = 2x interval, Medium = 1.5x interval, Hard = 0.5x interval

### Marking as Mastered

**Mastery Criteria** (all must be met):
- ✅ Completed 6+ reviews successfully
- ✅ Applied in real code 5+ times
- ✅ Can explain clearly without notes
- ✅ Can apply to new problems independently
- ✅ Can teach it to someone else

**When mastered:**
1. Update status to ✅ Mastered
2. Add to "Recently Mastered" section
3. Archive to `/archive/mastered-concepts/YYYY-MM-DD-[concept-name].md`
4. Keep a reference line in CLAUDE.md if frequently used

---

## Automation Helpers

### Quick Add Script (Optional)

Create a helper script to quickly add concepts:

```bash
#!/bin/bash
# add-concept.sh
# Usage: ./add-concept.sh "Concept Name" "Category"

CONCEPT_NAME=$1
CATEGORY=$2
TODAY=$(date +%Y-%m-%d)
DAY2=$(date -v+1d +%Y-%m-%d)
DAY4=$(date -v+3d +%Y-%m-%d)
DAY8=$(date -v+7d +%Y-%m-%d)
DAY15=$(date -v+14d +%Y-%m-%d)
DAY30=$(date -v+29d +%Y-%m-%d)
DAY60=$(date -v+59d +%Y-%m-%d)

cat >> LEARNING_TRACKER.md << EOF

### Concept: $CONCEPT_NAME

**Category**: $CATEGORY
**Status**: 🌱 Learning
**Learned On**: $TODAY
**What It Is**: [TODO: Add explanation]
**Why It Matters**: [TODO: Add importance]
**How We Use It**: [TODO: Add code example]
**Key Insight**: [TODO: Add aha moment]
**Connected To**: [TODO: Add related concepts]

**Review Schedule**:
- ✅ Day 1 ($TODAY): Learned
- 📅 Day 2 ($DAY2): First review - **DUE NEXT SESSION**
- ⏳ Day 4 ($DAY4): Second review
- ⏳ Day 8 ($DAY8): Third review
- ⏳ Day 15 ($DAY15): Fourth review
- ⏳ Day 30 ($DAY30): Fifth review
- ⏳ Day 60 ($DAY60): Mastery check

**Practice Exercises**:
- [ ] [TODO: Add exercise]

**Real-World Applications** (Count: 0):
[Add as you use it]

EOF

echo "Added $CONCEPT_NAME to LEARNING_TRACKER.md"
```

### Claude Automation

**Add to Claude's session start routine:**
```
1. Read LEARNING_TRACKER.md
2. Check today's date
3. Identify concepts due for review (review date = today or earlier)
4. At session start, ask: "Before we begin, let's review [concept]. Can you explain how it works?"
5. Based on user's response, update recall difficulty
6. Calculate next review date
```

---

## Integration with Other Files

### Relationship to CLAUDE.md

**LEARNING_TRACKER.md**: Detailed review schedule, practice tracking
**CLAUDE.md - Critical Concepts**: Summarized version of important concepts

**Flow**:
1. New concept learned → Add to LEARNING_TRACKER.md with full details
2. After 2-3 successful reviews → Add summary to CLAUDE.md Critical Concepts
3. After mastery → Archive from LEARNING_TRACKER, keep summary in CLAUDE.md
4. If frequently used → Keep in both places for easy reference

### Relationship to SESSION_TEMPLATE.md

**At session end:**
1. New concepts from session → Add to LEARNING_TRACKER.md
2. Reviews conducted → Update in LEARNING_TRACKER.md
3. Session log references → Note concepts learned with "Added to LEARNING_TRACKER: Yes"

---

## Success Metrics

**This tracker is working when:**
- ✅ You remember to review concepts at scheduled times
- ✅ Recall improves with each review
- ✅ Concepts successfully applied in real code multiple times
- ✅ You feel confident explaining concepts without notes
- ✅ Learning feels systematic, not random

**Needs adjustment if:**
- ❌ Reviews feel like a chore (make them more interactive)
- ❌ Recall isn't improving (shorten intervals)
- ❌ Too many concepts tracked (focus on fewer at a time)
- ❌ Reviews aren't happening (improve session start routine)
- ❌ File exceeds 250 lines (archive mastered concepts)

---

## Anti-Patterns to Avoid

❌ **Don't:**
- Add every tiny thing you learn (focus on key concepts)
- Skip reviews (defeats the entire purpose)
- Mark as mastered too early (need proof through application)
- Let file grow indefinitely (archive mastered concepts)
- Review passively (actively recall, don't just re-read)
- Track too many concepts at once (5-10 active is ideal)

✅ **Do:**
- Add concepts that are important to master
- Stick to the review schedule consistently
- Require multiple real-world applications before mastery
- Archive to keep file manageable
- Use active recall for reviews
- Focus depth over breadth (master few things well)

---

**Remember**: Spaced repetition works because of the spacing. The intervals are based on research into how memory works. Trust the system, stick to the schedule, and watch your retention soar! 🚀
