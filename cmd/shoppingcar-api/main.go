package main

import (
	"fmt"
	"fxdemo/cmd/shoppingcar-api/internal/interfaces/interfaces"
	"fxdemo/cmd/shoppingcar-api/internal/services"

	"go.uber.org/fx"
)

// ==================== Coordinator 實作 ====================
type Coordinator struct {
	userService  interfaces.UserServiceInterface
	orderService interfaces.OrderServiceInterface
}

func NewCoordinator(userService interfaces.UserServiceInterface, orderService interfaces.OrderServiceInterface) interfaces.CoordinatorInterface {
	return &Coordinator{
		userService:  userService,
		orderService: orderService,
	}
}

func (c *Coordinator) CreateUserWithOrder(name string) (userID string, orderID string) {
	fmt.Printf("Coordinator: 開始建立使用者及其預設訂單\n")

	// 建立使用者
	userID = c.userService.CreateUser(name)

	// 為使用者建立預設訂單
	orderID = c.orderService.CreateOrder(userID)
	fmt.Printf("Coordinator: 為使用者建立了訂單 %s\n", orderID)

	return userID, orderID
}

func (c *Coordinator) CreateOrderForUser(userId string) string {
	fmt.Printf("Coordinator: 為使用者 %s 建立訂單\n", userId)

	// 取得使用者的訂單數量
	count := c.userService.GetOrderCount(userId)
	fmt.Printf("Coordinator: 使用者已有 %d 個訂單\n", count)

	// 建立新訂單
	orderID := c.orderService.CreateOrder(userId)

	return orderID
}

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
	fx.Provide(fx.Annotate(NewCoordinator, fx.As(new(interfaces.CoordinatorInterface)))),
)

var AppModule = fx.Module("app",
	fx.Provide(NewApp),
	fx.Invoke(func(app *App) {
		app.Run()
	}),
)

// ==================== 主程式 ====================
func main() {
	fx.New(
		UserModule,
		OrderModule,
		CoordinatorModule,
		AppModule,
	).Run()
}
