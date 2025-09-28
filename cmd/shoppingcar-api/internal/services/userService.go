package services

import "fmt"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(name string) string {
	userID := fmt.Sprintf("user_%s", name)
	fmt.Printf("UserService: 建立使用者 %s\n", userID)
	return userID
}

func (s *UserService) GetOrderCount(userID string) int {
	fmt.Printf("UserService: 取得使用者 %s 的訂單數量\n", userID)
	return 2 // 模擬數據
}
