package moderation

// ModerationEngine is an interface that defines the methods for content moderation engines.
type ModerationEngine interface {
	Moderate(content string, filename string) error
}
