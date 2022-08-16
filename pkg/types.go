package pkg

type ArenaUpdate struct {
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
	} `json:"_links"`
	Arena struct {
		Dimensions []int                  `json:"dims"`
		State      map[string]PlayerState `json:"state"`
	} `json:"arena"`
}

type PlayerState struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Direction string `json:"direction"`
	WasHit    bool   `json:"wasHit"`
	Score     int    `json:"score"`
}

type Player interface {
	Play(ArenaUpdate) string
}

const (
	DirectionNorth = "N"
	DirectionWest = "W"
	DirectionSouth = "S"
	DirectionEast = "E"

	MoveForward = "F"
	MoveThrow = "T"
	MoveTurnLeft = "L"
	MoveTurnRight = "R"
)

var (
	Directions = []string{DirectionNorth, DirectionWest, DirectionSouth, DirectionEast}
	Moves = []string{MoveForward, MoveThrow, MoveTurnLeft, MoveTurnRight}
)
