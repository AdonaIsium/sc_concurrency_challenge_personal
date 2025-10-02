package types

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════
// 🎮 BOOT CAMP - CORE TYPES
// ═══════════════════════════════════════════════════════════════════════════
//
// Welcome to Boot Camp! This file is your training ground for Go concurrency fundamentals.
//
// 🎯 MISSION OBJECTIVES:
// 1. Understand how to model concurrent entities (Units = Goroutines)
// 2. Learn state management with mutexes (protecting shared data)
// 3. Practice communication through channels (how units coordinate)
// 4. Master lifecycle management with context (graceful shutdown)
//
// 🎴 MTG ANALOGY:
//   - Unit = Creature card
//   - State = Card state (tapped/untapped, +1/+1 counters, etc.)
//   - Commands = Instants/Sorceries targeting creatures
//   - Lifecycle = Battlefield → Graveyard
//
// ⚔️ SC:BW ANALOGY:
//   - Unit = Marine, Zealot, Zergling
//   - State = Idle, Moving, Attacking, Dead
//   - Commands = Move orders, attack orders
//   - Lifecycle = Train → Exists → Dies
//
// ⚠️ CRITICAL CONCEPTS YOU'LL LEARN:
//   - Why mutexes are needed (data races are the enemy!)
//   - When to use channels vs shared memory
//   - How context enables graceful shutdown
//   - Proper goroutine lifecycle management
//
// 📚 Read CONCEPTS_GUIDE.md before implementing if anything feels unclear.
//
// ═══════════════════════════════════════════════════════════════════════════

// ═══════════════════════════════════════════════════════════════════════════
// SECTION 1: ENUMS & CONSTANTS
// ═══════════════════════════════════════════════════════════════════════════
//
// 🎓 LEARNING MOMENT: Go doesn't have enums like other languages.
// Instead, we use custom types with iota constants. This pattern is
// everywhere in Go codebases—master it here!
//
// ═══════════════════════════════════════════════════════════════════════════

// UnitType represents different unit types in our StarCraft simulation
//
// 🎴 MTG: Like creature types (Goblin, Elf, Dragon)
// ⚔️ SC:BW: Marine, Zealot, Zergling, etc.
//
// 🤔 BEFORE YOU CODE:
// Q1: Why use a custom type (UnitType int) instead of just string?
// Q2: What does iota do in Go constants?
// Q3: How would you add a new unit type without breaking existing code?
//
// 🎯 HINT LEVEL 1: Type safety! Compiler prevents you from assigning "Bob" to a UnitType
// 🎯 HINT LEVEL 2: iota auto-increments: Marine=0, Zergling=1, Zealot=2, etc.
// 🎯 HINT LEVEL 3: Add it before unitTypeCount, update unitTypeNames map
type UnitType int

const (
	SCV UnitType = iota
	Marine
	Firebat
	Medic
	Vulture
	SiegeTank
	Goliath
	Wraith
	DropShip
	Valkyrie
	ScienceVessel
	Battlecruiser
	Drone
	Overlord
	Zergling
	Hydralisk
	Lurker
	Mutalisk
	Guardian
	Devourer
	Queen
	Ultralisk
	Defiler
	Probe
	Zealot
	Dragoon
	Templar
	DarkTemplar
	Shuttle
	Reaver
	Observer
	Corsair
	Carrier
	Arbiter
	// 🏗️ YOUR CHALLENGE: Add 2-3 more unit types here
	// Consider: Mutalisk, Carrier, Firebat, Dragoon
	// Remember to update unitTypeNames map below!

	unitTypeCount // Sentinel value for validation
)

// Compiler check: Ensure UnitType implements fmt.Stringer interface
// 🎓 LEARNING: The _ = syntax means "compile-time check, no runtime cost"
var _ fmt.Stringer = UnitType(0)

