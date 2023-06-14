package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

//create qr code using contents from txt file

func main() {
	a, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	b, err := ioutil.ReadAll(a)
	if err != nil {
		log.Fatal(err)
	}

	temp_var := string(b)
	new_qr_name := "sample.png"

	err = qrcode.WriteFile(temp_var, qrcode.Medium, 512, new_qr_name)

	if err != nil {
		log.Fatal(err)
	}
}
