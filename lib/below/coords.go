package below

import (
	"math/rand"
)

type Coords [2]int

func (coords Coords) X() int {
	return coords[0]
}

func (coords Coords) Y() int {
	return coords[1]
}

func RandomCoords() Coords {
	return Coords{rand.Intn(WORLD_COLS), rand.Intn(WORLD_ROWS)}
}

func DestinationCoords(coords Coords, dir string) Coords {
	return offsetCoords(coords, dirToOffset(dir))
}

func offsetCoords(coords Coords, dcoords Coords) Coords {
	for idx, coord := range coords {
		coords[idx] = coord + dcoords[idx]
	}
	return coords
}

func dirToOffset(dir string) Coords {
	switch dir {
	case "w":
		return Coords{-1, 0}
	case "e":
		return Coords{1, 0}
	case "n":
		return Coords{0, -1}
	case "s":
		return Coords{0, 1}
	case "nw":
		return Coords{-1, -1}
	case "ne":
		return Coords{1, -1}
	case "sw":
		return Coords{-1, 1}
	case "se":
		return Coords{1, 1}
	}
	return Coords{0, 0}
}
