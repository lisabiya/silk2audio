package silk

/*
#cgo LDFLAGS: -L . -lSKP_SILK_SDK
#include <stdio.h>
#include <stdlib.h>
#include <signal.h>

#include "SKP_Silk_SDK_API.h"
#include "Decoder.h"
static void print_usage() {
     printf("********** Silk Decoder (Fixed Point) v %s ********************\n", SKP_Silk_SDK_get_version());
}
*/
import "C"
import (
	"fmt"
	"os"
	"silk2audio/transcoder/ffmpeg"
	"strings"
	"unsafe"
)

func TransSilkToWav(inputPath string) string {
	var outputPath = strings.Replace(inputPath, ".silk", ".pcm", -1)
	var wavPath = strings.Replace(inputPath, ".silk", ".wav", -1)

	inputPathC := C.CString(inputPath)
	outPathC := C.CString(outputPath)
	var _ = C.Decoder(inputPathC, outPathC) //调用C函数
	C.free(unsafe.Pointer(inputPathC))
	C.free(unsafe.Pointer(outPathC))
	//ffmpeg pcm转码音频
	transPcmToAudio(outputPath, wavPath)

	_ = FileRemove(outputPath)
	return wavPath
}

func transPcmToAudio(inputPath, OutputPath string) {
	format := "s16le"
	overwrite := true
	audioCodec := "pcm_s16le"
	audioChannels := 2
	audioRate := 12000
	opts := ffmpeg.Options{
		Overwrite:     &overwrite,
		OutputFormat:  &format,
		AudioChannels: &audioChannels,
		AudioRate:     &audioRate,
		AudioCodec:    &audioCodec,
	}
	ffmpegConf := &ffmpeg.Config{
		FfmpegBinPath: "silk/ffmpeg",
	}
	_, err := ffmpeg.
		New(ffmpegConf).
		Input(inputPath).
		Output(OutputPath).
		WithOptions(opts).
		Start(opts)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(OutputPath)
	}
}

func FileRemove(logFile string) error {
	_, err := os.Stat(logFile)
	if err == nil {
		return os.Remove(logFile)
	}
	return nil
}
