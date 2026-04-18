package entity

// ValueObject
type ItemCategory int

const (
	CategoryHAT ItemCategory = iota + 1
	CategorySHIRT
	CategoryJACKET
	CategoryBOTTOMS
	CategorySHOES
)

// Item : Itemマスタ集約
type Item struct {
	ID       int
	Name     string
	Category ItemCategory
	ImageURL string
}
