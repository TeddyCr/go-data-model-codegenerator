package tokens

// PropertyFieldToken is a token for properties of the schema
type PropertiesToken struct {
	Type_ string
	Value map[string]any
}

func NewPropertiesToken(value map[string]any) *PropertiesToken {
	return &PropertiesToken{
		Type_: "properties",
		Value: value,
	}
}

func (propertyFieldToken *PropertiesToken) GetType() string {
	return propertyFieldToken.Type_;
}

func (propertyFieldToken *PropertiesToken) GetStringValue() string {
	return "";
}

func (propertyFieldToken *PropertiesToken) GetMapValue() map[string]any {
	return propertyFieldToken.Value;
}