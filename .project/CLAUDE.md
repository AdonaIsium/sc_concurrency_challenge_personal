# Project Memory - CLAUDE.md

> **Purpose**: This file serves as Claude's primary memory for your project. Keep it concise (~400-600 lines) for efficient session starts.
>
> **Update Frequency**: Update at end of each session or when making significant architectural decisions.

---

## 1. Project Overview

### Project Name
**StarCraft Concurrency War - The Masterclass**

### Description
A discovery-based learning project that teaches Go concurrency through guided exploration using StarCraft: Brood War and Magic: The Gathering as teaching frameworks. The goal is deep understanding through Socratic questions, progressive challenges, and game-based analogies—not just copy-paste coding. This is a **learning tool**, not a production application.

### Tech Stack
- **Language**: Go 1.21+
- **Framework**: Standard library (no external frameworks)
- **Key Libraries**:
  - `sync` - Core concurrency primitives (Mutex, WaitGroup, etc.)
  - `context` - Cancellation and timeout management
  - Standard library packages only (teaching fundamentals)
- **Testing**: Built-in `testing` package + race detector (`-race` flag)
- **Infrastructure**: Local development only (no deployment needed)

### Project Structure
```
starcraft_concurrency_war_claude/
├── .project/                    # Learning system memory files
│   ├── CLAUDE.md               # This file - long-term project memory
│   ├── ACTIVE_CONTEXT.md       # Working memory - current session state
│   ├── LEARNING_TRACKER.md     # Spaced repetition & concept tracking
│   └── SYSTEM_PROMPT.md        # Teaching methodology configuration
│
├── internal/                    # Core learning modules
│   ├── types/                  # Boot Camp - Core primitives
│   ├── units/                  # Mission 1 - Unit management patterns
│   ├── resources/              # Mission 2 - Resource coordination
│   ├── battle/                 # Mission 3 - Real-time event systems
│   └── coordination/           # Mission 4 - Hierarchical coordination
│
├── pkg/utils/                   # Utility functions (minimal)
├── cmd/starcraft-war/          # Final Mission - Full application
├── examples/                    # Practice missions
│   ├── basic_combat.go         # Mission 1 example
│   └── resource_management.go  # Mission 2 example
│
├── README.md                    # Campaign overview & getting started
├── QUICK_START.md              # First 30 minutes orientation
├── CONCEPTS_GUIDE.md           # Strategy manual (MTG/SC:BW analogies)
├── PROGRESSION_MAP.md          # Tech tree & mission dependencies
└── Makefile                     # Build & learning commands

```

### Quick Start
```bash
# Setup
go version  # Verify Go 1.21+
go mod download

# Start learning
make quickstart  # Interactive setup

# Run with race detector (ALWAYS)
go run -race ./examples/basic_combat.go

# Check progress
make progress
```

---

## 2. Architecture Decision Records (ADRs)

> **Purpose**: Document key technical decisions with their rationale. Focus on WHY, not just WHAT.
>
> **Format**: Each decision includes: Context, Options Considered, Decision, Rationale, Consequences

### ADR-001: Use StarCraft & MTG as Primary Teaching Frameworks - 2024-10-02

**Context**: Learning concurrency is notoriously difficult because concepts are abstract. Students often struggle to build intuition for goroutines, channels, and synchronization primitives.

**Options Considered**:
1. **Traditional Tutorial Approach**: Explain concepts technically with code examples
   - Pros: Direct, comprehensive, industry-standard
   - Cons: Abstract, hard to build intuition, often leads to copy-paste without understanding

2. **Generic Real-World Analogies**: Use factories, restaurants, traffic lights
   - Pros: Relatable, concrete examples
   - Cons: Overused, lack depth, don't engage learners emotionally

3. **Game-Based Frameworks (Chosen)**: Use StarCraft: Brood War and Magic: The Gathering
   - Pros: Rich strategic depth, learner likely has passion for these, concepts map beautifully, emotionally engaging
   - Cons: Requires familiarity with games (but target audience likely has it)

**Decision**: Build entire teaching framework around SC:BW and MTG analogies. Every concurrency concept maps to both games.

**Rationale**:
- Both games require deep strategic thinking (like concurrency)
- SC:BW has real-time resource management, unit coordination, macro/micro—perfect parallels
- MTG has the stack, priority, state management—ideal for channels and synchronization
- Learners passionate about games will be more engaged than with generic analogies
- The analogies aren't superficial—they reveal deep structural similarities

