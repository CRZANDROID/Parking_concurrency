package main

import (
	"parking-concurrency/scenes"

	"github.com/oakmound/oak/v4"
)

func main() {
	parkingScene := scenes.NewMainScene()

	parkingScene.Start()

	_ = oak.Init("mainScene", func(c oak.Config) (oak.Config, error) {
		c.Screen.Width = 1000
		c.Screen.Height = 200
		return c, nil
	})
}
