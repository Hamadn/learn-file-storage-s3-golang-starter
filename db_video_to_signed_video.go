package main

import (
	"errors"
	"github.com/bootdotdev/learn-file-storage-s3-golang-starter/internal/database"
	"strings"
	"time"
)

func (cfg *apiConfig) dbVideoToSignedVideo(video database.Video) (database.Video, error) {

	if video.VideoURL == nil {
		return video, nil
	}

	if !strings.Contains(*video.VideoURL, ",") {
		return video, errors.New("video URL does not contain a comma")
	}

	parts := strings.SplitN(*video.VideoURL, ",", 2)
	bucket, key := parts[0], parts[1]

	presignedURL, err := generatePresignedURL(cfg.s3Client, bucket, key, time.Minute*15)
	if err != nil {
		return video, err
	}
	video.VideoURL = &presignedURL
	return video, nil
}
