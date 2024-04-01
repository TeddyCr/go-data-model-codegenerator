// Package tokens interface for all tokens
// defines the methods that all tokens should implement
package tokens


// IToken interface for all tokens
type IToken interface {
	GetType() string
	GetStringValue() string
	GetMapValue() map[string]any
}


