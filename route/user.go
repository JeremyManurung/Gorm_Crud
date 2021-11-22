package route

import (
	"res-task/controller"
	"github.com/labstack/echo"
)

func NewRoutes(app *echo.Echo) {
	app.GET("/user/:id", controller.GetUserbyIDC)
	app.GET("/user", controller.GetUserC)
	app.POST("/user", controller.SaveBookC)
	app.PUT("/user/:id", controller.UpdateUserbyIDC)
	app.DELETE("/user/:id", controller.DeleteUserbyIDC)
}
