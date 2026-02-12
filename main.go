package main

import (
	"HyprWallsChanger-go/config"
	"HyprWallsChanger-go/image_handler"
	"fmt"
	"log"
	"os"
	"slices"

	// "encoding/json"

	"github.com/anthonynsimon/bild/imgio"
)

/*

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	jsonData := []byte(`{"item": "laptop", "price": 999.99}`)

	var data map[string]interface{}

	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Item: %v, Price: %v\n", data["item"], data["price"])
	// Output: Item: laptop, Price: 999.99
}


*/

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("[!] Invalid argument. Run program like: hypr-walls-blur <image>")
	}

	var filename string = os.Args[1]
	var format string = filename[len(filename)-3:]
	var blur_radius float64 = 100.0
	var save_flag bool = false
	var paths_to_img map[string]string = config.LoadImagePath("")

	paths_to_img["wall_dir"] = config.FormatPath(paths_to_img["wall_dir"])
	paths_to_img["lock_dir"] = config.FormatPath(paths_to_img["lock_dir"])

	if image_handler.ImageIsExists(paths_to_img["lock_dir"] + "/" + filename){
		fmt.Println("[*] The image already exists")
		return
	}

	img, err := imgio.Open( paths_to_img["wall_dir"] + "/" + filename )
	if err != nil {
		log.Fatalf("[!] failed to load image: %v", err)
	}

	fmt.Println("[*] Start blurring...")
	blurred_img := image_handler.GaussianBlur(img, blur_radius)

	fmt.Println("[*] Start save image...")
	if slices.Contains([]string{"jpg", "jpeg"}, format) {
		save_flag = image_handler.SaveJPEG(paths_to_img["lock_dir"] + "/" + filename, blurred_img)
	} else if format == "png" {
		save_flag = image_handler.SavePNG(paths_to_img["lock_dir"] + "/" + filename, blurred_img)
	} else {
		log.Fatalf("[!] Unsupported format \"%s\"", format)
	}
	if !save_flag {
		log.Fatalf("[!] The image is not saved")
	}

	fmt.Println("[+] Image blurred and saved as blur.png")
}