// unitTypeNames maps enum values to human-readable strings
// 🎓 LEARNING: This pattern makes debugging WAY easier
var unitTypeNames = map[UnitType]string{
	SCV:           "SCV",
	Marine:        "Marine",
	Firebat:       "Firebat",
	Medic:         "Medic",
	Vulture:       "Vulture",
	SiegeTank:     "SiegeTank",
	Goliath:       "Goliath",
	Wraith:        "Wraith",
	DropShip:      "DropShip",
	Valkyrie:      "Valkyrie",
	ScienceVessel: "ScienceVessel",
	Battlecruiser: "Battlecruiser",
	Drone:         "Drone",
	Overlord:      "Overlord",
	Zergling:      "Zergling",
	Hydralisk:     "Hydralisk",
	Lurker:        "Lurker",
	Mutalisk:      "Mutalisk",
	Guardian:      "Guardian",
	Devourer:      "Devourer",
	Queen:         "Queen",
	Ultralisk:     "Ultralisk",
	Defiler:       "Defiler",
	Probe:         "Probe",
	Zealot:        "Zealot",
	Dragoon:       "Dragoon",
	Templar:       "Templar",
	DarkTemplar:   "DarkTemplar",
	Shuttle:       "Shuttle",
	Reaver:        "Reaver",
	Observer:      "Observer",
	Corsair:       "Corsair",
	Carrier:       "Carrier",
	Arbiter:       "Arbiter",
	// 🏗️ YOUR CHALLENGE: Add entries for your new unit types
}

// String implements the fmt.Stringer interface for better debugging
//
// 🤔 WHY THIS MATTERS:
// When you fmt.Println(unit.Type), you get "Marine" not "UnitType(0)"
// When debugging with logs, readable output = faster bug fixes
//
// 🔥 PRO TIP FROM FLASH:
// "In SC:BW replays, we need to see unit names, not unit IDs. Same in code—
//
//	readable debugging output saves hours of head-scratching."
//
// ✅ ALREADY IMPLEMENTED (Study the pattern—you'll use it everywhere!)
func (ut UnitType) String() string {
	if name, ok := unitTypeNames[ut]; ok {
		return name
	}
	// Fallback for unknown types (defensive programming)
	return fmt.Sprintf("UnitType(%d)", ut)
}

// IsValid checks if a UnitType value is within valid range
//
// 🤔 BEFORE YOU CODE:
// Q1: Why do we need validation? Can't we trust the code?
// A2: We can trust the code to be as good as it is written!
//
//	However, there are always possibilities for bugs or unexpected
//	states to occur so we want to make sure to implement checks and
//	balances to test our work.
//
// Q2: What if someone does UnitType(999)?
// A2: Since our list of units is < 999, that would not be valid and would
//
//	return false.
//
// Q3: How does this prevent bugs in production?
// A3: It stops our code from ending up in a state that it should not be in
//
//	which ensures that we can reliably take any actions on it since its
//	state is always a known value
//
// 🎯 HINT LEVEL 1: Type conversion can bypass const safety
// 🎯 HINT LEVEL 2: ut >= Marine (minimum) && ut < unitTypeCount (maximum)
// 🎯 HINT LEVEL 3: See the pattern in String() above
//
// ✅ ALREADY IMPLEMENTED (Defensive programming—always validate!)
func (ut UnitType) IsValid() bool {
	return ut >= SCV && ut < unitTypeCount
}

// ═══════════════════════════════════════════════════════════════════════════
// UnitState - State Machine Pattern
// ═══════════════════════════════════════════════════════════════════════════
//
// 🎴 MTG ANALOGY: Card states (summoning sickness, tapped, phased out, etc.)
// ⚔️ SC:BW ANALOGY: Unit behaviors (idle, moving, attacking, dead)
//
// 🎓 LEARNING: State machines are EVERYWHERE in concurrent systems.
// A unit can be in one state at a time. Transitions between states must be
// carefully controlled to prevent race conditions.
//
// 🤔 THINK ABOUT THIS:
// In MTG: Can a creature attack while tapped? No—invalid state transition.
// In our code: Can a dead unit attack? No—same principle!
//
// ═══════════════════════════════════════════════════════════════════════════

type UnitState int

const (
	Idle UnitState = iota
	Moving
	Attacking
	Defending
	HoldingPosition
	Patrolling
	Repairing
	Building
	Dead
	// 🏗️ YOUR CHALLENGE: Add more states
	// Ideas: Retreating, Repairing, Building, Casting (for abilities)
	// Consider: What state transitions make sense?
	//   Idle → Moving ✓
	//   Dead → Attacking ✗ (invalid!)

	unitStateCount
)

