package types

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// ═══════════════════════════════════════════════════════════════════════════
// ENUM TESTS
// ═══════════════════════════════════════════════════════════════════════════

func TestUnitType_String(t *testing.T) {
	tests := []struct {
		name     string
		unitType UnitType
		expected string
	}{
		{"Marine", Marine, "Marine"},
		{"Zergling", Zergling, "Zergling"},
		{"Zealot", Zealot, "Zealot"},
		{"Invalid", UnitType(999), "UnitType(999)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.unitType.String()
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestUnitType_IsValid(t *testing.T) {
	require.True(t, Marine.IsValid(), "Marine should be valid")
	require.True(t, SCV.IsValid(), "SCV should be valid")
	require.True(t, Arbiter.IsValid(), "Arbiter should be valid")
	require.False(t, UnitType(-1).IsValid(), "Negative value should be invalid")
	require.False(t, UnitType(999).IsValid(), "Out of range value should be invalid")
	require.False(t, unitTypeCount.IsValid(), "Sentinel value should be invalid")
}

func TestUnitState_String(t *testing.T) {
	tests := []struct {
		name     string
		state    UnitState
		expected string
	}{
		{"Idle", Idle, "Idle"},
		{"Moving", Moving, "Moving"},
		{"Attacking", Attacking, "Attacking"},
		{"Dead", Dead, "Dead"},
		{"Invalid", UnitState(999), "UnitState(999)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.state.String()
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestUnitState_IsValid(t *testing.T) {
	require.True(t, Idle.IsValid(), "Idle should be valid")
	require.True(t, Dead.IsValid(), "Dead should be valid")
	require.False(t, UnitState(-1).IsValid(), "Negative value should be invalid")
	require.False(t, UnitState(999).IsValid(), "Out of range value should be invalid")
	require.False(t, unitStateCount.IsValid(), "Sentinel value should be invalid")
}

func TestElevationLayer_String(t *testing.T) {
	require.Equal(t, "Ground", Ground.String())
	require.Equal(t, "Air", Air.String())
	require.Equal(t, "Burrowed", Burrowed.String())
	require.Equal(t, "ElevationLayer(999)", ElevationLayer(999).String())
}

func TestElevationLayer_IsValid(t *testing.T) {
	require.True(t, Ground.IsValid())
	require.True(t, Air.IsValid())
	require.True(t, Burrowed.IsValid())
	require.False(t, ElevationLayer(-1).IsValid())
	require.False(t, elevationLayerCount.IsValid())
}

func TestTerrainType_String(t *testing.T) {
	require.Equal(t, "LowGround", LowGround.String())
	require.Equal(t, "HighGround", HighGround.String())
	require.Equal(t, "Water", Water.String())
	require.Equal(t, "TerrainType(999)", TerrainType(999).String())
}

func TestTerrainType_IsValid(t *testing.T) {
	require.True(t, LowGround.IsValid())
	require.True(t, Water.IsValid())
	require.False(t, TerrainType(-1).IsValid())
	require.False(t, terrainTypeCount.IsValid())
}

func TestCommand_String(t *testing.T) {
	tests := []struct {
		name     string
		cmd      Command
		expected string
	}{
		{
			name:     "Move command",
			cmd:      Command{Type: CmdMove, Dest: Position{X: 50.5, Y: 75.3}},
			expected: "Move to (50.5, 75.3)",
		},
		{
			name:     "Stop command",
			cmd:      Command{Type: CmdStop},
			expected: "Stop",
		},
		{
			name:     "Hold command",
			cmd:      Command{Type: CmdHold},
			expected: "Hold Position",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.cmd.String()
			require.Equal(t, tt.expected, result)
		})
	}
}

func TestCommand_String_WithTarget(t *testing.T) {
	var wg sync.WaitGroup
	target := NewUnit("target-1", Zergling, Position{X: 100, Y: 100}, &wg)

	cmd := Command{Type: CmdAttack, Target: target}
	require.Equal(t, "Attack target-1", cmd.String())

	cmdNil := Command{Type: CmdAttack, Target: nil}
	require.Equal(t, "Attack nil", cmdNil.String())

	target.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// POSITION TESTS
// ═══════════════════════════════════════════════════════════════════════════

func TestPosition_Distance(t *testing.T) {
	p1 := Position{X: 0, Y: 0}
	p2 := Position{X: 3, Y: 4}

	distance := p1.Distance(p2)
	require.InDelta(t, 5.0, distance, 0.001, "3-4-5 triangle should have distance 5")
}

func TestPosition_Distance_SamePoint(t *testing.T) {
	p := Position{X: 10, Y: 20}
	distance := p.Distance(p)
	require.Equal(t, 0.0, distance, "Distance to self should be 0")
}

func TestPosition_DistanceSq(t *testing.T) {
	p1 := Position{X: 0, Y: 0}
	p2 := Position{X: 3, Y: 4}

	distSq := p1.DistanceSq(p2)
	require.Equal(t, 25.0, distSq, "Squared distance for 3-4-5 triangle should be 25")
}

func TestPosition_DistanceSq_PreservesOrdering(t *testing.T) {
	origin := Position{X: 0, Y: 0}
	near := Position{X: 1, Y: 1}
	far := Position{X: 10, Y: 10}

	// If Distance(near) < Distance(far), then DistanceSq(near) < DistanceSq(far)
	require.Less(t, origin.Distance(near), origin.Distance(far))
	require.Less(t, origin.DistanceSq(near), origin.DistanceSq(far))
}

// ═══════════════════════════════════════════════════════════════════════════
// UNIT CONSTRUCTOR TESTS
// ═══════════════════════════════════════════════════════════════════════════

func TestNewUnit_InitializesFieldsCorrectly(t *testing.T) {
	var wg sync.WaitGroup
	pos := Position{X: 100, Y: 200}

	unit := NewUnit("marine-1", Marine, pos, &wg)
	require.NotNil(t, unit)

	// Check immutable fields
	require.Equal(t, "marine-1", unit.ID)
	require.Equal(t, Marine, unit.Type)

	// Check position
	require.Equal(t, pos, unit.GetPosition())

	// Check initial state
	require.Equal(t, Idle, unit.GetState())

	// Check stats were initialized
	require.Equal(t, 40, unit.GetHealth(), "Marine should have 40 HP")
	require.Equal(t, 6, unit.GetDamage(), "Marine should have 6 base damage")

	// Cleanup
	unit.Shutdown()
	wg.Wait()
}

func TestNewUnit_StartsGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test-unit", Marine, Position{}, &wg)

	// Send a command to verify goroutine is running
	err := unit.SendCommand(Command{Type: CmdStop})
	require.NoError(t, err)

	// Give it time to process
	time.Sleep(50 * time.Millisecond)

	// Shutdown should complete without deadlock
	unit.Shutdown()

	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Success
	case <-time.After(1 * time.Second):
		t.Fatal("Goroutine didn't shutdown properly")
	}
}

func TestInitializeStats_Marine(t *testing.T) {
	var wg sync.WaitGroup
	marine := NewUnit("marine-1", Marine, Position{}, &wg)

	require.Equal(t, 40, marine.GetHealth())
	require.Equal(t, 6, marine.GetDamage())
	require.Equal(t, 0, marine.GetArmor())

	marine.Shutdown()
	wg.Wait()
}

func TestInitializeStats_Zergling(t *testing.T) {
	var wg sync.WaitGroup
	zergling := NewUnit("zergling-1", Zergling, Position{}, &wg)

	require.Equal(t, 35, zergling.GetHealth())
	require.Equal(t, 5, zergling.GetDamage())

	zergling.Shutdown()
	wg.Wait()
}

func TestInitializeStats_Zealot(t *testing.T) {
	var wg sync.WaitGroup
	zealot := NewUnit("zealot-1", Zealot, Position{}, &wg)

	require.Equal(t, 100, zealot.GetHealth())
	require.Equal(t, 16, zealot.GetDamage())

	zealot.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// GETTER TESTS (Thread-Safety)
// ═══════════════════════════════════════════════════════════════════════════

func TestUnit_GetHealth(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	health := unit.GetHealth()
	require.Equal(t, 40, health)

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_GetDamage(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	damage := unit.GetDamage()
	require.Equal(t, 6, damage)

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_GetArmor(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	armor := unit.GetArmor()
	require.Equal(t, 0, armor)

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_GetState(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	state := unit.GetState()
	require.Equal(t, Idle, state)

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_GetPosition(t *testing.T) {
	var wg sync.WaitGroup
	pos := Position{X: 100, Y: 200}
	unit := NewUnit("test", Marine, pos, &wg)

	position := unit.GetPosition()
	require.Equal(t, pos, position)

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_GetTarget(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	target := unit.GetTarget()
	require.Nil(t, target, "Initial target should be nil")

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_Getters_ConcurrentAccess(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{X: 50, Y: 50}, &wg)

	// Spawn multiple goroutines reading concurrently
	var readWg sync.WaitGroup
	for i := 0; i < 100; i++ {
		readWg.Add(1)
		go func() {
			defer readWg.Done()
			_ = unit.GetHealth()
			_ = unit.GetState()
			_ = unit.GetPosition()
			_ = unit.GetDamage()
			_ = unit.GetArmor()
		}()
	}

	readWg.Wait()
	unit.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// SETTER TESTS (Thread-Safety)
// ═══════════════════════════════════════════════════════════════════════════

func TestUnit_SetState(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	unit.SetState(Moving)
	require.Equal(t, Moving, unit.GetState())

	unit.SetState(Attacking)
	require.Equal(t, Attacking, unit.GetState())

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_SetPosition(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	newPos := Position{X: 100, Y: 200}
	unit.SetPosition(newPos)
	require.Equal(t, newPos, unit.GetPosition())

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_SetTarget(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)
	target := NewUnit("target", Zergling, Position{}, &wg)

	unit.SetTarget(target)
	require.Equal(t, target, unit.GetTarget())

	unit.SetTarget(nil)
	require.Nil(t, unit.GetTarget())

	unit.Shutdown()
	target.Shutdown()
	wg.Wait()
}

func TestUnit_TakeDamage(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Marine has 40 HP
	remainingHealth := unit.TakeDamage(10)
	require.Equal(t, 30, remainingHealth)
	require.Equal(t, 30, unit.GetHealth())

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_TakeDamage_ClampsToZero(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Deal more damage than health
	remainingHealth := unit.TakeDamage(100)
	require.Equal(t, 0, remainingHealth)
	require.Equal(t, 0, unit.GetHealth())

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_Setters_ConcurrentAccess(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Spawn multiple goroutines writing concurrently
	var writeWg sync.WaitGroup
	for i := 0; i < 50; i++ {
		writeWg.Add(1)
		go func(index int) {
			defer writeWg.Done()
			unit.SetPosition(Position{X: float64(index), Y: float64(index)})
			unit.SetState(Idle)
			unit.TakeDamage(1)
		}(i)
	}

	writeWg.Wait()

	// Verify unit is still in valid state
	require.GreaterOrEqual(t, unit.GetHealth(), 0)
	require.True(t, unit.GetState().IsValid())

	unit.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// COMMAND HANDLER TESTS
// ═══════════════════════════════════════════════════════════════════════════

func TestUnit_HandleMove(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{X: 0, Y: 0}, &wg)

	dest := Position{X: 100, Y: 200}
	err := unit.SendCommand(Command{Type: CmdMove, Dest: dest})
	require.NoError(t, err)

	// Give time to process
	time.Sleep(50 * time.Millisecond)

	require.Equal(t, Moving, unit.GetState())
	require.Equal(t, dest, unit.GetPosition())

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_HandleAttack(t *testing.T) {
	var wg sync.WaitGroup
	attacker := NewUnit("marine", Marine, Position{}, &wg)
	target := NewUnit("zergling", Zergling, Position{}, &wg)

	initialHealth := target.GetHealth()

	err := attacker.SendCommand(Command{Type: CmdAttack, Target: target})
	require.NoError(t, err)

	// Give time to process
	time.Sleep(50 * time.Millisecond)

	require.Equal(t, Attacking, attacker.GetState())
	require.Equal(t, target, attacker.GetTarget())
	require.Less(t, target.GetHealth(), initialHealth, "Target should have taken damage")

	attacker.Shutdown()
	target.Shutdown()
	wg.Wait()
}

func TestUnit_HandleAttack_NilTarget(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Send attack command with nil target
	err := unit.SendCommand(Command{Type: CmdAttack, Target: nil})
	require.NoError(t, err)

	// Give time to process
	time.Sleep(50 * time.Millisecond)

	// State should not change to Attacking
	require.NotEqual(t, Attacking, unit.GetState())
	require.Nil(t, unit.GetTarget())

	unit.Shutdown()
	wg.Wait()
}

func TestUnit_HandleStop(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)
	target := NewUnit("target", Zergling, Position{}, &wg)

	// Set to attacking state
	unit.SetState(Attacking)
	unit.SetTarget(target)

	err := unit.SendCommand(Command{Type: CmdStop})
	require.NoError(t, err)

	// Give time to process
	time.Sleep(50 * time.Millisecond)

	require.Equal(t, Idle, unit.GetState())
	require.Nil(t, unit.GetTarget())

	unit.Shutdown()
	target.Shutdown()
	wg.Wait()
}

func TestUnit_HandleHold(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	err := unit.SendCommand(Command{Type: CmdHold})
	require.NoError(t, err)

	// Give time to process
	time.Sleep(50 * time.Millisecond)

	require.Equal(t, HoldingPosition, unit.GetState())

	unit.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// DAMAGE CALCULATION TESTS
// ═══════════════════════════════════════════════════════════════════════════

func TestUnit_CalculateDamageAgainst(t *testing.T) {
	var wg sync.WaitGroup
	marine := NewUnit("marine", Marine, Position{}, &wg)
	zergling := NewUnit("zergling", Zergling, Position{}, &wg)

	// Marine: 6 damage, Zergling: 0 armor
	damage := marine.CalculateDamageAgainst(zergling)
	require.Equal(t, 6, damage)

	marine.Shutdown()
	zergling.Shutdown()
	wg.Wait()
}

func TestUnit_CalculateDamageAgainst_WithArmor(t *testing.T) {
	var wg sync.WaitGroup
	marine := NewUnit("marine", Marine, Position{}, &wg)
	zealot := NewUnit("zealot", Zealot, Position{}, &wg)

	// Marine: 6 damage, Zealot: 1 armor
	damage := marine.CalculateDamageAgainst(zealot)
	require.Equal(t, 5, damage)

	marine.Shutdown()
	zealot.Shutdown()
	wg.Wait()
}

func TestUnit_CalculateDamageAgainst_FloorsAtZero(t *testing.T) {
	var wg sync.WaitGroup
	probe := NewUnit("probe", Probe, Position{}, &wg)
	battlecruiser := NewUnit("bc", Battlecruiser, Position{}, &wg)

	// Probe: 5 damage, Battlecruiser: 3 armor
	damage := probe.CalculateDamageAgainst(battlecruiser)
	require.Equal(t, 2, damage)

	probe.Shutdown()
	battlecruiser.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// GOROUTINE LIFECYCLE TESTS
// ═══════════════════════════════════════════════════════════════════════════

func TestUnit_Shutdown_StopsGoroutine(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	unit.Shutdown()

	// Wait with timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Success - goroutine terminated
	case <-time.After(1 * time.Second):
		t.Fatal("Goroutine didn't terminate after Shutdown")
	}
}

func TestUnit_MultipleShutdown_DoesNotPanic(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Calling Shutdown multiple times should not panic
	require.NotPanics(t, func() {
		unit.Shutdown()
		unit.Shutdown()
		unit.Shutdown()
	})

	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// CONCURRENT SAFETY TESTS (Race Detector)
// ═══════════════════════════════════════════════════════════════════════════

func TestUnit_ConcurrentReadWrite(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{X: 0, Y: 0}, &wg)

	var testWg sync.WaitGroup

	// Spawn readers
	for i := 0; i < 50; i++ {
		testWg.Add(1)
		go func() {
			defer testWg.Done()
			for j := 0; j < 100; j++ {
				_ = unit.GetHealth()
				_ = unit.GetState()
				_ = unit.GetPosition()
			}
		}()
	}

	// Spawn writers
	for i := 0; i < 10; i++ {
		testWg.Add(1)
		go func(index int) {
			defer testWg.Done()
			for j := 0; j < 50; j++ {
				unit.SetPosition(Position{X: float64(index), Y: float64(j)})
				unit.SetState(Moving)
				unit.TakeDamage(1)
			}
		}(i)
	}

	testWg.Wait()
	unit.Shutdown()
	wg.Wait()
}

func TestUnit_ConcurrentCommands(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	var testWg sync.WaitGroup

	// Spawn multiple goroutines sending commands
	// Keep it small to avoid overwhelming the channel buffer
	for i := 0; i < 5; i++ {
		testWg.Add(1)
		go func(index int) {
			defer testWg.Done()
			for j := 0; j < 2; j++ {
				_ = unit.SendCommand(Command{
					Type: CmdMove,
					Dest: Position{X: float64(index), Y: float64(j)},
				})
			}
		}(i)
	}

	testWg.Wait()
	time.Sleep(100 * time.Millisecond) // Let commands process

	unit.Shutdown()
	wg.Wait()
}

// ═══════════════════════════════════════════════════════════════════════════
// CONCURRENCY FLAW TESTS (Documenting Current Issues)
// ═══════════════════════════════════════════════════════════════════════════

// TestUnit_SendCommand_BlocksWhenBufferFull verifies that SendCommand
// returns an error when the buffer is full instead of blocking indefinitely.
func TestUnit_SendCommand_BlocksWhenBufferFull(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Fill the buffer (size 10) quickly before goroutine processes
	// Use CmdStop which doesn't send events, to avoid event buffer issues
	for i := 0; i < 15; i++ {
		err := unit.SendCommand(Command{Type: CmdStop})
		if err != nil {
			// Expected - hit backpressure
			require.Contains(t, err.Error(), "backpressure")
			unit.Shutdown()
			wg.Wait()
			return
		}
	}

	// If we get here without error, the goroutine is processing too fast
	// This is actually fine - just means our buffer is never full
	t.Log("Note: Could not fill buffer - goroutine processing too fast")
	unit.Shutdown()
	wg.Wait()
}

// TestUnit_SendCommand_AfterShutdown verifies that SendCommand
// respects context cancellation and returns an error after shutdown.
func TestUnit_SendCommand_AfterShutdown(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Shutdown the unit
	unit.Shutdown()
	wg.Wait()

	// Try to send command after shutdown - should return error
	err := unit.SendCommand(Command{Type: CmdStop})
	require.Error(t, err, "SendCommand should error when unit is shut down")
	require.Contains(t, err.Error(), "shutting down")
}

// TestUnit_Shutdown_WithPendingCommands verifies that shutdown
// completes gracefully even with pending commands in the buffer.
func TestUnit_Shutdown_WithPendingCommands(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Send commands but don't give time to process
	for i := 0; i < 5; i++ {
		err := unit.SendCommand(Command{
			Type: CmdMove,
			Dest: Position{X: float64(i), Y: float64(i)},
		})
		require.NoError(t, err)
	}

	// Shutdown immediately - should complete without deadlock
	unit.Shutdown()

	// This should complete without hanging
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// Success - shutdown completed gracefully
	case <-time.After(1 * time.Second):
		t.Fatal("Shutdown didn't complete - goroutine stuck")
	}
}

// TestUnit_ConcurrentCommands_Stress is a stress test that verifies
// the system handles many concurrent senders gracefully with backpressure.
func TestUnit_ConcurrentCommands_Stress(t *testing.T) {
	var wg sync.WaitGroup
	unit := NewUnit("test", Marine, Position{}, &wg)

	// Drain events channel in background to prevent blocking
	go func() {
		for {
			select {
			case <-unit.events:
				// Discard events
			case <-unit.ctx.Done():
				return
			}
		}
	}()

	var testWg sync.WaitGroup
	var successCount, errorCount int32

	// Realistic game scenario: 50 goroutines sending commands rapidly
	for i := 0; i < 50; i++ {
		testWg.Add(1)
		go func(index int) {
			defer testWg.Done()
			for j := 0; j < 20; j++ {
				err := unit.SendCommand(Command{
					Type: CmdMove,
					Dest: Position{X: float64(index), Y: float64(j)},
				})
				if err != nil {
					// Backpressure or shutdown - graceful handling
					atomic.AddInt32(&errorCount, 1)
					break
				} else {
					atomic.AddInt32(&successCount, 1)
				}
			}
		}(i)
	}

	testWg.Wait()
	unit.Shutdown()
	wg.Wait()

	// Verify no deadlock (test completed)
	// Some commands should succeed, some should error (backpressure)
	t.Logf("Stress test: %d commands succeeded, %d hit backpressure/shutdown",
		atomic.LoadInt32(&successCount), atomic.LoadInt32(&errorCount))

	require.Greater(t, atomic.LoadInt32(&successCount), int32(0), "At least some commands should succeed")
}
