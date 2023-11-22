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
		directions = append(directions, *newSpotDirection("right", 90))
	}
	if row == 2 {
		directions = append(directions, *newSpotDirection("right", 190))
	}
	if row == 3 {
		directions = append(directions, *newSpotDirection("right", 280))
	}
	if row == 4 {
		directions = append(directions, *newSpotDirection("right", 370))
	}
	if row == 5 {
		directions = append(directions, *newSpotDirection("right", 460))
	}
	if row == 6 {
		directions = append(directions, *newSpotDirection("right", 550))
	}
	if row == 7 {
		directions = append(directions, *newSpotDirection("right", 640))
	}
	if row == 8 {
		directions = append(directions, *newSpotDirection("right", 730))
	}
	if row == 9 {
		directions = append(directions, *newSpotDirection("right", 820))
	}
	if row == 10 {
		directions = append(directions, *newSpotDirection("right", 910))
	}

	directions = append(directions, *newSpotDirection("up", y+4))
	directions = append(directions, *newSpotDirection("right", x+2))

	return &directions
}

func directionsForLeaving() *[]SpotDirection {
	var directions []SpotDirection

	directions = append(directions, *newSpotDirection("up", 40))
	directions = append(directions, *newSpotDirection("left", 50))
	directions = append(directions, *newSpotDirection("down", 120))

	return &directions
}
