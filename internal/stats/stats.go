package stats

import (
	"encoding/json"
	"os"
)

type Stats struct {
	GamesWon    uint
	GamesLost   uint
	ShotsHit    uint
	ShotsMissed uint
}

// Get the stats from the stats.json file
func GetStats() *Stats {
	// Check if stats.json exists
	_, err := os.Stat("stats.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Notify that their is no saved stats & return empty stats
			return &Stats{}
		} else {
			panic(err)
		}
	}

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
	// Check if stats.json exists
	_, err := os.Stat("stats.json")
	if err != nil {
		if os.IsNotExist(err) {
			// Create stats.json if it doesn't exist
			_, err = os.Create("stats.json")
			if err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}

	// Marshal json file
	content, err := json.Marshal(stats)
	if err != nil {
		panic(err)
	}

	// Write json file
	err = os.WriteFile("stats.json", content, 0644)
	if err != nil {
		panic(err)
	}
}

func AddShotHit() {
	stats := GetStats()
	stats.ShotsHit++
	SaveStats(*stats)
}

func AddShotMissed() {
	stats := GetStats()
	stats.ShotsMissed++
	SaveStats(*stats)
}

func AddGameWon() {
	stats := GetStats()
	stats.GamesWon++
	SaveStats(*stats)
}

func AddGameLost() {
	stats := GetStats()
	stats.GamesLost++
	SaveStats(*stats)
}
