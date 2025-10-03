package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/AdonaIsium/sc_concurrency_challenge_personal/internal/battle"
	"github.com/AdonaIsium/sc_concurrency_challenge_personal/internal/types"
)

// LEARNING NOTE: This example demonstrates:
// - Basic unit creation and management
// - Simple battle setup and execution
// - Event monitoring and reporting
// - Resource cleanup and shutdown
// - Error handling in concurrent systems

// BasicCombatExample demonstrates a simple battle between two unit types
// LEARNING: This is a complete, runnable example showing key concurrency concepts
func main() {
	fmt.Println("=== Basic Combat Example ===")
	fmt.Println("This example demonstrates basic unit combat with concurrency patterns")
	fmt.Println()

	// TODO: Implement the complete basic combat example
	// This should be a working example that students can run and learn from

	// Example structure (you need to implement each TODO):
	if err := runBasicCombatExample(); err != nil {
		log.Fatalf("Example failed: %v", err)
	}
}

// runBasicCombatExample executes the basic combat scenario
func runBasicCombatExample() error {
	// TODO: Implement basic combat example
	// Steps to implement:

	// 1. Set up context and cancellation
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	// defer cancel()
	// var wg sync.WaitGroup

	// 2. Create unit manager
	// unitManager := units.NewUnitManager(ctx, 4) // 4 command workers
	// defer unitManager.Shutdown(5 * time.Second)

	// 3. Create battle simulator
	// battleSim := battle.NewBattleSimulator(ctx, 100*time.Millisecond)
	// defer battleSim.Shutdown(5 * time.Second)

	// 4. Create Marine units (attacking force)
	// marines := createMarineSquad(5, &wg)
	// for _, marine := range marines {
	//     unitManager.AddUnit(marine)
	// }

	// 5. Create Zealot units (defending force)
	// zealots := createZealotSquad(3, &wg)
	// for _, zealot := range zealots {
	//     unitManager.AddUnit(zealot)
	// }

	// 6. Set up battle
	// battleConfig := createBattleConfig(marines, zealots)
	// battle, err := battleSim.CreateBattle(battleConfig)
	// if err != nil {
	//     return fmt.Errorf("failed to create battle: %w", err)
	// }

	// 7. Monitor battle progress
	// return monitorBattleProgress(ctx, battleSim, battle.GetID())

	fmt.Println("TODO: Implement runBasicCombatExample function")
	fmt.Println("See comments above for implementation steps")
	return nil
}

// createMarineSquad creates a squad of Marine units
// LEARNING: Unit factory pattern with proper initialization
func createMarineSquad(count int, wg *sync.WaitGroup) []*types.Unit {
	// TODO: Implement Marine squad creation
	// Steps needed:
	// 1. Create slice to hold units
	// marines := make([]*types.Unit, 0, count)

	// 2. Create each Marine unit
	// for i := 0; i < count; i++ {
	//     // Generate unique ID
	//     id := fmt.Sprintf("marine-%d", i+1)
	//
	//     // Set position (spread them out in a line)
	//     position := types.Position{X: float64(i * 2), Y: 0}
	//
	//     // Create the unit
	//     marine := types.NewUnit(id, types.Marine, position, wg)
	//     marines = append(marines, marine)
	// }

	// 3. Return the squad
	// return marines

	fmt.Println("TODO: Implement createMarineSquad")
	fmt.Println("This should create", count, "Marine units with proper positioning")
	return nil
}

// createZealotSquad creates a squad of Zealot units
// LEARNING: Different unit types with different characteristics
func createZealotSquad(count int, wg *sync.WaitGroup) []*types.Unit {
	// TODO: Implement Zealot squad creation
	// Similar to Marines but:
	// 1. Use types.Zealot as the unit type
	// 2. Position them on the opposite side (e.g., Y: 20)
	// 3. Zealots are typically stronger but fewer in number

	fmt.Println("TODO: Implement createZealotSquad")
	fmt.Println("This should create", count, "Zealot units positioned opposite Marines")
	return nil
}

