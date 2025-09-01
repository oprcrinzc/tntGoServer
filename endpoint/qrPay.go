package endpoint

import (
	"os"
	"usersys/def"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func QrPay(c *fiber.Ctx) error {
	token := c.Params("token")
	qrId := c.Params("id")
	t, err := def.VerifyJwt(token)
	if err != nil {
		c.Status(500).JSON("token error")
	}
	username := t.Claims.(jwt.MapClaims)["username"].(string)
	if username != "" {
		qrPath := os.Getenv("storage_qr_path")
		if qrPath == "" {
			log.Error("storage_qr_path is empty in .env")
			return c.SendStatus(500)
		}
		return c.SendFile(qrPath + username + "/qr/" + qrId + ".png")

	}
	return c.Status(500).JSON("error username not found")
}
