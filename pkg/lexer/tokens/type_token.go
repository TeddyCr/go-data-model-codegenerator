package tokens

// TypeToken is a token for type of the schema
type TypeToken struct {
	Type_ string
	Value string
}

func NewTypeToken(value string) *TypeToken {
	return &TypeToken{
		Type_: "type",
		Value: value,
	}
}

func (typeToken *TypeToken) GetType() string {
	return typeToken.Type_;
}

func (typeToken *TypeToken) GetStringValue() string {
	return typeToken.Value;
}

func (typeToken *TypeToken) GetMapValue() map[string]any {
	return nil;
}