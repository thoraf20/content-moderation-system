package moderation

import (
	"fmt"
	"log"
	"strings"
)

type ImageModerationEngine struct{}

func (i *ImageModerationEngine) Moderate(content string, filename string) error {
	log.Printf("[ImageModeration] Moderating image: %s", filename)

	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	lowerName := strings.ToLower(filename)

	valid := false
	for _, ext := range allowedExtensions {
		if strings.HasSuffix(lowerName, ext) {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("image file type not allowed: %s", filename)
	}

	return nil
}