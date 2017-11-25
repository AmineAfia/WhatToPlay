package qrcode

import (
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

const folder = "qrs/"

func CreateQr(baseurl string, roomName string) {
	url := baseurl + "room/" + roomName
	mode := int(0755)
	_ = os.Mkdir(folder, os.FileMode(mode))
	err := qrcode.WriteFile(url, qrcode.Medium, 256, folder+roomName+".png")

	if err != nil {
		panic(err)
	}
}

func DeleteQr(roomName string) {
	path := folder + roomName + ".png"
	err := os.Remove(path)

	if err != nil {
	}
}
