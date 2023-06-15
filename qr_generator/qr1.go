package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//prompt user for qr content
	fmt.Println("Enter text to encode:")
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	//prompt user for qr file name
	fmt.Println("Enter file name (include file extension):")
	file_name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	file_name1 := strings.Replace(file_name, "\n", "", -1)

	//create qr code
	err = qrcode.WriteFile(content, qrcode.Medium, 512, file_name1)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("QR code image created.")
	}
}
