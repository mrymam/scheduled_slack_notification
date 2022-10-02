package view

type TextBlockType string

const (
	TextBlockMarkdown TextBlockType = "markdown"
)

type TextBlockObject struct {
	TextType TextBlockType
	Text     string
}
