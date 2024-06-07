package libVideoUtils

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gogf/gf/v2/os/gfile"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io"
	"os"
)

func ProcessVideo(srcPath, destDir string) (videoPath, coverPath string, err error) {
	fileName := gfile.Basename(srcPath)
	introFileName := "intro_" + fileName
	videoPath = destDir + "/" + introFileName
	err = ffmpeg.Input(srcPath, ffmpeg.KwArgs{"ss": 1}).
		Output(videoPath, ffmpeg.KwArgs{"t": 15, "vcodec": "libx264"}).OverWriteOutput().Run()

	//err = ffmpeg.Input(videoPath).
	//	Output(videoPath, ffmpeg.KwArgs{"c:v": "libx265"}).
	//	OverWriteOutput().ErrorToStdOut().Run()
	reader := ExampleReadFrameAsJpeg(srcPath, 1)
	img, err := imaging.Decode(reader)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	name := gfile.Basename(srcPath)

	coverFileName := "cover_" + name + ".png"
	coverPath = destDir + "/" + coverFileName
	err = imaging.Save(img, coverPath)
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}

	return
}

func in(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}
