package endpoint

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strconv"
	"usersys/db"
	"usersys/def"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func Login(c *fiber.Ctx) error {
	log.Info("Login arrived")
	body := def.UserLogin{}
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	checkNil := 0x0000
	if body.Name == "" {
		checkNil |= 0x0001
	} else if body.Password == "" {
		checkNil |= 0x0100
	}

	if checkNil != 0x0000 {
		return c.Status(400).JSON(def.Msg{
			Header:  "Ingredient Error",
			Content: strconv.Itoa(checkNil),
		})
	}

	sha512Pwd := sha512.Sum512([]byte(body.Password))
	// log.Info(sha512Pwd)
	// return c.JSON(fmt.Sprintf("%x", sha512Pwd))

	// find user from db

	client := db.Conn()
	col := client.Database("tnt").Collection("user")

	// cursor, err := col.Find(context.TODO(), bson.D{{Key: "name", Value: body.Name}, {Key: "password", Value: sha512Pwd}})
	cursor, err := col.Find(context.TODO(), bson.M{
		"name":     body.Name,
		"password": fmt.Sprintf("%x", sha512Pwd),
	})
	if err != nil {
		return c.Status(500).JSON(def.Msg{
			Header:  "Database Error",
			Content: err.Error(),
		})
	}
	var res def.Users
	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return c.Status(500).JSON(def.Msg{
			Header:  "Database Error",
			Content: err.Error(),
		})
	}
	if len(res) == 1 {
		s, err := def.GenJwt("oprc")
		if err != nil {
			c.JSON(err)
		}
		return c.Status(200).JSON(s)
	}

	// token, err := VerifyJwt(s)
	// if err != nil {
	// 	return c.Status(400).JSON(err)
	// }
	// claims := token.Claims.(jwt.MapClaims)
	return c.Status(500).JSON(res)
	// return c.Status(500).JSON(def.Msg{
	// 	Header:  "Database Error",
	// 	Content: "user not exists",
	// })

}
