// moderation/video_engine.go
package moderation

import (
	"fmt"
	"log"
	"strings"
)

type VideoModerationEngine struct{}

func (v *VideoModerationEngine) Moderate(content string, filename string) error {
	log.Printf("[VideoModeration] Moderating video: %s", filename)

	allowedExtensions := []string{".mp4", ".mov", ".avi"}
	lowerName := strings.ToLower(filename)

	valid := false
	for _, ext := range allowedExtensions {
		if strings.HasSuffix(lowerName, ext) {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("video file type not allowed: %s", filename)
	}

	return nil
}

