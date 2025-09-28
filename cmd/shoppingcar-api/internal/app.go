package app

import (
	"fmt"
	"fxdemo/cmd/shoppingcar-api/internal/interfaces/interfaces"
	"fxdemo/cmd/shoppingcar-api/internal/services"

	"go.uber.org/fx"
)

// ==================== App 主邏輯 ====================
type App struct {
	coordinator  interfaces.CoordinatorInterface
	orderService interfaces.OrderServiceInterface
}

func NewApp(coordinator interfaces.CoordinatorInterface, orderService interfaces.OrderServiceInterface) *App {
	return &App{
		coordinator:  coordinator,
		orderService: orderService,
	}
}

func (app *App) Run() {
	fmt.Println("=== 開始測試模組互相調用 ===")

	// 測試 Coordinator 協調 UserService 和 OrderService
	userID, orderID := app.coordinator.CreateUserWithOrder("Alice")

	fmt.Println()

	// 測試 Coordinator 為現有使用者建立訂單
	orderID2 := app.coordinator.CreateOrderForUser("user_Bob")
	userName := app.orderService.GetUserName("user_Bob")

	fmt.Printf("\n=== 測試結果 ===\n")
	fmt.Printf("建立的使用者: %s\n", userID)
	fmt.Printf("建立的訂單1: %s\n", orderID)
	fmt.Printf("建立的訂單2: %s\n", orderID2)
	fmt.Printf("查詢的使用者名稱: %s\n", userName)
}

// ==================== FX 模組定義 ====================
var UserModule = fx.Module("user",
	fx.Provide(fx.Annotate(services.NewUserService, fx.As(new(interfaces.UserServiceInterface)))),
)

var OrderModule = fx.Module("order",
	fx.Provide(fx.Annotate(services.NewOrderService, fx.As(new(interfaces.OrderServiceInterface)))),
)

var CoordinatorModule = fx.Module("coordinator",
	fx.Provide(fx.Annotate(services.NewCoordinator, fx.As(new(interfaces.CoordinatorInterface)))),
)

var AppModule = fx.Module("app",
	fx.Provide(NewApp),
	fx.Invoke(func(app *App) {
		app.Run()
	}),
)