var _ fmt.Stringer = UnitState(0)

var unitStateNames = map[UnitState]string{
	Idle:            "Idle",
	Moving:          "Moving",
	Attacking:       "Attacking",
	Defending:       "Defending",
	HoldingPosition: "HoldingPosition",
	Patrolling:      "Patrolling",
	Repairing:       "Repairing",
	Building:        "Building",
	Dead:            "Dead",
	// 🏗️ YOUR CHALLENGE: Add names for your new states
}

// 🏗️ YOUR CHALLENGE: Implement String() method for UnitState
// Follow the exact pattern from UnitType.String() above
//
// 🤔 BEFORE YOU CODE:
// Q: Why is copy-paste sometimes OK in programming?
// A: Consistent patterns > clever variations. Make it obvious, not clever.
//
// 🎯 HINT LEVEL 1: Look at UnitType.String() above—it's identical logic
// 🎯 HINT LEVEL 2: Change "unitTypeNames" to "unitStateNames"
// 🎯 HINT LEVEL 3: Here's the template:

func (us UnitState) String() string {
	if name, ok := unitStateNames[us]; ok {
		return name
	}
	return fmt.Sprintf("UnitState(%d)", us)
}

// 🏗️ YOUR CHALLENGE: Implement IsValid() method for UnitState
// Follow the pattern from UnitType.IsValid()
//
// 🎯 HINT: Replace UnitType with UnitState, Marine with Idle
func (us UnitState) IsValid() bool {
	return us >= Idle && us < unitStateCount
}

// ═══════════════════════════════════════════════════════════════════════════
// SECTION 2: VALUE TYPES (NO CONCURRENCY... YET)
// ═══════════════════════════════════════════════════════════════════════════
//
// 🎓 LEARNING: Not everything needs concurrency! Simple data types like
// Position don't need mutexes or channels. They're just data.
//
// 🤔 WHEN DO WE NEED CONCURRENCY PROTECTION?
// ✓ Multiple goroutines reading AND writing? → Need mutex or channel
// ✓ Shared state that changes? → Need synchronization
// ✗ Read-only data? → No protection needed
// ✗ Short-lived local variables? → No protection needed
//
// ═══════════════════════════════════════════════════════════════════════════

// Position represents a unit's location on the battlefield
//
// 🎴 MTG: Not applicable (no spatial positions in MTG)
// ⚔️ SC:BW: Exact pixel coordinates on the map
//
// 🤔 DESIGN DECISION:
// Q: Why float64 instead of int for coordinates?
// A: Smooth movement, precise calculations. int would cause choppy motion.
//
// Q: Do we need Z coordinate?
// A: Not for SC:BW (2D game), but useful for elevation/cliffs
type Position struct {
	X float64
	Y float64
	// 🏗️ OPTIONAL: Add Z float64 if you want 3D or elevation
}

// Distance calculates the Euclidean distance between two positions
//
// 🤔 BEFORE YOU CODE:
// Q1: What's the formula for distance between two points?
// Q2: Why use math.Sqrt? Can we avoid it for performance?
// Q3: What's "Euclidean distance" vs "Manhattan distance"?
//
// 🎯 HINT LEVEL 1: Pythagorean theorem: a² + b² = c²
// 🎯 HINT LEVEL 2: sqrt((x2-x1)² + (y2-y1)²)
// 🎯 HINT LEVEL 3: Use math.Sqrt and math.Pow (or just multiply dx*dx)
//
// 🏗️ YOUR CHALLENGE: Implement this method
// Template:
/*
func (p Position) Distance(other Position) float64 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	// Calculate and return distance using Pythagorean theorem
	return 0.0 // Replace this!
}
*/

// DistanceSquared returns the squared distance (faster, no sqrt)
//
// 🔥 PRO TIP FROM FLASH:
// "In SC:BW, we compare distances constantly (attack range, vision, etc.).
//  Square root is expensive! If you only need to compare distances,
//  compare the SQUARED distances instead. Way faster."
//
// 🤔 WHY THIS MATTERS:
// Comparing: dist1 < dist2 is same as: dist1² < dist2²
// But dist1² is faster to calculate (no sqrt!)
//
// 🏗️ YOUR CHALLENGE: Implement DistanceSquared
// It's like Distance() but without the math.Sqrt() call
//
// 🎯 HINT: Just return dx*dx + dy*dy

