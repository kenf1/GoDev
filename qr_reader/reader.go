package main

import (
	"log"
	"os"

	"github.com/tuotoo/qrcode"
)

func main() {
	fi, err := os.Open("code.png")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(qrmatrix.Content)
}
