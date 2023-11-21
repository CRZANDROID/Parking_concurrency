package scenes

import (
	"sync"
	"time"

	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/scene"
	"image/color"
	"parking-concurrency/models"
	"parking-concurrency/utils"
)

var (
	spots = []*models.Spot{
		models.NewSpot(410, 210, 440, 240, 1, 1),
		models.NewSpot(410, 255, 440, 285, 1, 2),
		models.NewSpot(410, 300, 440, 330, 1, 3),
		models.NewSpot(410, 345, 440, 375, 1, 4),
		models.NewSpot(320, 210, 350, 240, 2, 5),
		models.NewSpot(320, 255, 350, 285, 2, 6),
		models.NewSpot(320, 300, 350, 330, 2, 7),
		models.NewSpot(320, 345, 350, 375, 2, 8),
		models.NewSpot(230, 210, 260, 240, 3, 9),
		models.NewSpot(230, 255, 260, 285, 3, 10),
		models.NewSpot(230, 300, 260, 330, 3, 11),
		models.NewSpot(230, 345, 260, 375, 3, 12),
		models.NewSpot(140, 210, 170, 240, 4, 13),
		models.NewSpot(140, 255, 170, 285, 4, 14),
		models.NewSpot(140, 300, 170, 330, 4, 15),
		models.NewSpot(140, 345, 170, 375, 4, 16),
		models.NewSpot(50, 210, 80, 240, 5, 17),
		models.NewSpot(50, 255, 80, 285, 5, 18),
		models.NewSpot(50, 300, 80, 330, 5, 19),
		models.NewSpot(50, 345, 80, 375, 5, 20),
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

				for i := 0; i < 100; i++ {
					go func() {
						car := models.NewCar(ctx)
						car.Run(manager, parking, &doorM)
					}()
					time.Sleep(time.Millisecond * time.Duration(utils.RandomInt(1000, 2000)))
				}

				return 0
			})
		},
	})
}

func prepare(ctx *scene.Context) {
	// background brown color

	entities.New(ctx, entities.WithRect(floatgeom.NewRect2(0, 0, 800, 500)), entities.WithColor(color.RGBA{R: 139, G: 69, B: 19, A: 255}), entities.WithDrawLayers([]int{0}))

	parkingArea := floatgeom.NewRect2(20, 180, 500, 405)
	entities.New(ctx, entities.WithRect(parkingArea), entities.WithColor(color.RGBA{R: 86, G: 101, B: 115, A: 255}), entities.WithDrawLayers([]int{0}))

	parkingDoor := floatgeom.NewRect2(440, 170, 500, 180)
	entities.New(ctx, entities.WithRect(parkingDoor), entities.WithColor(color.RGBA{R: 255, G: 255, B: 255, A: 255}), entities.WithDrawLayers([]int{0}))

	for _, spot := range spots {
		entities.New(ctx, entities.WithRect(*spot.GetArea()), entities.WithColor(color.RGBA{R: 212, G: 172, B: 13, A: 255}), entities.WithDrawLayers([]int{1}))
	}
}
