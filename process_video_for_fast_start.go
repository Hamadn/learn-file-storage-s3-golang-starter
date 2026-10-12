package main

import (
	"fmt"
	"os/exec"
)

func processVideoForFastStart(filePath string) (string, error) {

	outputFilePath := filePath + ".processing"

	cmd := exec.Command("ffmpeg", "-i", filePath, "-c", "copy", "-movflags", "faststart", "-f", "mp4", outputFilePath)

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to run ffmpeg: %v", err)
	}
	return outputFilePath, nil
}
