package view

type Block interface {
	GetType() BlockType
}
type BlockType string

const (
	BlockSection BlockType = "section"
	BlockDivider BlockType = "divider"
)

type SectionBlock struct {
	Object TextBlockObject
}

func (b SectionBlock) GetType() BlockType {
	return BlockSection
}

type DeviderBlock struct{}

func (b DeviderBlock) GetType() BlockType {
	return BlockDivider
}
