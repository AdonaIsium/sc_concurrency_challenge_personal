# 🚀 Quick Start - Your First 30 Minutes

> *"Before Flash became the Ultimate Weapon, he spent hours studying basic mechanics. Before you master concurrency, you need to understand how this learning system works."*

Welcome! You've found this project because you want to **truly understand** Go concurrency, not just copy-paste it. Smart choice. This guide gets you oriented in 30 minutes.

---

## 🎯 What You're About to Experience

This isn't a traditional tutorial. This is a **discovery-based learning adventure** using two powerful frameworks:

### 1. The Socratic Method (Questions > Answers)
Instead of: *"Here's how to use a mutex"*
You'll get: *"Before using a mutex, ask yourself: what happens if two goroutines access this at once?"*

**Why?** Because answering questions builds **deep understanding**. Just reading answers builds **shallow knowledge**.

### 2. Game-Based Analogies (Familiar > Abstract)
Instead of: *"Channels provide communication between goroutines"*
You'll get: *"Channels are like the MTG stack—things resolve in order, priority matters, and timing is everything"*

**Why?** Because your brain already understands StarCraft army control and MTG card interactions. We're just mapping those concepts to concurrency.

---

## 📚 The Three Memory Files (Your Training Regimen)

Before you start coding, set up your learning infrastructure. Think of this as building your training facility before ladder games.

### `.project/CLAUDE.md` - Your Strategic Playbook

**What it is**: Long-term memory for architectural decisions, concepts, and project state

**MTG Analogy**: This is your deck tech and sideboard guide
**SC:BW Analogy**: This is your build order library and strategy notes

**What to do NOW**:
1. Open `.project/CLAUDE.md`
2. Fill in "Project Overview" section (5 minutes)
3. Add your first learning goal: "Master Go concurrency patterns"
4. Set communication preferences (how much detail you want)

```markdown
## 1. Project Overview
### Project Name
StarCraft Concurrency War - My Learning Journey

### Description
A hands-on project to master Go concurrency through discovery-based learning,
using StarCraft: Brood War and Magic: The Gathering as teaching frameworks.

### My Learning Goals
- Understand channels, mutexes, and goroutines at a deep level
- Learn when to use each concurrency primitive
- Build confidence in designing concurrent systems from scratch

## 8. Collaboration Notes
### Communication Preferences
- Explanation Style: Use MTG and SC:BW analogies heavily
- Code Review Focus: Explain WHY, not just HOW
- Learning Pace: Moderate - let me think before giving hints
```

### `.project/ACTIVE_CONTEXT.md` - Your Current Match

**What it is**: Working memory for what you're doing RIGHT NOW

**MTG Analogy**: This is the current game state—what's in play, what's in your hand
**SC:BW Analogy**: This is your current match—supply count, army position, build progress

**What to do NOW**:
1. Open `.project/ACTIVE_CONTEXT.md`
2. Set your current sprint goal
3. Add your first task

```markdown
## 🎯 Current Sprint/Focus
**Sprint Goal**: Complete Boot Camp - Master the core types and basic patterns
**Target Completion**: Next 3 sessions
**Why This Matters**: Can't build concurrent systems without understanding the foundation

## 🚧 Work In Progress
### Primary Task: Understanding Go Concurrency Primitives
**Status**: Just started
**Current Step**: Reading types.go to understand the structure
**Progress**:
- 🔄 Review types.go and understand the design
- ⏳ Implement String() methods
- ⏳ Implement NewUnit() constructor
- ⏳ Complete basic getters/setters

## ⏭️ Next Steps (Priority Order)
### Immediate (This Session):
1. Read through types.go completely
2. Answer the Socratic questions in comments
3. Implement 2-3 simple methods
```

### `.project/LEARNING_TRACKER.md` - Your Training Schedule

**What it is**: Spaced repetition system for mastering concepts

**MTG Analogy**: Your practice schedule—when to review different deck matchups
**SC:BW Analogy**: Your training regimen—when to drill different build orders

**What to do NOW**:
1. Open `.project/LEARNING_TRACKER.md`
2. You'll add concepts as you learn them
3. The system will schedule reviews automatically

*Don't worry about this file yet—you'll populate it as you learn.*

---

## 🎮 Your First Mission: Boot Camp

Now that your training facility is set up, time for your first mission.

### Step 1: Orient Yourself (5 minutes)

```bash
# Navigate to the project
cd starcraft_concurrency_war_claude

# See the structure
ls -la

# Key directories:
# internal/types/     ← Start here (Boot Camp)
# internal/units/     ← Mission 1
# internal/resources/ ← Mission 2
# internal/battle/    ← Mission 3
# internal/coordination/ ← Mission 4
# examples/           ← Practice missions
```

### Step 2: Read the Foundation (10 minutes)

