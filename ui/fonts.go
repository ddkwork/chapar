package ui

import (
	"embed"
	"fmt"
	"gioui.org/font"
	"gioui.org/font/opentype"
)

//go:embed fonts/*
var fonts embed.FS

func prepareFonts() ([]font.FontFace, error) {
	var fontFaces []font.FontFace
	robotoRegularTTF, err := getFont("Roboto-Regular.ttf")
	if err != nil {
		return nil, err
	}

	robotoRegular, err := opentype.Parse(robotoRegularTTF)
	if err != nil {
		return nil, err
	}
	fontFaces = append(fontFaces, font.FontFace{Font: font.Font{}, Face: robotoRegular})
	return fontFaces, nil
}

func getFont(path string) ([]byte, error) {
	data, err := fonts.ReadFile(fmt.Sprintf("fonts/%s", path))
	if err != nil {
		return nil, err
	}

	return data, err
}