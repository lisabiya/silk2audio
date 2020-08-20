package main

import (
	"silk2audio/silk"
)

func TransSilk(inputPath string) string {
	return silk.TransSilkToWav(inputPath)
}
