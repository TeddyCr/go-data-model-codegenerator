package tokens

// PropertyFieldToken is a token for properties of the schema
type PropertyFieldToken struct {
	Type_ string
	Value string
	Path string
}

func NewPropertyFieldToken(type_ string, value string, path string) *PropertyFieldToken {
	return &PropertyFieldToken{
		Type_: type_, // "represents the data type of the field"
		Value: value,
		Path: path, // "represents the path of the field in the json structure (e.g. 'properties.name')"
	}
}

func (propertyFieldToken *PropertyFieldToken) GetType() string {
	return propertyFieldToken.Type_;
}

func (propertyFieldToken *PropertyFieldToken) GetStringValue() string {
	return propertyFieldToken.Value;
}

func (propertyFieldToken *PropertyFieldToken) GetMapValue() map[string]any {
	return nil;
}

func (propertyFieldToken *PropertyFieldToken) GetFieldPath() string {
	return propertyFieldToken.Path;
}