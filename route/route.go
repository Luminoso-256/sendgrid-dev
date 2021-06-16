package route

import (
	"github.com/labstack/echo"
	"github.com/yKanazawa/sendgrid-dev/api/v3/send"
)

func Init() *echo.Echo {
	e := echo.New()

	// Routes
	v3 := e.Group("/v3")
	{
		v3.GET("/send", send.GetSend())
		v3.POST("/send", send.PostSend())
	}
	
	return e
}
