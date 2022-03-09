package screenshot

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image"
	"image/png"
	"log"
	"os"
)

func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func WorkFlow() {
	const imagePath = "./0_1920x1080.png"
	if !IsPathExist(imagePath) {
		n := screenshot.NumActiveDisplays()
		for i := 0; i < n; i++ {
			bounds := screenshot.GetDisplayBounds(i)

			img, err := screenshot.CaptureRect(bounds)
			if err != nil {
				panic(err)
			}
			fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
			file, _ := os.Create(fileName)
			defer file.Close()
			png.Encode(file, img)

			fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
		}
	} else {
		log.Println("image had existed, now to show.")
		f, err := os.Open(imagePath)
		if err != nil {
			panic(err)
		}
		img, formatName, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		fmt.Println(formatName)
		fmt.Println(img.Bounds())
		fmt.Println(img.ColorModel())
	}

}