// createBattleConfig sets up the battle configuration
// LEARNING: Battle configuration and setup patterns
func createBattleConfig(attackers, defenders []*types.Unit) battle.BattleConfig {
	// TODO: Implement battle configuration
	// Create a BattleConfig struct with:
	// 1. Unique battle ID
	// 2. Attacking and defending units
	// 3. Battlefield area (rectangle)
	// 4. Battle objectives (eliminate all enemies)
	// 5. Time limit
	// 6. Environment settings

	// Example structure:
	// return battle.BattleConfig{
	//     ID:        "basic-combat-" + time.Now().Format("20060102150405"),
	//     Attackers: attackers,
	//     Defenders: defenders,
	//     Battlefield: types.Rectangle{
	//         TopLeft:     types.Position{X: -10, Y: -5},
	//         BottomRight: types.Position{X: 30, Y: 25},
	//     },
	//     Objectives: []battle.BattleObjective{
	//         {
	//             Type:        battle.EliminateAll,
	//             Description: "Eliminate all enemy units",
	//             Points:      100,
	//         },
	//     },
	//     TimeLimit: 5 * time.Minute,
	//     Environment: battle.EnvironmentConfig{
	//         Weather:    battle.Clear,
	//         Visibility: 1.0,
	//         TerrainType: battle.Open,
	//     },
	//     Rules: battle.BattleRules{
	//         FriendlyFire:     false,
	//         Reinforcements:   false,
	//         SpecialAbilities: false,
	//         TimeScale:        1.0,
	//     },
	// }

	fmt.Println("TODO: Implement createBattleConfig")
	return battle.BattleConfig{}
}

// monitorBattleProgress watches the battle and reports on progress
// LEARNING: Event monitoring and real-time reporting
func monitorBattleProgress(ctx context.Context, simulator *battle.BattleSimulator, battleID string) error {
	// TODO: Implement battle monitoring
	// Steps needed:
	// 1. Subscribe to battle events
	// observer := simulator.AddObserver()

	// 2. Start battle
	// simulator.Start()

	// 3. Monitor events and status
	// ticker := time.NewTicker(1 * time.Second)
	// defer ticker.Stop()
	//
	// fmt.Println("Battle started! Monitoring progress...")
	//
	// for {
	//     select {
	//     case <-ctx.Done():
	//         return ctx.Err()
	//
	//     case event := <-observer:
	//         handleSimulatorEvent(event)
	//
	//         // Check if battle is complete
	//         if event.Type == battle.BattleCompleted {
	//             return reportBattleResults(simulator, battleID)
	//         }
	//
	//     case <-ticker.C:
	//         // Periodic status update
	//         reportBattleStatus(simulator, battleID)
	//     }
	// }

	fmt.Println("TODO: Implement monitorBattleProgress")
	fmt.Println("This should monitor battle events and provide real-time updates")
	return nil
}

// handleSimulatorEvent processes individual simulator events
// LEARNING: Event handling and display patterns
func handleSimulatorEvent(event battle.SimulatorEvent) {
	// TODO: Implement event handling
	// Different event types should be handled differently:
	// switch event.Type {
	// case battle.BattleCreated:
	//     fmt.Printf("[%s] ⚔️  Battle created\n", formatTime(event.Timestamp))
	//
	// case battle.BattleCompleted:
	//     fmt.Printf("[%s] 🏁 Battle completed\n", formatTime(event.Timestamp))
	//
	// case battle.TickProcessed:
	//     // Don't log every tick, too noisy
	//
	// default:
	//     fmt.Printf("[%s] 📡 Event: %v\n", formatTime(event.Timestamp), event.Type)
	// }

	fmt.Printf("TODO: Handle event of type %T\n", event)
}

