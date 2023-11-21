package models

import "sync"

type Manager struct {
	Cars  []*Car
	Mutex sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Cars: make([]*Car, 0),
	}
}

func (cm *Manager) Add(car *Car) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()
	cm.Cars = append(cm.Cars, car)
}

func (cm *Manager) GetCars() []*Car {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()
	return cm.Cars
}

func (cm *Manager) RemoveCar(car *Car) {
	cm.Mutex.Lock()
	defer cm.Mutex.Unlock()
	for i, c := range cm.Cars {
		if c == car {
			cm.Cars = append(cm.Cars[:i], cm.Cars[i+1:]...)
			break
		}
	}
}
