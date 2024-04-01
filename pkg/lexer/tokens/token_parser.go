package tokens

// TokenParser is a function that returns a token based on the key and value
// of the token.
func TokenParser(key string, value any) IToken {
	var token IToken;
	switch key {
	case "$id":
		token = NewIDToken(value.(string));
		return token;
	case "type":
		token = NewTypeToken(value.(string));
		return token;
	case "properties":
		token = NewPropertiesToken(value.(map[string]any));
	}
	return token;
}