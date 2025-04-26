// moderation/dispatcher.go
package moderation

import "fmt"

func GetModerationEngine(contentType string) (ModerationEngine, error) {
	switch contentType {
	case "text":
		return &TextModerationEngine{}, nil
	case "image":
		return &ImageModerationEngine{}, nil
	case "video":
		return &VideoModerationEngine{}, nil
	default:
		return nil, fmt.Errorf("unsupported content type: %s", contentType)
	}
}
