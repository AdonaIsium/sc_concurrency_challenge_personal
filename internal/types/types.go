package types

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

// ═══════════════════════════════════════════════════════════════════════════
// UNIT STATS DATA - Embedded at Compile Time
// ═══════════════════════════════════════════════════════════════════════════
//
// 🎓 LEARNING: go:embed Directive
//
// This is a compile-time feature that embeds external files into your binary.
// The JSON file is read at BUILD time and baked into the executable.
//
// 🎴 MTG ANALOGY: Like having all card data printed on the back of each card.
// No need to look up the rulebook—the information is right there.
//
// ⚔️ SC:BW ANALOGY: Like having unit stats compiled into the game executable
// instead of loading from external files. Fast, reliable, no I/O overhead.
//
// 🔥 PRO TIP FROM FLASH:
// "In professional games, we compile game data into the binary for tournaments.
//  Can't have the game failing because a config file is missing!"
//
// ═══════════════════════════════════════════════════════════════════════════

//go:embed unit_stats.json
var unitStatsJSON []byte

// UnitStatsData holds the parsed unit stats for fast lookup
type UnitStatsData struct {
	MaxHealth      int    `json:"maxHealth"`
	BaseDamage     int    `json:"baseDamage"`
	BaseArmor      int    `json:"baseArmor"`
	ArmorModifier  int    `json:"armorModifier"`
	AttackModifier int    `json:"attackModifier"`
	AttackRange    int    `json:"attackRange"`
	VisionRange    int    `json:"visionRange"`
	ElevationLayer string `json:"elevationLayer"`
}

// unitStatsCache is populated once at startup from the embedded JSON
var unitStatsCache map[string]UnitStatsData

// init runs once when the package is loaded
// Parses the embedded JSON into our cache for fast runtime lookups
func init() {
	if err := json.Unmarshal(unitStatsJSON, &unitStatsCache); err != nil {
		log.Fatalf("Failed to load unit stats: %v", err)
	}
}

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
}

func (us UnitState) String() string {
	if name, ok := unitStateNames[us]; ok {
		return name
	}
	return fmt.Sprintf("UnitState(%d)", us)
}

func (us UnitState) IsValid() bool {
	return us >= Idle && us < unitStateCount
}

type ElevationLayer int

const (
	Burrowed ElevationLayer = iota
	Ground
	Air

	elevationLayerCount
)

var _ fmt.Stringer = ElevationLayer(0)

var elevationLayerNames = map[ElevationLayer]string{
	Burrowed: "Burrowed",
	Ground:   "Ground",
	Air:      "Air",
}

func (el ElevationLayer) String() string {
	if name, ok := elevationLayerNames[el]; ok {
		return name
	}
	return fmt.Sprintf("ElevationLayer(%d)", el)
}

func (el ElevationLayer) IsValid() bool {
	return el >= Burrowed && el < elevationLayerCount
}

type TerrainType int

const (
	LowGround TerrainType = iota
	MiddleGround
	HighGround
	LowMiddleRamp
	MiddleHighRamp
	Water

	terrainTypeCount
)

var _ fmt.Stringer = TerrainType(0)

var terrainTypeNames = map[TerrainType]string{
	LowGround:      "LowGround",
	MiddleGround:   "MiddleGround",
	HighGround:     "HighGround",
	LowMiddleRamp:  "LowMiddleRamp",
	MiddleHighRamp: "MiddleHighRamp",
	Water:          "Water",
}

func (tt TerrainType) String() string {
	if name, ok := terrainTypeNames[tt]; ok {
		return name
	}
	return fmt.Sprintf("TerrainType(%d)", tt)
}

func (tt TerrainType) IsValid() bool {
	return tt >= LowGround && tt < terrainTypeCount
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
}

