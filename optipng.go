package optipng

import (
	"bytes"
	"image"
	"image/png"
	"os/exec"
	"strings"
)

func Compress(input image.Image, args []string) (output image.Image, err error) {
	var w bytes.Buffer
	err = png.Encode(&w, input)
	if err != nil {
		return
	}

	b := w.Bytes()
	compressed, err := CompressBytes(b, args)
	if err != nil {
		return
	}

	output, err = png.Decode(bytes.NewReader(compressed))
	return
}

func CompressBytes(input []byte, args []string) (output []byte, err error) {
	// "-", "--speed", speed
	cmd := exec.Command("optipng", args...)
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o
	err = cmd.Run()

	if err != nil {
		return
	}

	output = o.Bytes()
	return
}