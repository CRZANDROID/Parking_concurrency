package models

type SpotDirection struct {
	Direction string
	Point     float64
}

func newSpotDirection(direction string, point float64) *SpotDirection {
	return &SpotDirection{
		Direction: direction,
		Point:     point,
	}
}
