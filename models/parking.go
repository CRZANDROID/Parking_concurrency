package models

import (
	"sync"
)

type Parking struct {
	spots         []*Spot
	mu            sync.Mutex
	availableCond *sync.Cond
}

func NewParking(spots []*Spot) *Parking {
	p := &Parking{
		spots: spots,
	}
	p.availableCond = sync.NewCond(&p.mu)
	return p
}

func (p *Parking) GetSpots() []*Spot {
	return p.spots
}

func (p *Parking) GetParkingSpotAvailable() *Spot {
	p.mu.Lock()
	defer p.mu.Unlock()

	for {
		for _, spot := range p.spots {
			if spot.GetIsAvailable() {
				spot.SetIsAvailable(false)
				return spot
			}
		}
		p.availableCond.Wait()
	}
}

func (p *Parking) ReleaseParkingSpot(spot *Spot) {
	p.mu.Lock()
	defer p.mu.Unlock()

	spot.SetIsAvailable(true)
	p.availableCond.Signal()
}