// ═══════════════════════════════════════════════════════════════════════════
// SECTION 3: THE UNIT STRUCT (Where Concurrency Begins!)
// ═══════════════════════════════════════════════════════════════════════════
//
// ⚠️ CRITICAL MOMENT: This is where we start dealing with concurrency!
//
// 🎴 MTG ANALOGY:
//   Unit = Creature on the battlefield
//   - Can be tapped/untapped (state)
//   - Can have +1/+1 counters (stats)
//   - Can be targeted by spells (commands)
//   - Can be destroyed (lifecycle)
//
// ⚔️ SC:BW ANALOGY:
//   Unit = Marine, Zealot, Zergling
//   - Runs independently (goroutine!)
//   - Has health, damage, position (state)
//   - Receives orders (channel commands)
//   - Dies and cleans up (context cancellation)
//
// 🎓 KEY CONCURRENCY CONCEPTS IN THIS STRUCT:
//   1. sync.RWMutex - Protects shared state (health, position, etc.)
//   2. Channels - For receiving commands
//   3. Context - For lifecycle/cancellation
//   4. Goroutine - The unit "runs" in its own goroutine
//
// 🔥 PRO TIP FROM BISU:
// "In Protoss play, every unit has a role. Same in concurrent code—every
//  field in this struct has a purpose. Understand WHY each field exists."
//
// ═══════════════════════════════════════════════════════════════════════════

// Unit represents a single unit in the simulation
//
// 🤔 CRITICAL DESIGN QUESTIONS (Answer before implementing):
//
// Q1: Why use sync.RWMutex instead of sync.Mutex?
// Q2: Why are some fields protected by mutex and others aren't?
// Q3: Why use channels for commands instead of just calling methods?
// Q4: What happens if we forget to close channels or cancel context?
//
// 🎯 ANSWERS TO PONDER:
// A1: RWMutex allows multiple readers OR one writer. Read-heavy? RWMutex wins.
// A2: Immutable/constant fields (like ID, Type) don't need protection!
// A3: Channels = asynchronous, non-blocking. Method calls = synchronous, blocking.
// A4: GOROUTINE LEAKS! The unit goroutine runs forever, eating memory. BAD!
type Unit struct {
	// ═══════════════════════════════════════════════════════════════════════
	// IMMUTABLE FIELDS (Never change after creation → No mutex needed!)
	// ═══════════════════════════════════════════════════════════════════════
	ID   string   // Unique identifier, never changes
	Type UnitType // Marine stays Marine, can't morph (unless you're Zerg!)

	// ═══════════════════════════════════════════════════════════════════════
	// MUTABLE STATE (Protected by mutex—multiple goroutines access this!)
	// ═══════════════════════════════════════════════════════════════════════
	mu        sync.RWMutex // Protects ALL fields below
	health    int          // Current health (0 = dead)
	maxHealth int          // Maximum health
	damage    int          // Attack damage
	state     UnitState    // Current state (Idle, Moving, etc.)
	position  Position     // Current position
	target    *Unit        // Currently attacking this unit (nil if none)

	// ═══════════════════════════════════════════════════════════════════════
	// CONCURRENCY PRIMITIVES (Channels, Context, Coordination)
	// ═══════════════════════════════════════════════════════════════════════
	commands chan Command       // Receives commands (attack, move, etc.)
	events   chan UnitEvent     // Sends events (took damage, killed unit, etc.)
	ctx      context.Context    // For cancellation/shutdown
	cancel   context.CancelFunc // Call this to stop the unit's goroutine
	wg       *sync.WaitGroup    // For coordinated shutdown
}

// 🤔 PAUSE AND REFLECT:
// Look at that struct. See the three categories?
// 1. Immutable (no protection needed)
// 2. Mutable state (RWMutex protects it)
// 3. Concurrency primitives (channels, context)
//
// This is the PATTERN for all concurrent objects in Go.
// Master it here, use it everywhere.