Open `internal/types/types.go` and **just read**. Don't code yet.

**What to notice**:
1. The comments are LONG—that's intentional. They're teaching you.
2. You'll see questions like *"🤔 BEFORE YOU CODE:"* - These are crucial
3. You'll see analogies to MTG and SC:BW—use them to build intuition
4. Most functions return `nil` or `0`—you'll implement them

**Pro Tip**: Use `less internal/types/types.go` or your favorite editor. Read from top to bottom once.

### Step 3: Answer Your First Questions (10 minutes)

After reading, before writing any code, answer these questions **out loud or in writing**:

**Question 1: Goroutines**
> *In StarCraft, when you train a Marine, it exists independently and can be controlled. In MTG, when you cast a creature, it resolves and sits on the battlefield. What's the Go equivalent?*

**Question 2: Communication**
> *If two Marines need to coordinate an attack, they communicate. If two creatures need to trigger an ability, the stack handles it. In Go, when two goroutines need to coordinate, what do they use?*

**Question 3: Protection**
> *In StarCraft, only one player can control a unit at a time. In Go, if multiple goroutines try to update the same variable simultaneously, what prevents chaos?*

**Write your answers down in ACTIVE_CONTEXT.md under "Session Notes & Insights"**

These questions aren't tests—they're setting up your mental framework.

### Step 4: Implement Something Small (15 minutes)

Now you're ready to code. Start with the **simplest** function in `types.go`:

**Your first implementation: `String()` methods**

Find the `UnitType.String()` method. Notice:
- It's already implemented! (Good news)
- Study HOW it works
- See the pattern: map lookup, default case

**Your turn**: Look at other String() methods in the file. Are they all implemented?

**Implementation Challenge**: Find 2-3 String() methods that need implementation and complete them using the same pattern.

```go
// Example you'll find:
func (us UnitState) String() string {
    if name, ok := unitStateNames[us]; ok {
        return name
    }
    return fmt.Sprintf("UnitState(%d)", us)
}
```

**🎯 Success Check**: Can you run `go build ./internal/types` without errors?

---

## 🧠 The Socratic Workflow (Use This Every Session)

This is **how you learn** with this project:

### Phase 1: Read & Understand (Don't Code Yet!)
1. Open the file you're working on
2. Read all the comments—they're the curriculum
3. Read the questions (🤔 markers)
4. Form hypotheses about the answers

### Phase 2: Think & Answer
1. Answer the Socratic questions in your head (or out loud!)
2. Write your thoughts in ACTIVE_CONTEXT.md
3. Check if your reasoning makes sense
4. Look for analogies to MTG/SC:BW to solidify understanding

### Phase 3: Implement & Test
1. Write the minimal code to solve the problem
2. Test immediately: `go run -race ./your_file.go`
3. If it works, ask yourself: *"WHY does this work?"*
4. If it fails, ask yourself: *"What did I misunderstand?"*

### Phase 4: Review & Reinforce
1. Add the concept to LEARNING_TRACKER.md
2. Update ACTIVE_CONTEXT.md with progress
3. Schedule your review (the system will remind you)
4. Move to the next challenge

**Critical**: Never skip Phase 1 and 2. The questions BUILD understanding. The code DEMONSTRATES understanding.

---

## 🎯 Hint System (When You're Stuck)

Throughout the code, you'll find three levels of hints:

```go
// 🎯 HINT LEVEL 1: Gentle nudge in the right direction
// 🎯 HINT LEVEL 2: Stronger guidance with patterns
// 🎯 HINT LEVEL 3: Template with blanks to fill in
```

**Use them progressively**:
1. Try Level 1 first, think for 5+ minutes
2. If still stuck, try Level 2
3. Only use Level 3 if you're genuinely blocked

**Pro Tip**: Getting stuck is GOOD. That's when learning happens. Don't rush to hints.

---

## ⚠️ Common Beginner Mistakes

### Mistake 1: "I'll just skip the reading and code"
**Result**: Code that works but you don't understand
**Fix**: Force yourself to read and think first. Set a timer: 10 min reading, then code.

### Mistake 2: "This analogy is silly, I'll ignore it"
**Result**: Abstract concepts stay abstract
**Fix**: Embrace the analogies. Your brain LOVES familiar frameworks. Use them.

### Mistake 3: "I'll do all the missions in one session"
**Result**: Cognitive overload, shallow learning
**Fix**: One mission per session (2-4 hours). Take breaks. Let concepts marinate.

### Mistake 4: "I don't need to write things down"
**Result**: Forgotten concepts, no spaced repetition
**Fix**: Update your memory files after EVERY session. Future you will thank current you.

### Mistake 5: "I'll copy-paste from Stack Overflow"
**Result**: You learn nothing
**Fix**: Struggle is learning. Use hints in the code, not external solutions.

