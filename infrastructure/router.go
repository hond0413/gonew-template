package infrastructure

import "github.com/labstack/echo"

func Init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":9999"))
}