// Distance calculates the Euclidean distance between two positions
func (p Position) Distance(other Position) float64 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// DistanceSq returns the squared distance (faster, no sqrt)
//
// 🔥 PRO TIP FROM FLASH:
// "In SC:BW, we compare distances constantly (attack range, vision, etc.).
//
//	Square root is expensive! If you only need to compare distances,
//	compare the SQUARED distances instead. Way faster."
func (p Position) DistanceSq(other Position) float64 {
	dx := other.X - p.X
	dy := other.Y - p.Y
	return dx*dx + dy*dy
}

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
	mu             sync.RWMutex // Protects ALL fields below
	health         int          // Current health (0 = dead)
	maxHealth      int          // Maximum health
	baseDamage     int          // Attack damage
	baseArmor      int
	attackModifier int
	armorModifier  int
	attackUpgrades int
	armorUpgrades  int
	attackRange    int
	visionRange    int
	state          UnitState      // Current state (Idle, Moving, etc.)
	elevationLayer ElevationLayer // Current Elevation (Burrowed, Flying, Ground)
	position       Position       // Current position
	target         *Unit          // Currently attacking this unit (nil if none)

	// ═══════════════════════════════════════════════════════════════════════
	// CONCURRENCY PRIMITIVES (Channels, Context, Coordination)
	// ═══════════════════════════════════════════════════════════════════════
	commands     chan Command       // Receives commands (attack, move, etc.)
	events       chan UnitEvent     // Sends events (took damage, killed unit, etc.)
	ctx          context.Context    // For cancellation/shutdown
	cancel       context.CancelFunc // Call this to stop the unit's goroutine
	wg           *sync.WaitGroup    // For coordinated shutdown
	shutdownOnce sync.Once          // Ensures Shutdown only executes once
}

type Tile struct {
	Position Position
	Terrain  TerrainType
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
// 🔥 PRO TIP FROM JAEDONG:
// "In Zerg, I queue up commands for my units (attack, move, attack). They
//
//	execute them in order. Buffered channel = command queue. Makes sense?"
func NewUnit(id string, unitType UnitType, pos Position, wg *sync.WaitGroup) *Unit {
	ctx, cancel := context.WithCancel(context.Background())

	u := &Unit{
		ID:       id,
		Type:     unitType,
		position: pos,
		state:    Idle,

		commands: make(chan Command, 10),
		events:   make(chan UnitEvent, 10),

		ctx:    ctx,
		cancel: cancel,
		wg:     wg,
	}

	initializeStats(u, unitType)

	wg.Add(1)
	go u.run()

	return u
}

// initializeStats loads unit stats from the embedded JSON cache
//
// 🎓 LEARNING: Data-Driven Design Pattern
//
// Instead of hardcoding 250+ lines of switch cases, we:
// 1. Store data in external JSON file (maintainable!)
// 2. Embed at compile time (zero runtime I/O overhead!)
// 3. Parse once at startup in init() (one-time cost)
// 4. Lookup from map in O(1) time (blazing fast!)
//
// 🎴 MTG ANALOGY: Like having all card stats in a database instead of
// hardcoding each card's P/T, mana cost, and abilities in the game code.
// MTG Arena does exactly this—thousands of cards, one clean data file.
//
// ⚔️ SC:BW ANALOGY: StarCraft stores unit stats in data files, not code.
// When Blizzard wants to balance a unit (nerf Mutalisk damage), they
// edit a data file, not recompile the entire game engine!
//
// 🔥 PRO TIP FROM FLASH:
// "In professional play, we analyze unit stats constantly. Having them
//
//	in a readable data format makes balance analysis much easier than
//	digging through game code!"
//
// 📊 PERFORMANCE:
// - Startup: ~2-5ms one-time JSON parse
// - Runtime: O(1) map lookup, same speed as switch statement
// - Memory: ~5KB for all 34 unit types
// - I/O: ZERO (file embedded at compile time)
func initializeStats(u *Unit, unitType UnitType) {
	// Lookup stats from our cache (populated by init() at startup)
	stats, ok := unitStatsCache[unitType.String()]
	if !ok {
		// Fallback for unknown unit types (defensive programming)
		log.Printf("Warning: Unknown unit type %s, using default stats", unitType)
		stats = UnitStatsData{
			MaxHealth:      50,
			BaseDamage:     5,
			BaseArmor:      0,
			ArmorModifier:  1,
			AttackModifier: 0,
			AttackRange:    4,
			VisionRange:    7,
			ElevationLayer: "Ground",
		}
	}

	// Apply stats from cache
	u.maxHealth, u.health = stats.MaxHealth, stats.MaxHealth
	u.baseDamage = stats.BaseDamage
	u.baseArmor = stats.BaseArmor
	u.armorModifier = stats.ArmorModifier
	u.attackModifier = stats.AttackModifier
	u.armorUpgrades = 0
	u.attackUpgrades = 0
	u.attackRange = stats.AttackRange
	u.visionRange = stats.VisionRange

	// Parse elevation layer string to enum
	switch stats.ElevationLayer {
	case "Ground":
		u.elevationLayer = Ground
	case "Air":
		u.elevationLayer = Air
	case "Burrowed":
		u.elevationLayer = Burrowed
	default:
		u.elevationLayer = Ground
	}
}

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
}

