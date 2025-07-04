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

func CreateUser(c *fiber.Ctx) error {
	// load user input data
	body := def.UserCreationIngredient{}
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(400).SendString("Error")
	}

	checkNil := 0x0000
	if body.Name == "" {
		checkNil |= 0x0001
	} else if body.Email == "" {
		checkNil |= 0x0010
	} else if body.Password == "" {
		checkNil |= 0x0100
	}

	if checkNil != 0x0000 {
		return c.Status(400).JSON(def.Msg{
			Header:  "Ingredient Error",
			Content: strconv.Itoa(checkNil),
		})
	}

	// check is name used in db

	var res def.Users
	client := db.Conn()
	col := client.Database("tnt").Collection("user")
	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
	}()
	cursor, err := col.Find(context.TODO(), bson.D{{Key: "name", Value: body.Name}})
	if err != nil {
		return c.Status(400).JSON(err)
	}
	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	// last check

	if len(res) == 0 {
		sha512Pwd := sha512.Sum512([]byte(body.Password))
		// create user
		newUser := def.UserCreationIngredient{
			Name:     body.Name,
			Email:    body.Email,
			Password: string(fmt.Sprintf("%x", sha512Pwd)),
		}
		createUserRes, err := col.InsertOne(context.TODO(), newUser)
		if err != nil {
			return c.Status(500).JSON(def.Msg{
				Header:  "Database Error",
				Content: err.Error(),
			})
		}
		return c.Status(200).JSON(def.Msg{
			Header:  "Create user",
			Content: strconv.FormatBool(createUserRes.Acknowledged),
		})
	}

	return c.Status(400).JSON(def.Msg{
		Header:  "Database Error",
		Content: "username is already taken",
	})
}
