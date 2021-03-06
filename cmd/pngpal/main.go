package main

import (
	"flag"
	"image/png"
	"log"
	"os"

	"github.com/donatj/pngpal"
)

var (
	force  = flag.Bool("f", false, "force re-encode of existing paletted images")
	output = flag.String("o", "", "output path - if not specified, passed file will be replaced")
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
		log.Fatal(err, " - ", path)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		log.Fatal(err, " - ", path)
	}
	file.Close()

	img2, err := pngpal.ImageToPaletted(img)
	if err != nil {
		log.Fatal(err, " - ", path)
	}

	if !*force && img2 == img {
		os.Stderr.WriteString("image already paletted, left unmodified\n")
		os.Exit(0)
	}

	opath := path
	if *output != "" {
		opath = *output
	}

	save, err := os.Create(opath)
	if err != nil {
		log.Fatal(err, " - ", opath)
	}
	defer save.Close()

	err = png.Encode(save, img2)
	if err != nil {
		log.Fatal(err, " - ", opath)
	}
}
