package models

import (
	"github.com/oakmound/oak/v4/render"
	"github.com/oakmound/oak/v4/render/mod"
	"parking-concurrency/utils"
	"sync"
	"time"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/scene"
)

const (
	entrySpot = 185.00
	velocity  = 15
)

type Car struct {
	area   floatgeom.Rect2
	entity *entities.Entity
	mu     sync.Mutex
}

func NewCar(ctx *scene.Context) *Car {
	area := floatgeom.NewRect2(445, -20, 465, 0)

	assetPath := "assets/car.png"

	carModel, _ := render.LoadSprite(assetPath)

	newSwitch := render.NewSwitch("up", map[string]render.Modifiable{
		"up":    carModel,
		"down":  carModel.Copy().Modify(mod.FlipY),
		"left":  carModel.Copy().Modify(mod.Rotate(90)),
		"right": carModel.Copy().Modify(mod.Rotate(-90)),
	})

	entity := entities.New(ctx, entities.WithRect(area), entities.WithRenderable(newSwitch), entities.WithDrawLayers([]int{2, 3}))

	return &Car{
		area:   area,
		entity: entity,
	}
}

func (c *Car) ShiftY(dy float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.ShiftY(dy)
}

func (c *Car) ShiftX(dx float64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.ShiftX(dx)
}

func (c *Car) X() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entity.X()
}

func (c *Car) Y() float64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.entity.Y()
}

func (c *Car) Remove() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entity.Destroy()
}

func (c *Car) isCollision(direction string, cars []*Car) bool {
	minDistance := 30.0
	for _, car := range cars {
		if direction == "left" && c.X() > car.X() && c.X()-car.X() < minDistance && c.Y() == car.Y() {
			return true
		} else if direction == "right" && c.X() < car.X() && car.X()-c.X() < minDistance && c.Y() == car.Y() {
			return true
		} else if direction == "up" && c.Y() > car.Y() && c.Y()-car.Y() < minDistance && c.X() == car.X() {
			return true
		} else if direction == "down" && c.Y() < car.Y() && car.Y()-c.Y() < minDistance && c.X() == car.X() {
			return true
		}
	}
	return false
}

func (c *Car) Enqueue(manager *Manager) {
	for c.Y() < 145 {
		if !c.isCollision("down", manager.GetCars()) {
			c.ShiftY(1)
			c.entity.Renderable.(*render.Switch).Set("down")
			time.Sleep(velocity * time.Millisecond)
		}
	}
}

func (c *Car) JoinDoor(manager *Manager) {
	for c.Y() < entrySpot {
		if !c.isCollision("down", manager.GetCars()) {
			c.ShiftY(1)
			c.entity.Renderable.(*render.Switch).Set("down")
			time.Sleep(velocity * time.Millisecond)
		}
	}
}

func (c *Car) ExitDoor(manager *Manager) {
	for c.Y() > 145 {
		if !c.isCollision("up", manager.GetCars()) {
			c.ShiftY(-1)
			c.entity.Renderable.(*render.Switch).Set("up")
			time.Sleep(velocity * time.Millisecond)
		}
	}
}

func (c *Car) Park(spot *Spot, manager *Manager) {
	for index := 0; index < len(*spot.GetDirectionsForParking()); index++ {
		directions := *spot.GetDirectionsForParking()

		switch directions[index].Direction {
		case "right":
			for c.X() < directions[index].Point {
				if !c.isCollision("right", manager.GetCars()) {
					c.ShiftX(1)
					c.entity.Renderable.(*render.Switch).Set("right")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		case "down":
			for c.Y() < directions[index].Point {
				if !c.isCollision("down", manager.GetCars()) {
					c.ShiftY(1)
					c.entity.Renderable.(*render.Switch).Set("down")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		case "left":
			for c.X() > directions[index].Point {
				if !c.isCollision("left", manager.GetCars()) {
					c.ShiftX(-1)
					c.entity.Renderable.(*render.Switch).Set("left")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		case "up":
			for c.Y() > directions[index].Point {
				if !c.isCollision("up", manager.GetCars()) {
					c.ShiftY(-1)
					c.entity.Renderable.(*render.Switch).Set("up")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		}
	}
}

func (c *Car) Leave(spot *Spot, manager *Manager) {
	for index := 0; index < len(*spot.GetDirectionsForLeaving()); index++ {
		directions := *spot.GetDirectionsForLeaving()

		switch directions[index].Direction {
		case "left":
			for c.X() > directions[index].Point {
				if !c.isCollision("left", manager.GetCars()) {
					c.ShiftX(-1)
					c.entity.Renderable.(*render.Switch).Set("left")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		case "right":
			for c.X() < directions[index].Point {
				if !c.isCollision("right", manager.GetCars()) {
					c.ShiftX(1)
					c.entity.Renderable.(*render.Switch).Set("right")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		case "up":
			for c.Y() > directions[index].Point {
				if !c.isCollision("up", manager.GetCars()) {
					c.ShiftY(-1)
					c.entity.Renderable.(*render.Switch).Set("up")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		case "down":
			for c.Y() < directions[index].Point {
				if !c.isCollision("down", manager.GetCars()) {
					c.ShiftY(1)
					c.entity.Renderable.(*render.Switch).Set("down")
					time.Sleep(velocity * time.Millisecond)
				}
			}
		}
	}
}

func (c *Car) LeaveSpot(manager *Manager) {
	spotX := c.X()
	for c.X() > spotX-30 {
		if !c.isCollision("left", manager.GetCars()) {
			c.ShiftX(-1)
			c.entity.Renderable.(*render.Switch).Set("left")
			time.Sleep(velocity * time.Millisecond)
		}
		time.Sleep(20 * time.Millisecond)
	}
}

func (c *Car) GoAway(manager *Manager) {
	for c.Y() > -20 {
		if !c.isCollision("up", manager.GetCars()) {
			c.ShiftY(-1)
			c.entity.Renderable.(*render.Switch).Set("up")
			time.Sleep(velocity * time.Millisecond)
		}
	}
}

func (c *Car) Run(manager *Manager, parking *Parking, doorM *sync.Mutex) {
	manager.Add(c)
	c.Enqueue(manager)

	spotAvailable := parking.GetParkingSpotAvailable()

	doorM.Lock()
	c.JoinDoor(manager)
	doorM.Unlock()

	c.Park(spotAvailable, manager)
	time.Sleep(time.Millisecond * time.Duration(utils.RandomInt(40000, 50000)))

	c.LeaveSpot(manager)
	parking.ReleaseParkingSpot(spotAvailable)

	c.Leave(spotAvailable, manager)

	doorM.Lock()
	c.ExitDoor(manager)
	doorM.Unlock()

	c.GoAway(manager)
	c.Remove()

	manager.RemoveCar(c)
}
