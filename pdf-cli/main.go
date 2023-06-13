package main

import (
	"fmt"
	"os"

	"github.com/sunshineplan/imgconv"
)

func main() {
	f, err := os.Open("./Files/sample.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	img, err := imgconv.Decode(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = imgconv.Write(os.Stdout, img, &imgconv.FormatOption{Format: imgconv.JPEG})
	if err != nil {
		fmt.Println(err)
		return
	}
}
