package tokens

// PropertiesToken is a token for properties of the schema
// revive:disable
type PropertiesToken struct {
	Type_ string
	Value map[string]any
}

// NewPropertiesToken creates a new PropertiesToken
func NewPropertiesToken(value map[string]any) *PropertiesToken {
	return &PropertiesToken{
		Type_: "properties",
		Value: value,
	}
}

// GetType returns the type of the token
func (propertyFieldToken *PropertiesToken) GetType() string {
	return propertyFieldToken.Type_;
}

// GetStringValue returns the value of the token
func (propertyFieldToken *PropertiesToken) GetStringValue() string {
	return "";
}

// GetMapValue returns the value of the token
func (propertyFieldToken *PropertiesToken) GetMapValue() map[string]any {
	return propertyFieldToken.Value;
}