// ═══════════════════════════════════════════════════════════════════════════
// CONSTRUCTOR: NewUnit
// ═══════════════════════════════════════════════════════════════════════════
//
// 🎴 MTG: Casting a creature spell
//   1. Pay mana cost
//   2. Spell goes on stack
//   3. Resolves → creature enters battlefield
//   4. Creature is now "alive" and can act
//
// ⚔️ SC:BW: Training a unit
//   1. Select building
//   2. Click train button
//   3. Wait for production
//   4. Unit pops out, ready for orders
//
// 🎓 IN OUR CODE:
//   1. Call NewUnit()
//   2. Initialize all fields
//   3. Start goroutine (go unit.run())
//   4. Unit is "alive" and processing
//
// ═══════════════════════════════════════════════════════════════════════════

// NewUnit creates a new unit and starts its lifecycle goroutine
//
// 🤔 BEFORE YOU CODE - CRITICAL QUESTIONS:
//
// Q1: What would happen if we forgot to start the goroutine (go u.run())?
// Q2: Why do we need to pass a WaitGroup pointer?
// Q3: What if we create a unit but never cancel its context?
// Q4: Should we use buffered or unbuffered channels for commands?
//
// 🎯 THINK ABOUT IT:
// A1: Unit would be "created" but never actually run. Like training a Marine
//     but it never leaves the barracks!
// A2: For graceful shutdown—we need to wait for ALL units to finish cleanly
// A3: GOROUTINE LEAK! The go u.run() runs forever, never stops. Memory leak!
// A4: Unbuffered = sender blocks until processed. Buffered = can queue commands.
//     What makes sense for game units? (Hint: buffer prevents missed commands)
//
// 🔥 PRO TIP FROM JAEDONG:
// "In Zerg, I queue up commands for my units (attack, move, attack). They
//  execute them in order. Buffered channel = command queue. Makes sense?"
//
// 🏗️ YOUR BIG CHALLENGE: Implement NewUnit
//
// You'll need to:
// 1. Create the Unit struct with initial values
// 2. Set up context with cancellation
// 3. Initialize channels (commands, events)
// 4. Set initial stats based on unit type
// 5. Start the goroutine
// 6. Return the unit
//
// 🎯 HINT LEVEL 1: Structure your code in the order listed above
// 🎯 HINT LEVEL 2: Use context.WithCancel(context.Background())
// 🎯 HINT LEVEL 3: Buffered channels → make(chan Command, 10)
//
// Template:
/*
func NewUnit(id string, unitType UnitType, pos Position, wg *sync.WaitGroup) *Unit {
	ctx, cancel := context.WithCancel(context.Background())

	u := &Unit{
		ID:   id,
		Type: unitType,
		// TODO: Set position
		// TODO: Set initial state (probably Idle?)
		// TODO: Create channels
		// TODO: Store ctx, cancel, wg
	}

	// TODO: Set stats based on unitType (use switch statement?)
	// Hint: initializeStats(u, unitType) helper function?

	// TODO: Start goroutine
	// Hint: wg.Add(1) BEFORE go u.run()

	return u
}
*/

// Helper function to set unit stats based on type
// 🤔 DESIGN QUESTION: Why extract this to a separate function?
// A: Single Responsibility Principle. NewUnit creates, this initializes stats.
//
// 🏗️ YOUR CHALLENGE: Implement this
/*
func initializeStats(u *Unit, unitType UnitType) {
	switch unitType {
	case Marine:
		u.maxHealth = 40
		u.health = 40
		u.damage = 6
	case Zergling:
		u.maxHealth = 35
		u.health = 35
		u.damage = 5
	// TODO: Add cases for other unit types
	// Refer to SC:BW stats or make up balanced numbers!
	default:
		// Fallback for unknown types
		u.maxHealth = 50
		u.health = 50
		u.damage = 5
	}
}
*/

// ═══════════════════════════════════════════════════════════════════════════
// CHECKPOINT: Before you continue...
// ═══════════════════════════════════════════════════════════════════════════
//
// ✅ Can you explain what RWMutex does and why we use it?
// ✅ Can you explain the difference between buffered and unbuffered channels?
// ✅ Can you explain why we need context cancellation?
// ✅ Can you explain what happens if we forget wg.Add(1)?
//
// If ANY of these are unclear, READ CONCEPTS_GUIDE.md before continuing!
// Understanding beats implementation. Always.
//
// ═══════════════════════════════════════════════════════════════════════════