// reportBattleStatus provides periodic status updates
// LEARNING: Status reporting and monitoring patterns
func reportBattleStatus(simulator *battle.BattleSimulator, battleID string) {
	// TODO: Implement status reporting
	// Steps needed:
	// 1. Get current battle status
	// status, err := simulator.GetBattleStatus(battleID)
	// if err != nil {
	//     fmt.Printf("❌ Error getting battle status: %v\n", err)
	//     return
	// }

	// 2. Display key information
	// fmt.Printf("⏱️  Battle Duration: %v\n", status.Duration)
	//
	// // Count units by faction
	// for faction, units := range status.Participants {
	//     aliveCount := 0
	//     for _, unit := range units {
	//         if unit.GetState() != types.Dead {
	//             aliveCount++
	//         }
	//     }
	//     fmt.Printf("🪖 %s: %d/%d units alive\n", faction, aliveCount, len(units))
	// }
	//
	// // Show objective progress
	// for i, obj := range status.Objectives {
	//     fmt.Printf("🎯 Objective %d: %.1f%% complete\n", i+1, obj.Progress*100)
	// }
	// fmt.Println("---")

	fmt.Println("TODO: Implement reportBattleStatus for", battleID)
}

// reportBattleResults provides final battle analysis
// LEARNING: Results analysis and reporting
func reportBattleResults(simulator *battle.BattleSimulator, battleID string) error {
	// TODO: Implement results reporting
	// Steps needed:
	// 1. Get final battle status
	// status, err := simulator.GetBattleStatus(battleID)
	// if err != nil {
	//     return fmt.Errorf("failed to get final status: %w", err)
	// }

	// 2. Determine winner
	// winner := determineWinner(status.Participants)

	// 3. Display comprehensive results
	// fmt.Println("\n🏆 BATTLE RESULTS 🏆")
	// fmt.Printf("Winner: %s\n", winner)
	// fmt.Printf("Duration: %v\n", status.Duration)
	// fmt.Println()
	//
	// // Show casualties
	// fmt.Println("📊 Casualties:")
	// for faction, units := range status.Participants {
	//     total := len(units)
	//     alive := 0
	//     for _, unit := range units {
	//         if unit.GetState() != types.Dead {
	//             alive++
	//         }
	//     }
	//     casualties := total - alive
	//     fmt.Printf("  %s: %d/%d lost (%.1f%%)\n",
	//         faction, casualties, total, float64(casualties)/float64(total)*100)
	// }
	//
	// // Show statistics if available
	// if len(status.Statistics.TotalDamageDealt) > 0 {
	//     fmt.Println("\n⚔️ Combat Statistics:")
	//     for faction, damage := range status.Statistics.TotalDamageDealt {
	//         fmt.Printf("  %s total damage: %d\n", faction, damage)
	//     }
	// }

	fmt.Println("TODO: Implement reportBattleResults for", battleID)
	return nil
}

// Helper functions

// determineWinner analyzes battle participants to determine the winner
func determineWinner(participants map[string][]*types.Unit) string {
	// TODO: Implement winner determination logic
	// Check which faction has units still alive
	// Return "Draw" if multiple factions have survivors
	// Return "None" if all units are dead

	return "TODO: Determine winner"
}

// formatTime formats a timestamp for display
func formatTime(t time.Time) string {
	return t.Format("15:04:05")
}

// LEARNING EXERCISES for Students:
//
// After implementing the TODOs above, try these exercises:
//
// 1. UNIT BEHAVIOR: Modify unit stats (health, damage) and observe battle outcomes
// 2. FORMATIONS: Implement different starting formations for units
// 3. SPECIAL ABILITIES: Add special unit abilities that activate during combat
// 4. TERRAIN: Implement terrain effects that influence battle outcomes
// 5. REINFORCEMENTS: Add mid-battle reinforcements for one side
// 6. MORALE: Implement unit morale that affects combat effectiveness
// 7. RESOURCE COSTS: Add resource costs for unit actions
// 8. AI STRATEGIES: Implement different AI behaviors for units
//
// Each exercise teaches different aspects of concurrent programming:
// - State management
// - Event-driven programming
// - Resource coordination
// - Real-time processing
// - Complex system interactions

