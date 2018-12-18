package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/donatj/pngpal"
)

func init() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	path := flag.Arg(0)

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	img2, err := pngpal.ImageToPaletted(img)
	if err != nil {
		log.Fatal(err)
	}

	if img2 == img {
		os.Stderr.WriteString("image already paletted, left unmodified\n")
		os.Exit(0)
	}

	save, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	defer save.Close()

	err = png.Encode(save, img2)
	if err != nil {
		log.Fatal(err)
	}
}
