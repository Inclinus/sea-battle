package stats

import (
	"fmt"
	"encoding/json"
	"os"
)

type Stats struct {
	GamesWon uint
	GamesLost uint
	ShotsHit uint
	ShotsMissed uint
	BoatsDestroyed uint
}

// Get the stats from the stats.json file
func GetStats() *Stats {
	// Open json file
	content, err := os.ReadFile("stats.json")
	if err != nil {
		panic(err)
	}

	// Unmarshal json file
	var stats Stats
	err = json.Unmarshal(content, &stats)
	if err != nil {
		panic(err)
	}

	return &stats
}

// Overwrite the stats.json file with given stats
func SaveStats(stats Stats) {
	// Open json file
	content, err := os.ReadFile("stats.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// Marshal json file
	content, err = json.Marshal(stats)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)

	// Write json file
	err = os.WriteFile("stats.json", content, 0644)
	if err != nil {
		panic(err)
	}
}	
