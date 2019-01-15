package service

import (
	"github.com/godcong/go-ffmpeg/ffmpeg"
	"github.com/godcong/go-ffmpeg/ffprobe"
	"log"
	"os"
)

// ToM3U8WithKey ...
func ToM3U8WithKey(id string) error {

	output := config.Media.Transfer + "/" + id
	source := config.Media.Upload + "/" + id
	probe := ffprobe.New(source)

	//err := rdsQueue.Set(id, StatusTransferring, 0).Err()
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}

	procFunc := ffmpeg.SplitWithKey
	if probe.Run().IsH264AndAAC() {
		procFunc = ffmpeg.QuickSplitWithKey
	}

	b, err := procFunc(source, output, config.Media.KeyInfoFile, "media", config.Media.M3U8)
	if err != nil {
		log.Println(string(b), err)
		return err
	}

	return nil
}

// ToM3U8 ...
func ToM3U8(id string) error {

	output := config.Media.Transfer + "/" + id
	_ = os.MkdirAll(output, os.ModePerm) //ignore err

	source := config.Media.Upload + "/" + id
	probe := ffprobe.New(source)

	//err := rdsQueue.Set(id, StatusTransferring, 0).Err()
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}

	procFunc := ffmpeg.Split
	if probe.Run().IsH264AndAAC() {
		procFunc = ffmpeg.QuickSplit
	}

	b, err := procFunc(source, output, "media", config.Media.M3U8)
	if err != nil {
		log.Println(string(b), err)
		return err
	}

	return nil
}
