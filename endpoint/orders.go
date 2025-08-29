package endpoint

import (
	"context"
	"usersys/db"
	"usersys/def"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Orders(c *fiber.Ctx) error {
	customer := ""
	token := c.Get("Authorization")
	if token != "" {
		t, err := def.VerifyJwt(token)
		if err != nil {
			return c.JSON(err.Error())
		}
		customer = t.Claims.(jwt.MapClaims)["username"].(string)
	}

	if customer != "" {
		client := db.Conn()
		defer func() {
			client.Disconnect(context.TODO())
		}()
		col := client.Database("tnt").Collection("order")
		cursor, err := col.Find(context.TODO(), bson.M{"customer": customer})
		if err != nil {
			return c.Status(500).JSON(def.Msg{
				Header:  "Database Error",
				Content: err.Error(),
			})
		}
		var orders def.Orders
		err = cursor.All(context.TODO(), &orders)
		if err != nil {
			return c.Status(500).JSON(def.Msg{
				Header:  "Database Error",
				Content: err.Error(),
			})
		}
		// log.Info(orders)
		return c.JSON(orders)
	}
	return c.Status(200).JSON("ERROR")
}
