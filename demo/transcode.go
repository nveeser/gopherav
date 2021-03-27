package main

import (
	"github.com/nveeser/gopherav"
	"log"
	"math/big"
	"os"
)

type params struct {
	outputExtension   string
	videoCodec        string
	audioCodec        string
	codecPrivateKey   string
	codecPrivateValue string
}

func main() {
	input := "/home/nicholas/demo.mp4"
	output := "/home/nicholas/output.mp4"

	// OpenInput the file

	format, err := gopherav.OpenInput(input, nil)
	if err != nil {
		log.Printf("Error: Couldn't open file: %s", err)
		os.Exit(1)
	}
	defer format.Close()

	// Decode Streams

	if err := format.InitStreamInfo(nil); err != nil {
		log.Printf("Error: Couldn't find stream information: %s\n", err)
		os.Exit(1)
	}

	decodeStream := &formatStream{
		filename: input,
		format:   format,
	}

	// Prepare Decoder
	for _, stream := range format.Streams() {
		log.Printf("MediaType: %s", stream.CodecParameters().MediaType)
		log.Printf("\tFrameRate: %s", stream.AvgFrameRate())

		cd, err := gopherav.FindCodec(stream.CodecParameters().CodecID, gopherav.Decoder)
		if err != nil {
			log.Printf("error: finding codec: %s", err)
			continue
		}
		log.Printf("\tCodec Name: %s", cd.String())
		log.Printf("\tCodec Name: %s", cd.LongName())
		codecCtx, err := stream.OpenCodecContext(gopherav.Decoder, nil)
		if err != nil {
			log.Printf("error: opening codec context: %s", err)
			continue
		}

		switch stream.CodecParameters().MediaType {
		case gopherav.MediaTypeVideo:
			decodeStream.video = &packetStream{
				stream:   stream,
				codec:    cd,
				codecCtx: codecCtx,
				index:    stream.Index(),
			}

		case gopherav.MediaTypeAudio:
			decodeStream.audio = &packetStream{
				stream:   stream,
				codec:    cd,
				codecCtx: codecCtx,
				index:    stream.Index(),
			}
		}
	}

	// Prepare Encoder

	encodeStream := &formatStream{
		filename: output,
	}

	format, err = gopherav.OpenOutput(output)

	defer format.Close()
	encodeStream.format = format

	frameRate := decodeStream.video.stream.GuessFrameRate()
	prepareVideoEncoder(encodeStream, decodeStream.video.codecCtx, frameRate, nil)

}

func prepareVideoEncoder(fs *formatStream, decoderCtx *gopherav.CodecContext, inputFramerate *big.Rat, params *params) error {
	codec, err := gopherav.FindCodecByName(params.videoCodec, gopherav.Encoder)
	if err != nil {
		return err
	}
	fs.video.stream = fs.format.NewStream(codec, params)
	//gopherav.NewCodecContext()
	return nil
}

type formatStream struct {
	filename string
	format   *gopherav.AvFormat
	video    *packetStream
	audio    *packetStream
}

type packetStream struct {
	stream   *gopherav.AvStream
	codec    *gopherav.Codec
	codecCtx *gopherav.CodecContext
	index    int
}
