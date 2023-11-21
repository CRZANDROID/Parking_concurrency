package models

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

type Spot struct {
	area                 *floatgeom.Rect2
	directionsForParking *[]SpotDirection
	directionsForLeaving *[]SpotDirection
	number               int
	isAvailable          bool
}

func NewSpot(x, y, x2, y2 float64, row, number int) *Spot {
	area := floatgeom.NewRect2(x, y, x2, y2)

	return &Spot{
		area:                 &area,
		directionsForParking: directionForParking(x, y, row),
		directionsForLeaving: directionsForLeaving(),
		number:               number,
		isAvailable:          true,
	}
}

func (p *Spot) GetArea() *floatgeom.Rect2 {
	return p.area
}

func (p *Spot) GetNumber() int {
	return p.number
}

func (p *Spot) GetDirectionsForParking() *[]SpotDirection {
	return p.directionsForParking
}

func (p *Spot) GetDirectionsForLeaving() *[]SpotDirection {
	return p.directionsForLeaving
}

func (p *Spot) GetIsAvailable() bool {
	return p.isAvailable
}

func (p *Spot) SetIsAvailable(isAvailable bool) {
	p.isAvailable = isAvailable
}

func directionForParking(x, y float64, row int) *[]SpotDirection {
	var directions []SpotDirection

	if row == 1 {
		directions = append(directions, *newSpotDirection("left", 445))
	}
	if row == 2 {
		directions = append(directions, *newSpotDirection("left", 355))
	}
	if row == 3 {
		directions = append(directions, *newSpotDirection("left", 265))
	}
	if row == 4 {
		directions = append(directions, *newSpotDirection("left", 175))
	}
	if row == 5 {
		directions = append(directions, *newSpotDirection("left", 85))
	}

	directions = append(directions, *newSpotDirection("down", y+5))
	directions = append(directions, *newSpotDirection("left", x+5))

	return &directions
}

func directionsForLeaving() *[]SpotDirection {
	var directions []SpotDirection

	directions = append(directions, *newSpotDirection("down", 380))
	directions = append(directions, *newSpotDirection("right", 475))
	directions = append(directions, *newSpotDirection("up", 185))

	return &directions
}
