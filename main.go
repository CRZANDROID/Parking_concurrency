package main

import (
	"github.com/oakmound/oak/v4"
	"parking-concurrency/scenes"
)

func main() {
	parkingScene := scenes.NewMainScene()

	parkingScene.Start()

	_ = oak.Init("mainScene")
}
