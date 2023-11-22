package scenes

import (
	"sync"
	"time"

	"image/color"
	"parking-concurrency/models"
	"parking-concurrency/utils"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/scene"
)

var (
	spots = []*models.Spot{
		models.NewSpot(110, 80, 140, 110, 1, 2),
		models.NewSpot(110, 120, 140, 150, 1, 1),

		models.NewSpot(210, 80, 240, 110, 2, 3),
		models.NewSpot(210, 120, 240, 150, 2, 4),

		models.NewSpot(300, 80, 330, 110, 3, 5),
		models.NewSpot(300, 120, 330, 150, 3, 2),

		models.NewSpot(390, 80, 420, 110, 4, 2),
		models.NewSpot(390, 120, 420, 150, 4, 2),

		models.NewSpot(480, 80, 510, 110, 5, 2),
		models.NewSpot(480, 120, 510, 150, 5, 2),

		models.NewSpot(570, 80, 600, 110, 6, 2),
		models.NewSpot(570, 120, 600, 150, 6, 2),

		models.NewSpot(660, 80, 690, 110, 7, 2),
		models.NewSpot(660, 120, 690, 150, 7, 2),

		models.NewSpot(750, 80, 780, 110, 8, 2),
		models.NewSpot(750, 120, 780, 150, 8, 2),

		models.NewSpot(840, 80, 870, 110, 9, 2),
		models.NewSpot(840, 120, 870, 150, 9, 2),

		models.NewSpot(930, 80, 960, 110, 10, 2),
		models.NewSpot(930, 120, 960, 150, 10, 2),
	}
	parking = models.NewParking(spots)
	doorM   sync.Mutex
	manager = models.NewManager()
)

type MainScene struct{}

func NewMainScene() *MainScene {
	return &MainScene{}
}

func (ps *MainScene) Start() {
	band := true

	_ = oak.AddScene("mainScene", scene.Scene{
		Start: func(ctx *scene.Context) {
			_ = ctx.Window.SetBorderless(true)
			prepare(ctx)

			event.GlobalBind(ctx, event.Enter, func(enterPayload event.EnterPayload) event.Response {
				if !band {
					return 0
				}

				band = false

				for {
					go func() {
						car := models.NewCar(ctx)
						car.Run(manager, parking, &doorM)
					}()
					time.Sleep(time.Millisecond * time.Duration(utils.RandomInt(1000, 2000)))
				}
			})
		},
	})
}

func prepare(ctx *scene.Context) {
	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(0, 0, 1000, 1000)), entities.WithColor(color.RGBA{R: 0, G: 128, B: 0, A: 255}), entities.WithDrawLayers([]int{0}))

	parkingArea := floatgeom.NewRect2(50, 40, 990, 180)
	entities.New(ctx, entities.WithRect(parkingArea), entities.WithColor(color.RGBA{R: 86, G: 101, B: 115, A: 255}), entities.WithDrawLayers([]int{0}))

	parkingDoor := floatgeom.NewRect2(45, 110, 50, 180)
	entities.New(ctx, entities.WithRect(parkingDoor), entities.WithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}), entities.WithDrawLayers([]int{0}))

	for _, spot := range spots {
		entities.New(ctx, entities.WithRect(*spot.GetArea()), entities.WithColor(color.RGBA{R: 139, G: 128, B: 0, A: 255}), entities.WithDrawLayers([]int{1}))
	}
}