**Consequences**:
- ✅ Learners build strong intuition through familiar frameworks
- ✅ High engagement due to passion for games
- ✅ Memorable learning (hard to forget "channels are like the MTG stack")
- ⚠️ Learners unfamiliar with games may need to learn basic concepts first
- ⚠️ Some analogies may not be perfect (trade-off for engagement)

**What You Learned**: Effective teaching connects new concepts to existing mental models. The deeper and more emotionally resonant those models, the better the learning.

---

### ADR-002: Socratic Method Over Direct Instruction - 2024-10-02

**Context**: Traditional tutorials give code and explanations. Learners copy-paste, things work, but they don't develop deep understanding or problem-solving ability.

**Options Considered**:
1. **Direct Instruction**: "Here's how mutexes work. Here's the code. Now you know."
   - Pros: Fast, comprehensive coverage, gets things working quickly
   - Cons: Shallow learning, no critical thinking, learner dependent on examples

2. **Discovery-Based with Hints**: Guide learners to discover solutions through questions
   - Pros: Deep understanding, builds problem-solving skills, memorable
   - Cons: Slower, requires more cognitive effort, some learners may struggle

3. **Socratic Method with Progressive Hints (Chosen)**: Questions first, 3-level hint system
   - Pros: Combines discovery with scaffolding, learner never truly stuck, builds confidence
   - Cons: Takes more time than direct instruction

**Decision**: Use Socratic questions throughout all code. Provide 3-level progressive hints (nudge → guidance → template) so learners can choose their support level.

**Rationale**:
- Research shows discovery-based learning leads to deeper retention
- Socratic method forces active thinking, not passive reading
- Progressive hints prevent frustration while preserving discovery
- Answering questions builds neural connections that reading doesn't
- Learners develop debugging and problem-solving skills, not just coding skills

**Consequences**:
- ✅ Deep, lasting understanding of concepts
- ✅ Learners develop independent problem-solving ability
- ✅ Can explain concepts to others (true mastery indicator)
- ⚠️ Slower progress than tutorial-style learning
- ⚠️ Requires discipline (learners must resist skipping to hints)

**What You Learned**: The struggle IS the learning. Making things too easy robs learners of the cognitive work that builds mastery.

---

### ADR-003: Minimal Code, Maximum Guidance - 2024-10-02

**Context**: Most learning projects provide either (a) complete implementations to study, or (b) empty files with TODO comments. Neither optimizes for learning.

**Options Considered**:
1. **Complete Implementations**: Full working code with extensive comments
   - Pros: Learners see professional patterns, code works immediately
   - Cons: No learning through implementation, just reading

2. **Empty Files with TODOs**: "TODO: Implement this function"
   - Pros: Forces implementation, learners write everything
   - Cons: No guidance, too intimidating, learners get stuck

3. **Rich Guidance with Templates (Chosen)**: Extensive questions, analogies, hints, starter code
   - Pros: Guided discovery, clear path forward, learner implements core logic
   - Cons: More work to create, large comment blocks

**Decision**: Remove most implementations. Replace with:
- Rich Socratic questions ("Before coding, ask yourself...")
- Game analogies explaining WHY this pattern matters
- 3-level progressive hints
- Template code with blanks to fill in
- "Pro tips" from "coaches" (Flash, Bisu, etc.)

**Rationale**:
- Cognitive load theory: provide scaffolding, not overwhelm or underwhelm
- The sweet spot is "productive struggle"—challenging but achievable
- Writing code cements understanding better than reading it
- Guidance prevents frustration, questions prevent mindless copying

**Consequences**:
- ✅ Learners implement core logic themselves
- ✅ Support prevents getting stuck indefinitely
- ✅ Questions build understanding before implementation
- ⚠️ Large comment blocks may feel overwhelming at first
- ⚠️ Some learners may want to skip straight to coding (must resist)

**What You Learned**: The "zone of proximal development"—not too easy, not too hard—is where maximum learning happens. This structure aims for that zone.

---

### ADR-004: Point-Based Rating System for Progress Tracking - 2025-10-02

**Context**: Learners need tangible feedback on their progress. Simply completing missions doesn't capture nuance—a perfectly implemented worker pool (first try, no hints) should be recognized differently than a copied implementation (with full hints). We need a system that motivates continued learning while fairly recognizing skill demonstrations.

