package main

import (
	app "fxdemo/cmd/shoppingcar-api/internal"

	"go.uber.org/fx"
)

// ==================== 主程式 ====================
func main() {
	fx.New(
		app.UserModule,
		app.OrderModule,
		app.CoordinatorModule,
		app.AppModule,
	).Run()
}
