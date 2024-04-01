package tokens

// IdToken is a token for id of the schema
type IdToken struct {
	Type_ string
	Value string
}

func NewIdToken(value string) *IdToken {
	return &IdToken{
		Type_: "id",
		Value: value,
	}
}

func (idToken *IdToken) GetType() string {
	return idToken.Type_
}

func (idToken *IdToken) GetStringValue() string {
	return idToken.Value
}

func (idToken *IdToken) GetMapValue() map[string]any {
	return nil
}