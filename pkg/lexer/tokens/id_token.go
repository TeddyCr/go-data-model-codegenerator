package tokens

// IDToken is a token for id of the schema
// revive:disable
type IDToken struct {
	Type_ string
	Value string
}

// NewIDToken creates a new IDToken
func NewIDToken(value string) *IDToken {
	return &IDToken{
		Type_: "id",
		Value: value,
	}
}

// GetType returns the type of the token
func (idToken *IDToken) GetType() string {
	return idToken.Type_
}

// GetStringValue returns the value of the token
func (idToken *IDToken) GetStringValue() string {
	return idToken.Value
}

// GetMapValue returns the value of the token
func (idToken *IDToken) GetMapValue() map[string]any {
	return nil
}