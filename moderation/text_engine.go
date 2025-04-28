package moderation

import (
	"fmt"
	"log"
	"strings"
)

type TextModerationEngine struct{}

func (t *TextModerationEngine) Moderate(content string, filename string) error {
	log.Printf("[TextModeration] Moderating text: %s", content)

	bannedWords := []string{"badword", "hate", "spam", "violence", "abuse", "harassment", "bullying", "racism", "sexism", "discrimination"}
	maxLength := 5000 

	contentLower := strings.ToLower(content)
	for _, word := range bannedWords {
		if strings.Contains(contentLower, word) {
			return fmt.Errorf("text contains banned word: %s", word)
		}
	}

	if len(content) == 0 {
		return fmt.Errorf("text content is empty")
	}
	
	if len(content) > maxLength {
		return fmt.Errorf("text content exceeds maximum length of %d characters", maxLength)
	}

	return nil
}
