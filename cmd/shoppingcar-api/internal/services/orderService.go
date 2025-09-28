package services

import (
	"fmt"
	"fxdemo/cmd/shoppingcar-api/internal/interfaces/interfaces"
)

type OrderService struct{}

func NewOrderService() interfaces.OrderServiceInterface {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(userID string) string {
	orderID := fmt.Sprintf("order_%s_001", userID)
	fmt.Printf("OrderService: 建立訂單 %s\n", orderID)
	return orderID
}

func (s *OrderService) GetUserName(userID string) string {
	fmt.Printf("OrderService: 取得使用者 %s 的名稱\n", userID)
	return "John Doe" // 模擬數據
}