**Options Considered**:
1. **Binary Completion Tracking**: Mission complete/incomplete checkboxes
   - Pros: Simple, clear milestones
   - Cons: No nuance, doesn't reward quality or depth, demotivating for partial progress

2. **Time-Based Tracking**: Track hours spent on each concept
   - Pros: Shows effort invested
   - Cons: Rewards time, not understanding; slow learners penalized

3. **Point-Based Rating System (Chosen)**: Award points for implementations with decay mechanics
   - Pros: Granular feedback, rewards first-time learning, recognizes complexity, motivates progress
   - Cons: More complex to track, requires point calibration

**Decision**: Implement ELO-style point-based rating system with:
- Starting rank: E (1150 points)
- Rank brackets: F(0-1149), E(1150-1424), D(1425-1549), C(1550-1699), B(1700-1999), A(2000-2499), S(2500+)
- Point values based on implementation complexity (10-600 points)
- Decay mechanics for repetitions (~15% per iteration, floor at 25%)
- Full points restored for new complexity facets
- Bonus multipliers for mastery demonstrations (+10-25%)

**Rationale**:
- **SC:BW Ladder Analogy**: Learners already understand rank progression from gaming
- **Decay Encourages Breadth**: Can't grind same pattern repeatedly; must learn new things
- **New Complexity Recognition**: Rewards genuine learning (new patterns) vs. repetition
- **Bonus System**: Recognizes quality (first attempt success, no hints, teaching back)
- **Granular Feedback**: Every implementation moves the needle, shows progress
- **Motivation**: Clear path from E → S rank creates compelling progression

**Consequences**:
- ✅ Tangible progress feedback after every implementation
- ✅ System rewards understanding over completion
- ✅ Prevents grinding same pattern for points (decay)
- ✅ Motivates learning breadth (new patterns = full points)
- ✅ Recognizes mastery (bonuses for teaching, optimization, etc.)
- ⚠️ Requires manual point tracking (could automate later)
- ⚠️ Point values may need calibration based on actual learner experience
- ⚠️ Risk of "gaming the system" (mitigated by decay and complexity detection)

**What You Learned**: Effective progress tracking must balance granularity (see small wins) with meaningfulness (points reflect real skill). Game-based metaphors (ladder ranks) leverage existing mental models for motivation.

---

---

## 3. Critical Concepts

> **Purpose**: Track what you're learning in this project. These are concepts to review and master.
>
> **Update**: Add new concepts as you learn them. Move to "Mastered" section after 3+ successful applications.

### Currently Learning

#### Concept: Goroutines - Lightweight Concurrency
**What It Is**: Goroutines are lightweight threads managed by the Go runtime. They're cheap to create (thousands or millions possible) and scheduled cooperatively.

**Why It Matters**: Foundation of all concurrency in Go. Understanding goroutines is like understanding unit production in SC:BW or casting creatures in MTG—you can't play the game without it.

**How We Use It**:
```go
// In types.go and throughout the project
go unit.run(ctx) // Spawn unit's lifecycle goroutine
// Each unit runs independently, like training a Marine

var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // Do concurrent work
}()
wg.Wait() // Wait for all to finish
```

**Key Insight**: Goroutines are not free—they cost memory and scheduling time. Like SC:BW supply limits, you can't just spawn infinite units. Design matters.

**Related To**:
- OS threads (but much lighter)
- SC:BW unit production (train → unit exists independently)
- MTG creature spells (cast → resolves → on battlefield)

**Practice Opportunities**:
- [ ] Spawn 1000 goroutines and observe memory usage
- [ ] Create goroutine leak and detect it
- [ ] Implement proper cleanup with context

---

#### Concept: Channels - Communication Primitive
**What It Is**: Typed conduits for sending/receiving values between goroutines. Can be buffered (queue) or unbuffered (synchronization point).

**Why It Matters**: Go's philosophy: "Don't communicate by sharing memory, share memory by communicating." Channels are THE way goroutines coordinate.

**How We Use It**:
```go
// Unbuffered: synchronous handoff
commandChan := make(chan Command)
go func() {
    cmd := <-commandChan // Blocks until sender
}()
commandChan <- AttackCommand // Blocks until receiver

// Buffered: asynchronous up to capacity
eventChan := make(chan Event, 100) // Holds 100 events
eventChan <- event // Doesn't block until full
```

**Key Insight**: Channels are like the MTG stack—things resolve in order, timing matters. Unbuffered = instant priority pass. Buffered = stack with capacity.

