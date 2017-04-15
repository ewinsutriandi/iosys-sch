package main

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const secretKey string = "ular lari lurus"

func login(c echo.Context) error {
	username := c.FormValue("username")
	//password := c.FormValue("password")
	isvaliduser := true
	if isvaliduser {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = username
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte(secretKey))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	mdjwt := middleware.JWT([]byte(secretKey))

	// Login route
	e.POST("/login", login)

	// Unauthenticated route
	e.GET("/", accessible)

	//restricted route group for authenticated user
	ru := e.Group("/restricted")
	ru.Use(mdjwt)
	//ru.GET("/user/:usr", getUser)
	//ru.POST("/telem", putTelemReport)

	e.Logger.Fatal(e.Start(":1323"))
}
