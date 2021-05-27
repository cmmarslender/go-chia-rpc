package types

// PuzzleHash Own type for future methods to encode/decode
type PuzzleHash string

// SerializedProgram Just represent as a string for now
type SerializedProgram string

// ClassgroupElement Classgroup Element
type ClassgroupElement struct {
	Data string `json:"data"`
}
