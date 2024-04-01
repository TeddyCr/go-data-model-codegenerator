package tokens

// TypeToken is a token for type of the schema
// revive:disable
type TypeToken struct {
	Type_ string
	Value string
}

// NewTypeToken creates a new TypeToken
func NewTypeToken(value string) *TypeToken {
	return &TypeToken{
		Type_: "type",
		Value: value,
	}
}

// GetType returns the type of the token
func (typeToken *TypeToken) GetType() string {
	return typeToken.Type_;
}

// GetStringValue returns the value of the token
func (typeToken *TypeToken) GetStringValue() string {
	return typeToken.Value;
}

// GetMapValue returns the value of the token
func (typeToken *TypeToken) GetMapValue() map[string]any {
	return nil;
}