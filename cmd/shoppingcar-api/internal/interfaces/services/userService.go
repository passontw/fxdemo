package servicesInterface

// ==================== 介面定義 ====================

type UserServiceInterface interface {
	CreateUser(name string) string
	GetOrderCount(userId string) int
}

type OrderServiceInterface interface {
	CreateOrder(userId string) string
	GetUserName(userId string) string
}

// 協調器介面 - 負責協調 User 和 Order 服務間的互動
type CoordinatorInterface interface {
	CreateUserWithOrder(name string) (userID string, orderID string)
	CreateOrderForUser(userId string) string
}
