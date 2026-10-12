package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"os/exec"
)

func getVideoAspectRatio(filePath string) (string, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-print_format", "json", "-show_streams", filePath)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run ffprobe: %v", err)
	}

	type Stream struct {
		Height int `json:"height"`
		Width  int `json:"width"`
	}

	type FFProbeOutput struct {
		Streams []Stream `json:"streams"`
	}

	var output FFProbeOutput
	err = json.Unmarshal(out.Bytes(), &output)
	if err != nil {
		return "", fmt.Errorf("failed to parse ffprobe output: %v", err)
	}

	if len(output.Streams) == 0 {
		return "other", nil
	}

	width := output.Streams[0].Width
	height := output.Streams[0].Height

	if height == 0 || width == 0 {
		return "other", nil
	}

	ratio := float64(width) / float64(height)
	target := 9.0 / 16.0

	if math.Abs(ratio-target) < 0.001 {
		return "9:16", nil
	} else if math.Abs(ratio-(16.0/9.0)) < 0.001 {
		return "16:9", nil
	} else {
		return "other", nil
	}

}
