package middlewares

import (
	"context"
	"fmt"
	"go-playground/firebase_client"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckAccessToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("accessToken")

		if err != nil {
			fmt.Println("Error", err)
			return c.NoContent(http.StatusUnauthorized)
		}

		fbClient := firebase_client.GetFirebaseClient()

		token := cookie.Value
		ctx := context.Background()

		userInfo, err := fbClient.VerifyIDTokenAndCheckRevoked(ctx, token)
		if err != nil {
			fmt.Println("Error", token, err)
			return c.NoContent(http.StatusUnauthorized)
		}

		//Đặt giá trị người dùng vào request context
		c.Set("uid", userInfo.UID)

		return next(c)
	}
}
