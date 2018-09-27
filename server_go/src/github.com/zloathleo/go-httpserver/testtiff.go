package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"github.com/chai2010/tiff"
	"image"
)

func main() {
	var data []byte
	var err error

	var files = []string{
		"D:/tiff/TMC.141838170.tif",
	}
	for _, filename := range files {
		// Load file data
		if data, err = ioutil.ReadFile(filename); err != nil {
			log.Fatal(err)
		}

		// Decode tiff
		m,_ := tiff.Decode(bytes.NewReader(data));
		decodeImage(m)

		//m, errors, err := tiff.DecodeAll(bytes.NewReader(data))
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//// Encode tiff
		//for i := 0; i < len(m); i++ {
		//	for j := 0; j < len(m[i]); j++ {
		//		newname := fmt.Sprintf("%s-%02d-%02d.tiff", filepath.Base(filename), i, j)
		//		if errors[i][j] != nil {
		//			log.Printf("%s: %v\n", newname, err)
		//			continue
		//		}
		//
		//		decodeImage(m[i][j])
		//
		//	}
		//}
	}
}

func decodeImage(currentImg image.Image){
	imagesBounds := currentImg.Bounds()
	imageSize := imagesBounds.Size()
	log.Println(imageSize.X ,imageSize.Y)
	for i := 0; i < imageSize.X; i++ {
		for j := 0; j < imageSize.Y; j++ {
			color := currentImg.At(i,j)
			color2 :=imagesBounds.At(i,j)

			log.Println(color)
			log.Println(color2)
		}
		}
}