**Related To**:
- Queues (but with goroutine synchronization)
- MTG stack (LIFO, ordered resolution)
- SC:BW command queue (orders execute in sequence)

**Practice Opportunities**:
- [ ] Build pipeline with 3 stages using channels
- [ ] Experiment with buffer sizes and observe behavior
- [ ] Create deadlock with channels, then fix it

---

#### Concept: Mutexes - Shared State Protection
**What It Is**: Mutual exclusion locks. Only one goroutine can hold a lock at a time. Protects shared data from concurrent access.

**Why It Matters**: Prevents data races—the bane of concurrent systems. Like controlling who can mine a mineral patch in SC:BW.

**How We Use It**:
```go
type SafeCounter struct {
    mu    sync.RWMutex
    count map[string]int
}

func (s *SafeCounter) Inc(key string) {
    s.mu.Lock()         // Take control
    defer s.mu.Unlock() // Always release
    s.count[key]++
}

func (s *SafeCounter) Value(key string) int {
    s.mu.RLock()        // Multiple readers OK
    defer s.mu.RUnlock()
    return s.count[key]
}
```

**Key Insight**: Mutexes protect DATA, not code. Minimize critical section (time holding lock). Like MTG priority—only hold it while you need it.

**Related To**:
- Semaphores (but simpler)
- MTG priority (only one player acts at a time)
- SC:BW resource patch (only one worker mines at once)

**Practice Opportunities**:
- [ ] Create data race, detect with `-race`, fix with mutex
- [ ] Compare sync.Mutex vs sync.RWMutex performance
- [ ] Implement deadlock through lock ordering, then fix

---

### Mastered Concepts

*(Concepts move here after 5+ applications and can be explained fluently to others)*

*None yet—this is where your conquered concepts will live!*

---

## 4. Current Status

> **Purpose**: Quick snapshot of where the project stands RIGHT NOW.
>
> **Update**: Every session end. Keep this current, not historical.

### What's Working
- ✅ Project structure established with clear mission progression
- ✅ Learning system documentation (README, QUICK_START, CONCEPTS_GUIDE, PROGRESSION_MAP)
- ✅ Memory system files created (.project/ directory)
- ✅ Go module initialized (go.mod)
- ✅ Makefile with basic commands

### What's In Progress
- 🚧 Converting code files from TODO-heavy to Socratic-question-rich
- 🚧 Implementing discovery-based learning throughout codebase
- 🚧 Adding MTG/SC:BW analogies to all code comments

### Known Issues
- ⚠️ Most code files return nil/0—this is intentional (learner implements)
- ⚠️ Code won't compile until learner implements functions—by design
- ⚠️ No tests yet—will add as part of learning exercises

### What's Next (Priority Order)
1. **Boot Camp**: Revamp `internal/types/types.go` with rich Socratic guidance
2. **Mission 1**: Revamp `internal/units/manager.go` with worker pool patterns
3. **Mission 2**: Revamp `internal/resources/manager.go` with resource coordination
4. **Examples**: Convert example files to interactive missions
5. **Makefile**: Add learning-support commands

---

## 5. Learning Goals

> **Purpose**: Track what you want to master through this project.
>
> **Review**: Check progress monthly. Celebrate wins, adjust goals as needed.

### Active Learning Goals

#### Goal: Master Go Concurrency Primitives
**Why This Matters**: Can't build production systems without deep understanding of goroutines, channels, and synchronization.

**Success Criteria**:
- [ ] Can explain goroutines, channels, mutexes using analogies to others
- [ ] Can identify data races before the race detector finds them
- [ ] Can design concurrent systems from scratch with confidence
- [ ] Can debug deadlocks and goroutine leaks systematically
- [ ] Can teach these concepts to someone else

**Progress**: Starting Boot Camp—foundation level

**Resources**:
- CONCEPTS_GUIDE.md (this project)
- Go Concurrency Patterns (Go blog)
- Effective Go (official docs)

---

#### Goal: Develop Problem-Solving Through Discovery
**Why This Matters**: Copy-paste makes code work; understanding makes engineers. Discovery builds deep knowledge.

**Success Criteria**:
- [ ] Can solve problems without immediately looking for solutions
- [ ] Can break complex problems into manageable pieces
- [ ] Can debug by forming hypotheses and testing them
- [ ] Comfortable with "productive struggle"
- [ ] Can explain my reasoning process to others

**Progress**: Learning the Socratic workflow

---

### Completed Learning Goals

*(Goals move here when all success criteria met)*

