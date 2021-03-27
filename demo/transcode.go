package main

import (
	"github.com/nveeser/gopherav"
	"log"
	"os"
)

func main() {
	filename := "/home/nicholas/demo.mp4"

	format, err := gopherav.Open(filename, nil)
	if err != nil {
		log.Printf("Error: Couldn't open file: %s", err)
		os.Exit(1)
	}
	defer format.Close()

	if err := format.InitStreamInfo(nil); err != nil {
		log.Printf("Error: Couldn't find stream information: %s\n", err)
		os.Exit(1)
	}

	for _, stream := range format.Streams() {
		log.Printf("MediaType: %s", stream.CodecParameters().MediaType())
		log.Printf("\tFrameRate: %s", stream.AvgFrameRate())
		cd, err := gopherav.FindCodec(stream.CodecParameters().CodecId(), gopherav.Decoder)
		if err != nil {
			log.Printf("error: finding codec: %s", err)
			continue
		}
		log.Printf("\tCodec Name: %s", cd.String())
		log.Printf("\tCodec Name: %s", cd.LongName())
		decoder, err := stream.OpenCodecContext(gopherav.Decoder, nil)
		if err != nil {
			log.Printf("error: opening codec context: %s", err)
			continue
		}
		decoder.Close()
	}
}
