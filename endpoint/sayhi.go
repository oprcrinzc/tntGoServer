package endpoint

import (

	// "usersys/define"

	"github.com/gofiber/fiber/v2"
	// "go.mongodb.org/mongo-driver/bson"
)

func Sayhi(c *fiber.Ctx) error {
	// client := db.Conn()
	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// coll := client.Database("tnt").Collection("user")
	// var res def.Users
	// cursor, err := coll.Find(context.TODO(), bson.D{{}})
	// if err != nil {
	// 	return c.SendString(err.Error())
	// }
	// if err = cursor.All(context.TODO(), &res); err != nil {
	// 	return c.SendString(err.Error())
	// }

	return c.SendString("Hi")
}
