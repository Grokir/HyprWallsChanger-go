package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"HyprWallsChanger-go/image_handler"
	"github.com/anthonynsimon/bild/imgio"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[!] Invalid argument. Run program like: hypr-walls-blur <image>")
	}

	var filename string = os.Args[1]
	var format string = filename[len(filename)-3:]
	var blur_radius float64 = 100.0
	var save_flag bool = false

	img, err := imgio.Open(filename)
	if err != nil {
		log.Fatalf("[!] failed to load image: %v", err)
	}

	fmt.Println("[*] Start blurring...")
	blurred_img := image_handler.GaussianBlur(img, blur_radius)

	fmt.Println("[*] Start save image...")
	if slices.Contains([]string{"jpg", "jpeg"}, format) {
		save_flag = image_handler.SaveJPEG(filename, blurred_img)
	} else if format == "png" {
		save_flag = image_handler.SavePNG(filename, blurred_img)
	} else {
		log.Fatalf("[!] Unsupported format \"%s\"", format)
	}
	if !save_flag {
		log.Fatalf("[!] The image is not saved")
	}

	fmt.Println("[+] Image blurred and saved as blur.png")
}
