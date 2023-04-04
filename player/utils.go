package player

import (
	"bufio"
	"encoding/binary"
	"io"
	"os/exec"
	"strconv"

	"github.com/Goscord/goscord/goscord/gateway"
	"layeh.com/gopus"
)

const (
	channels  int = 2                   // 1 for mono, 2 for stereo
	frameRate int = 48000               // audio sampling rate
	frameSize int = 960                 // uint16 size of each audio frame
	maxBytes      = (frameSize * 2) * 2 // max size of opus data
)

var opusEncoder *gopus.Encoder

func PlayUrlOrFile(v *gateway.VoiceConnection, filename string, stop <-chan bool) {
	run := exec.Command("ffmpeg", "-i", filename, "-f", "s16le", "-ar", strconv.Itoa(frameRate), "-ac", strconv.Itoa(channels), "pipe:1")
	ffmpegout, err := run.StdoutPipe()
	if err != nil {
		return
	}

	ffmpegbuf := bufio.NewReaderSize(ffmpegout, 16384)

	err = run.Start()
	if err != nil {
		return
	}

	defer run.Process.Kill()

	go func() {
		<-stop
		err = run.Process.Kill()
	}()

	v.Speaking(true)

	defer func() {
		v.Speaking(false)
	}()

	opusEncoder, err = gopus.NewEncoder(frameRate, channels, gopus.Audio)

	if err != nil {
		return
	}

	for {
		audiobuf := make([]int16, frameSize*channels)
		err = binary.Read(ffmpegbuf, binary.LittleEndian, &audiobuf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return
		}
		if err != nil {
			return
		}

		select {
		default:
			opus, err := opusEncoder.Encode(audiobuf, frameSize, maxBytes)
			if err != nil {
				return
			}

			if !v.Ready() {
				return
			}

			v.Write(opus)
		}
	}
}
