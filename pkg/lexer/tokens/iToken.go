package tokens


// Token interface for all tokens
type IToken interface {
	GetType() string
	GetStringValue() string
	GetMapValue() map[string]any
}


