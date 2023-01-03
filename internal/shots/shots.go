package shots

import (
	"sea-battle/internal/utils"
)

var AllShots []Shot

type Shot struct {
	// Player Player => To do in another branch
	Position utils.Position
	Hit      bool
}

func GetAllShots() *[]Shot {
	return &AllShots
}