*None yet—this is your trophy room!*

---

## 6. Code Patterns & Conventions

> **Purpose**: Document project-specific patterns and standards.
>
> **Why**: Ensures consistency and helps you learn good practices.

### Naming Conventions
- **Files**: [e.g., kebab-case for components, PascalCase for classes]
- **Variables**: [e.g., camelCase for variables, UPPER_CASE for constants]
- **Functions**: [e.g., verb-based names like getUserData()]

### Code Patterns

#### Pattern: [Pattern Name]
**When to Use**: [Situation where this pattern applies]

**Template**:
```[language]
// Standard way we implement this in the project
[code template]
```

**Example**:
```[language]
// Real example from your codebase
[actual code]
```

---

## 7. Key Resources

> **Purpose**: Quick links to documentation and resources you reference often.

### Documentation
- [Technology/Framework]: [URL]
- [Another key doc]: [URL]

### Learning Resources
- [Tutorial/Course you're following]: [URL]
- [Helpful article/video]: [URL]

### Project-Specific
- [Design doc]: [Path or URL]
- [API documentation]: [Path or URL]

---

## 8. Collaboration Notes

> **Purpose**: Track important context about how you work with Claude.

### Communication Preferences
- **Explanation Style**: Use MTG and SC:BW analogies heavily; think like Flash, explain like Bisu
- **Code Review Focus**: Explain WHY patterns work, not just HOW to implement them
- **Learning Pace**: Moderate—let me think and struggle, but provide hints when truly stuck
- **Question Depth**: Ask 2-3 Socratic questions before providing implementation guidance

### Teaching Configuration
- **Scaffolding Level**: Medium—challenge me but don't let me flounder
- **Explanation Detail**: Comprehensive for new concepts, concise for review
- **Practice Frequency**: High—create exercises and checkpoints frequently
- **Review Intensity**: Strict spaced repetition—remind me to review concepts

### Context Boundaries
- **What Claude Should Know**:
  - This is a learning project, not production code
  - Struggle and mistakes are expected and valuable
  - Speed is not the goal; depth of understanding is
  - I want to implement code myself, not have it written for me

- **What Claude Shouldn't Assume**:
  - Don't assume I've mastered a concept just because I implemented it once
  - Don't assume I want the "fast path"—I want to learn deeply
  - Don't provide solutions before I've had a chance to think
  - Always check: "Have you thought about X?" before giving answers

---

## 9. Session Quick Reference

> **Purpose**: Quick links to recent sessions and active context.

### Recent Sessions
- [Latest date]: [One-line summary] → See `/SESSION_HISTORY/YYYY-MM-DD.md`
- [Previous date]: [One-line summary] → See `/SESSION_HISTORY/YYYY-MM-DD.md`
- [Earlier date]: [One-line summary] → See `/SESSION_HISTORY/YYYY-MM-DD.md`

### Active Context
📍 **Current Focus**: [See ACTIVE_CONTEXT.md for details on current work]

---

## 10. Maintenance

### Archive Checklist (Monthly Review)
- [ ] Move ADRs older than 3 months to `/archive/decisions/`
- [ ] Move mastered concepts that haven't been used in 2 months to `/archive/concepts/`
- [ ] Archive session history older than 1 month to `/archive/sessions/`
- [ ] Update "What's Working" to reflect actual current state
- [ ] Review and update learning goals

### File Size Check
- **Current Line Count**: [Run `wc -l CLAUDE.md` and note here]
- **Target**: Keep under 600 lines
- **Action If Over**: Archive older content to `/archive/`

---

## Template Usage Instructions

**For New Projects:**
1. Copy this template to your project root
2. Fill in Section 1 (Project Overview) completely
3. Add your first ADR when you make your first architectural decision
4. Add concepts to Section 3 as you learn them
5. Update Section 4 (Current Status) at the end of each session
6. Set your learning goals in Section 5

**For Existing Projects:**
1. Fill in current project state in Section 1
2. Add recent architectural decisions as ADRs (last 3 months)
3. List concepts you're currently learning or mastering
4. Set current status and immediate next steps
5. Backfill history lightly - focus on current state over past details

**Daily Workflow:**
- **Session Start**: Claude reads this file + ACTIVE_CONTEXT.md (~700 lines total)
- **During Session**: Update as you make decisions or learn concepts
- **Session End**: Quick update to Current Status, create session log

**Remember**: This file is Claude's memory. Keep it current, concise, and focused on what matters for moving forward.
