package def

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/sigurn/crc16"
	"github.com/skip2/go-qrcode"
)

func GenQrString(amt float64) (string, error) {
	amtString := strconv.FormatFloat(amt, 'f', 2, 32)
	if amt < 10 {
		amtString = "0" + amtString
	}
	lenAmtString := strconv.Itoa(len(amtString))
	if len(amtString) < 10 {
		lenAmtString = "0" + lenAmtString
	}
	firstSeg := os.Getenv("first_seg_qr_string")
	if firstSeg == "" {
		return "", errors.New("first_seg_qr_string is empty in .env")
	}
	qrString := firstSeg + lenAmtString + amtString + "6304"
	table := crc16.MakeTable(crc16.Params{Poly: 0x1021, Init: 0xFFFF, RefIn: false, RefOut: false, XorOut: 0x0000, Check: 0x31C3, Name: "CRC-16/XMODEM"})
	crc := crc16.Checksum([]byte(qrString), table)
	qrString += fmt.Sprintf("%X", crc)
	return qrString, nil
}

func GenQr(qrString string, fileName string) error {
	qrPath := os.Getenv("storage_qr_path")
	if qrPath == "" {
		// log.Error("storage_qr_path is empty in .env")
		return errors.New("storage_qr_path is empty in .env")
	}
	if fileName == "" {
		return errors.New("fileName is empty")
	}
	if qrString == "" {
		return errors.New("qrString is empty")
	}
	err := qrcode.WriteFile(qrString, qrcode.Highest, 256, qrPath+fileName)
	if err != nil {
		panic(err)
	}
	return nil
}
