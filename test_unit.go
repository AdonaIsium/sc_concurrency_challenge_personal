package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/AdonaIsium/sc_concurrency_challenge_personal/internal/types"
)

func main() {
	fmt.Println("🎮 Testing Unit.run() goroutine loop...")
	fmt.Println()

	var wg sync.WaitGroup

	// Create a Marine
	marine := types.NewUnit("marine-1", types.Marine, types.Position{X: 100, Y: 100}, &wg)
	fmt.Printf("✅ Created %s at position (%.1f, %.1f)\n", marine.Type, 100.0, 100.0)
	fmt.Println()

	// Give it some time to start
	time.Sleep(100 * time.Millisecond)

	// Create a Zergling enemy
	zergling := types.NewUnit("zergling-1", types.Zergling, types.Position{X: 150, Y: 150}, &wg)
	fmt.Printf("✅ Created %s at position (%.1f, %.1f) with %d HP\n", zergling.Type, 150.0, 150.0, zergling.GetHealth())
	fmt.Println()

	// Send some commands
	fmt.Println("📡 Sending commands to marine...")

	err := marine.SendCommand(types.Command{
		Type: types.CmdMove,
		Dest: types.Position{X: 200, Y: 200},
	})
	if err != nil {
		fmt.Printf("❌ Error sending move command: %v\n", err)
	}

	err = marine.SendCommand(types.Command{
		Type:   types.CmdAttack,
		Target: zergling,
	})
	if err != nil {
		fmt.Printf("❌ Error sending attack command: %v\n", err)
	}

	err = marine.SendCommand(types.Command{
		Type: types.CmdStop,
	})
	if err != nil {
		fmt.Printf("❌ Error sending stop command: %v\n", err)
	}

	// Wait a bit to see the output
	time.Sleep(100 * time.Millisecond)

	fmt.Printf("\n💔 Zergling health after attack: %d HP\n", zergling.GetHealth())

	// Shutdown the units
	fmt.Println()
	fmt.Println("🛑 Shutting down units...")
	marine.Shutdown()
	zergling.Shutdown()

	// Wait for goroutines to finish
	wg.Wait()

	fmt.Println("✅ Marine shut down cleanly!")
	fmt.Println()
	fmt.Println("🎯 Your run() loop is working! You've mastered the goroutine lifecycle!")
}