// ═══════════════════════════════════════════════════════════════════════════
// SECTION 4: COMMAND & EVENT TYPES
// ═══════════════════════════════════════════════════════════════════════════
//
// 🎴 MTG ANALOGY:
//   Command = Instant/Sorcery targeting a creature ("Lightning Bolt target creature")
//   Event = Triggered ability ("Whenever a creature dies, draw a card")
//
// ⚔️ SC:BW ANALOGY:
//   Command = Player order (move here, attack that unit)
//   Event = Game event (unit died, unit killed another unit)
//
// 🎓 LEARNING: Commands flow DOWN (player → unit), Events flow UP (unit → game)
//
// ═══════════════════════════════════════════════════════════════════════════

// CommandType represents different types of commands a unit can receive
type CommandType int

const (
	CmdMove CommandType = iota
	CmdAttack
	CmdStop
	CmdHold
	// 🏗️ YOUR CHALLENGE: Add more command types
	// Ideas: CmdRetreat, CmdPatrol, CmdCast (for abilities)
)

// Command represents an order sent to a unit
//
// 🤔 DESIGN QUESTION: Why not just use methods like unit.Attack(target)?
// A: Commands through channels = asynchronous, non-blocking, queueable!
//
//	Methods = synchronous, blocking. Big difference in concurrent systems.
type Command struct {
	Type   CommandType
	Target *Unit    // For attack commands
	Dest   Position // For move commands
	// 🏗️ OPTIONAL: Add more fields for complex commands
}

// UnitEventType represents different types of events units can emit
type UnitEventType int

const (
	EventDamaged UnitEventType = iota
	EventKilled                // This unit killed another unit
	EventDied                  // This unit died
	EventMoved
	EventIdle
	// 🏗️ YOUR CHALLENGE: Add more event types as needed
)

// UnitEvent represents something that happened to/by a unit
//
// 🎴 MTG: Like triggered abilities ("When this creature deals damage...")
// ⚔️ SC:BW: Like game events (unit completed, unit died, etc.)
type UnitEvent struct {
	Type      UnitEventType
	Source    *Unit // Unit that generated this event
	Target    *Unit // For events involving another unit
	Timestamp time.Time
	Data      interface{} // For additional event-specific data
}

// ═══════════════════════════════════════════════════════════════════════════
// NEXT STEPS FOR YOU:
// ═══════════════════════════════════════════════════════════════════════════
//
// You've seen the foundation. Now implement:
//
// 1. **Getter Methods** (with RLock):
//    - GetHealth() int
//    - GetState() UnitState
//    - GetPosition() Position
//    - GetTarget() *Unit
//    Pattern: RLock → read value → RUnlock → return
//
// 2. **Setter Methods** (with Lock):
//    - SetState(UnitState)
//    - SetPosition(Position)
//    - TakeDamage(int)
//    Pattern: Lock → modify value → Unlock
//
// 3. **The run() Method** (The main goroutine loop):
//    - Use select to listen on multiple channels
//    - Handle commands from commands channel
//    - Check for context cancellation
//    - Send events to events channel
//    Pattern: for { select { case cmd := <-commands... case <-ctx.Done()... } }
//
// 4. **Command Handlers**:
//    - handleMoveCommand(Command)
//    - handleAttackCommand(Command)
//    - handleStopCommand(Command)
//
// 5. **Lifecycle Methods**:
//    - Shutdown() - Cancel context, wait for goroutine
//    - Die() - Set state to Dead, emit event, shutdown
//
// 🔥 PRO TIP FROM BISU:
// "Build one piece at a time. Get getters/setters working first. Test them.
//  Then build run(). Test it. Then add command handlers. One piece at a time."
//
// 📚 STUCK? Read CONCEPTS_GUIDE.md for pattern examples!
// 📊 NEED STRUCTURE? See PROGRESSION_MAP.md for the learning path!
//
// ═══════════════════════════════════════════════════════════════════════════

// Continue in types.go for your implementations...
// Remember: The goal isn't just to make it work—it's to UNDERSTAND why it works!
//
// **GG HF!** (Good Game, Have Fun learning!) 🎯
