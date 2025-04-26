package moderation

import (
	"strings"
)

var bannedWords = []string{
	"badword", "inappropriate", "hate", "violence", "badword1", "badword2", "offensive",// Add more as needed,
}

type ModerationResult struct {
	Flagged  bool
	Reason   string
	Offenses []string
}

func ModerateText(content, id string) (ModerationResult, error) {
	offenses := []string{}

	for _, word := range bannedWords {
		if strings.Contains(strings.ToLower(content), word) {
			offenses = append(offenses, word)
		}
	}

	if len(offenses) > 0 {
		return ModerationResult{
			Flagged:  true,
			Reason:   "offensive content detected",
			Offenses: offenses,
		}, nil
	}

	return ModerationResult{Flagged: false, Reason: "clean"}, nil
}
