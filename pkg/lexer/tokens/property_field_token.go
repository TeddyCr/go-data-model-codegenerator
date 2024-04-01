package tokens

// PropertyFieldToken is a token for properties of the schema
// revive:disable
type PropertyFieldToken struct {
	Type_ string
	Value string
	Path string
}

// NewPropertyFieldToken creates a new PropertyFieldToken
func NewPropertyFieldToken(type_ string, value string, path string) *PropertyFieldToken {
	return &PropertyFieldToken{
		Type_: type_, // "represents the data type of the field"
		Value: value,
		Path: path, // "represents the path of the field in the json structure (e.g. 'properties.name')"
	}
}

// GetType returns the type of the token
func (propertyFieldToken *PropertyFieldToken) GetType() string {
	return propertyFieldToken.Type_;
}

// GetStringValue returns the value of the token
func (propertyFieldToken *PropertyFieldToken) GetStringValue() string {
	return propertyFieldToken.Value;
}

// GetMapValue returns the value of the token
func (propertyFieldToken *PropertyFieldToken) GetMapValue() map[string]any {
	return nil;
}

// GetFieldPath returns the path of the field in the json structure
func (propertyFieldToken *PropertyFieldToken) GetFieldPath() string {
	return propertyFieldToken.Path;
}