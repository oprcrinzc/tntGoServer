package endpoint

import (
	"context"
	"os"
	"strconv"
	"time"
	"usersys/db"
	"usersys/def"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
)

func Ext(full string) string {
	i := len(full) - 1
	extR := ""
	ext := ""
	for j := i; j <= i; {
		if full[j] == '.' {
			break
		}
		extR += string(full[j])
		j -= 1
	}
	for j := len(extR) - 1; j <= len(extR)-1; {
		ext += string(extR[j])
		j -= 1
		if j < 0 {
			break
		}
	}
	return "." + ext
}

func Order(c *fiber.Ctx) error {
	order := def.Order{}
	cnt := c.FormValue("Content")
	color := c.FormValue("Color")
	material := c.FormValue("Material")
	multiP, err := c.FormFile("File")
	if err != nil {
		log.Error(err)
	}

	log.Info(Ext(multiP.Filename))

	log.Info(cnt)
	if cnt == "" {
		return c.Status(200).JSON("NO")
	}
	// err := c.BodyParser(&order)
	token := c.Get("Authorization")
	// log.Info(token)\
	fileName := ""
	orderTime := int(time.Now().UnixNano())
	order.Content = cnt
	order.Time = strconv.Itoa(orderTime)
	order.Color = color
	order.Material = material
	order.Status = def.StatusPending
	Ack := false

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
		fileName = order.Customer + strconv.Itoa(orderTime) + Ext(multiP.Filename)
		order.File = []string{fileName}

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
			// return c.JSON(iRes.Acknowledged)
			Ack = iRes.Acknowledged
		}

	}

	if err != nil && Ack {
		return c.JSON(err.Error())
	}

	storagePath := os.Getenv("storage_user_path")
	if storagePath == "" {
		log.Error("storage_user_path is empth in .env")
	}

	err = os.Mkdir(storagePath+order.Customer, os.ModeDir)
	if !os.IsExist(err) {
		os.Chmod(storagePath+order.Customer, os.ModePerm)
	}
	err = os.Mkdir(storagePath+order.Customer+"/qr/", os.ModeDir)
	if !os.IsExist(err) {
		os.Chmod(storagePath+order.Customer+"/qr/", os.ModePerm)
	}

	// log.Info(os.IsExist(err))
	err = c.SaveFile(multiP, storagePath+order.Customer+"/"+fileName)
	if err != nil {
		return c.JSON(err)
	}

	qrString, err := def.GenQrString(20)
	if err != nil {
		log.Error("qrString error", err)
		return c.SendStatus(500)
	}
	def.GenQr(qrString, order.Customer+"/qr/"+order.Time+".png")

	return c.JSON(order)
}
