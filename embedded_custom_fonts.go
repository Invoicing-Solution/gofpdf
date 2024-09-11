package gofpdf

import (
	_ "embed"
)

// Embed the font files

//go:embed font/Arial.ttf
var arial []byte

//go:embed font/Arial-Bold.ttf
var arialBold []byte

//go:embed font/Arial-Italic.ttf
var arialItalic []byte

//go:embed font/Arial-BoldItalic.ttf
var arialBoldItalic []byte

//go:embed font/Arimo.ttf
var arimo []byte

//go:embed font/Arimo-Bold.ttf
var arimoBold []byte

//go:embed font/Arimo-Italic.ttf
var arimoItalic []byte

//go:embed font/Arimo-BoldItalic.ttf
var arimoBoldItalic []byte

//go:embed font/Inter.ttf
var inter []byte

//go:embed font/Inter-Bold.ttf
var interBold []byte

//go:embed font/Inter-Italic.ttf
var interItalic []byte

//go:embed font/Inter-BoldItalic.ttf
var interBoldItalic []byte
