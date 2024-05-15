package main

import (
	_ "image/jpeg"
	"log"
	"os"
	"path"
	"regexp"

	"github.com/qeesung/image2ascii/convert"
)

func StripANSI(s string) string {
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)
	return re.ReplaceAllString(s, "")
}

func GenerateFrames() ([10]string, error) {
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40

	var pedroFrames [10]string

	var srcDir = "frames/pedro"
	converter := convert.NewImageConverter()

	images, err := os.ReadDir(srcDir)
	if err != nil {
		log.Fatal(err)
		return pedroFrames, err
	}
	var temp string
	for i, image := range images {
		temp = converter.ImageFile2ASCIIString(path.Join("frames", "pedro", image.Name()), &convertOptions)
		pedroFrames[i] = StripANSI(temp)
		if i == 9 {
			break
		}
	}

	return pedroFrames, nil
}