---

## 📅 Session Template (Copy This!)

At the **start of each session**:

```markdown
## Session [Date] - [What I'm Working On]

### Goals for Today:
1. [Primary goal]
2. [Secondary goal]

### Questions I Want to Answer:
- [Question 1]
- [Question 2]

### Progress Checklist:
- [ ] Read relevant files
- [ ] Answer Socratic questions
- [ ] Implement 2-3 functions
- [ ] Test with `go run -race`
- [ ] Update learning tracker
```

At the **end of each session**:

```markdown
### What I Accomplished:
- ✅ [Thing 1]
- ✅ [Thing 2]

### What I Learned:
- [Concept 1] - Aha moment: [insight]
- [Concept 2] - Connected to: [previous knowledge]

### Questions That Arose:
- [Question to explore next time]

### Next Session:
- Start with: [specific file/function]
- Review before starting: [concept to refresh]
```

**Save this in SESSION_HISTORY/** with the date: `YYYY-MM-DD-session.md`

---

## 🚦 Green Lights (You're Ready When...)

You're ready to start Boot Camp when:
- ✅ You've read this entire Quick Start guide
- ✅ You've set up your three memory files
- ✅ You understand the Socratic workflow
- ✅ You're mentally prepared to THINK before coding
- ✅ You have 2-4 hours blocked for focused learning

You're ready for Mission 1 when:
- ✅ Boot Camp is complete (types.go has several working methods)
- ✅ You can explain goroutines, channels, and mutexes using analogies
- ✅ You've successfully run code with `-race` flag
- ✅ You've added concepts to your learning tracker
- ✅ You're excited (not overwhelmed) to continue

---

## 🆘 Emergency Procedures

### "I'm Completely Lost"
1. Stop coding
2. Read CONCEPTS_GUIDE.md for the concept you're stuck on
3. Reread the questions in the code
4. Ask yourself: "What concept am I missing?"
5. Review your learning tracker—did you skip a prerequisite?

### "The Code Won't Compile"
```bash
# Run this to see detailed errors:
go build -v ./...

# Common issues:
# - Missing imports: go mod tidy
# - Syntax errors: read the error message carefully
# - Type mismatches: check the function signature
```

### "I Don't Get the Analogy"
That's okay! Not everyone plays MTG or SC:BW. Key concepts:
- **MTG Stack**: Things happen in order, last-in-first-out
- **SC:BW Macro**: Managing multiple things at once
- **SC:BW Timing Attack**: Everything happens in real-time
- **Resource Contention**: Multiple units want the same thing

If analogies don't help, focus on the concrete Go concepts. The analogies are helpers, not requirements.

### "I'm Not Learning Fast Enough"
**STOP RIGHT THERE.**

Learning speed ≠ learning depth. Ask yourself:
- Can you explain what you learned to someone else?
- Can you apply it to a new problem?
- Do you understand WHY, not just HOW?

If yes, you're learning at exactly the right speed.

---

## 🎯 Your 30-Minute Summary

**You've now learned**:
1. ✅ This project uses Socratic method + game analogies
2. ✅ Three memory files track your learning journey
3. ✅ The workflow: Read → Think → Code → Review
4. ✅ Hints exist in 3 levels—use progressively
5. ✅ Your first mission: Boot Camp in `internal/types/types.go`
6. ✅ Success means understanding WHY, not just making it work

**Your immediate next steps**:
1. Set up your three memory files (10 minutes)
2. Read `internal/types/types.go` completely (15 minutes)
3. Implement your first String() method (15 minutes)
4. Update ACTIVE_CONTEXT.md with progress (5 minutes)

**Total time to first code running**: ~45 minutes

---

## 🔥 Motivational Wisdom

### From Flash
> "I didn't become the best by copying builds. I became the best by understanding why builds work. That's what you're doing here—learning the 'why' behind concurrency."

### From Jon Finkel
> "The best players don't memorize every card interaction. They understand the rules so deeply that interactions become obvious. Same with concurrency—understand the primitives deeply, and patterns become obvious."

### From This Project
> "If you're struggling, you're learning. If it's too easy, you're not learning enough. Find the sweet spot where challenge meets capability. That's the zone of masimal growth."

---

## 🎮 Ready?

**You are now prepared to start your journey.**

Open `internal/types/types.go` and begin Boot Camp. Remember:
- Read everything
- Answer the questions
- Think before coding
- Test frequently
- Update your tracker
- Embrace the struggle

**The concurrency arena awaits. GG HF!** 🚀

---

*Next recommended reading: CONCEPTS_GUIDE.md (for deep dives on specific concepts)*
*Having issues? Check your ACTIVE_CONTEXT.md and see if you skipped a step*
