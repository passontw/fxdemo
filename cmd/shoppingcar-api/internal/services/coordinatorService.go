package services

import (
	"fmt"
	"fxdemo/cmd/shoppingcar-api/internal/interfaces/interfaces"
)

// ==================== Coordinator 實作 ====================
type CoordinatorService struct {
	userService  interfaces.UserServiceInterface
	orderService interfaces.OrderServiceInterface
}

func NewCoordinator(userService interfaces.UserServiceInterface, orderService interfaces.OrderServiceInterface) interfaces.CoordinatorInterface {
	return &CoordinatorService{
		userService:  userService,
		orderService: orderService,
	}
}

func (c *CoordinatorService) CreateUserWithOrder(name string) (userID string, orderID string) {
	fmt.Printf("Coordinator: 開始建立使用者及其預設訂單\n")

	// 建立使用者
	userID = c.userService.CreateUser(name)

	// 為使用者建立預設訂單
	orderID = c.orderService.CreateOrder(userID)
	fmt.Printf("Coordinator: 為使用者建立了訂單 %s\n", orderID)

	return userID, orderID
}

func (c *CoordinatorService) CreateOrderForUser(userId string) string {
	fmt.Printf("Coordinator: 為使用者 %s 建立訂單\n", userId)

	// 取得使用者的訂單數量
	count := c.userService.GetOrderCount(userId)
	fmt.Printf("Coordinator: 使用者已有 %d 個訂單\n", count)

	// 建立新訂單
	orderID := c.orderService.CreateOrder(userId)

	return orderID
}
