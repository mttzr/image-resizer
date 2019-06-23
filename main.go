package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"log"
	"github.com/disintegration/imaging"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter image file location: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")

	src, err := imaging.Open(text)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	new_img_resize := imaging.Resize(src, 500, 500, imaging.Lanczos)

	reader2 := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file location for saving output: ")
	out, _ := reader2.ReadString('\n')
	out = strings.TrimSuffix(out, "\n")
	full_out := out + "/square_image.jpg"

	err = imaging.Save(new_img_resize, full_out)
	if err != nil {
		log.Fatalf("Failed to save image: %v", err)
	}

}