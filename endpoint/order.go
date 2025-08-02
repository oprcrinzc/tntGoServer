package endpoint

import (
	"context"
	"usersys/db"
	"usersys/def"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func Order(c *fiber.Ctx) error {
	order := def.Order{}
	cnt := c.FormValue("Content")
	color := c.FormValue("Color")
	material := c.FormValue("Material")
	multiP, err := c.FormFile("File")
	if err != nil {
		log.Error(err)
	}

	log.Info(multiP.Filename)

	c.SaveFile(multiP, "/mnt/game/projects/tnt3dPrint/UserSys/storage/"+multiP.Filename)
	// file := "asdasd"
	order.Content = cnt
	order.Color = color
	order.Material = material
	order.File = []string{multiP.Filename}

	log.Info(cnt)
	if cnt == "" {
		return c.Status(200).JSON("NO")
	}
	// err := c.BodyParser(&order)
	token := c.Get("Authorization")
	// log.Info(token)

	if token != "" {
		t, err := def.VerifyJwt(token)
		if err != nil {
			return c.JSON(err.Error())
		}
		order.Customer = t.Claims.(jwt.MapClaims)["username"].(string)
		// n, err := t.Claims.GetExpirationTime()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		client := db.Conn()
		defer func() {
			err := client.Disconnect(context.TODO())
			if err != nil {
				log.Error(err)
			}
		}()

		col := client.Database("tnt").Collection("order")
		iRes, err := col.InsertOne(context.TODO(), order)
		if err != nil {
			log.Error(err)
			return c.JSON(err.Error())
		}
		if iRes.Acknowledged {
			return c.JSON(iRes.Acknowledged)
		}

	}

	if err != nil {
		return c.JSON(err.Error())
	}
	return c.JSON(order)
}
