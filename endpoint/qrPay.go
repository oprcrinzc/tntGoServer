package endpoint

import "github.com/gofiber/fiber/v2"

func QrPay(c *fiber.Ctx) error {
	return c.SendFile("/mnt/game/projects/tnt3dPrint/UserSys/storage/qr.png")
}