func (c Command) String() string {
	switch c.Type {
	case CmdMove:
		return fmt.Sprintf("Move to (%.1f, %.1f)", c.Dest.X, c.Dest.Y)
	case CmdAttack:
		targetID := "nil"
		if c.Target != nil {
			targetID = c.Target.ID
		}
		return fmt.Sprintf("Attack %s", targetID)
	case CmdStop:
		return "Stop"
	case CmdHold:
		return "Hold Position"
	default:
		return fmt.Sprintf("Command(%d)", c.Type)

	}
}

// UnitEventType represents different types of events units can emit
type UnitEventType int

const (
	EventDamaged UnitEventType = iota
	EventKilled                // This unit killed another unit
	EventDied                  // This unit died
	EventMoved
	EventIdle
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
// UNIT METHODS
// ═══════════════════════════════════════════════════════════════════════════

func (u *Unit) GetHealth() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.health
}

func (u *Unit) GetDamage() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.baseDamage + u.attackModifier*u.attackUpgrades
}

func (u *Unit) GetArmor() int {
	u.mu.RLock()
	defer u.mu.RUnlock()
	totalArmor := u.baseArmor + u.armorModifier*u.armorUpgrades
	return totalArmor
}

func (u *Unit) GetState() UnitState {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.state
}

func (u *Unit) GetPosition() Position {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.position
}

func (u *Unit) GetTarget() *Unit {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.target
}

func (u *Unit) SetState(state UnitState) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.state = state
}

func (u *Unit) SetPosition(position Position) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.position = position
}

func (u *Unit) SetTarget(target *Unit) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.target = target
}

func (u *Unit) TakeDamage(amount int) int {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.health -= amount
	if u.health < 0 {
		u.health = 0
	}

	return u.health
}

func (u *Unit) run() {
	defer u.wg.Done()

	for {
		select {
		case cmd, ok := <-u.commands:
			if !ok {
				return
			}
			switch cmd.Type {
			case CmdMove:
				u.handleMove(cmd)
			case CmdAttack:
				u.handleAttack(cmd)
			case CmdStop:
				u.handleStop(cmd)
			case CmdHold:
				u.handleHold(cmd)
			}
			fmt.Printf("[%s] %s\n", u.ID, cmd)
		case <-u.ctx.Done():
			return
		}
	}
}

func (u *Unit) SendCommand(cmd Command) error {
	// Check context first to avoid sending on closed channel
	select {
	case <-u.ctx.Done():
		return fmt.Errorf("unit is shutting down")
	default:
	}

	// Now try to send with timeout
	select {
	case u.commands <- cmd:
		return nil
	case <-u.ctx.Done():
		return fmt.Errorf("unit is shutting down")
	case <-time.After(100 * time.Millisecond):
		return fmt.Errorf("command queue full - backpressure")
	}
}

func (u *Unit) Shutdown() {
	// Use sync.Once to ensure this only runs once, even if called multiple times
	u.shutdownOnce.Do(func() {
		u.cancel()
		close(u.commands)
	})
}

func (u *Unit) handleMove(cmd Command) {
	u.SetState(Moving)
	u.SetPosition(cmd.Dest)
	moveEvent := UnitEvent{Type: EventMoved, Source: u, Target: nil, Timestamp: time.Now()}
	u.events <- moveEvent
}

func (u *Unit) handleAttack(cmd Command) {
	if cmd.Target == nil {
		fmt.Printf("[%s] Attack command has no target!\n", u.ID)
		return
	}
	u.SetState(Attacking)
	u.SetTarget(cmd.Target)
	dmg := u.CalculateDamageAgainst(cmd.Target)
	cmd.Target.TakeDamage(dmg)
	attackEvent := UnitEvent{Type: EventDamaged, Source: u, Target: cmd.Target, Timestamp: time.Now()}
	u.events <- attackEvent
}

func (u *Unit) handleStop(cmd Command) {
	u.SetState(Idle)
	u.SetTarget(nil)
}

func (u *Unit) handleHold(cmd Command) {
	u.SetState(HoldingPosition)
}

func (u *Unit) CalculateDamageAgainst(target *Unit) int {
	targetArmor := target.GetArmor()
	dmg := u.GetDamage()
	result := dmg - targetArmor
	if result < 0 {
		return 0
	}
	return result
}
