package tokens

func TokenParser(key string, value any) IToken {
	var token IToken = nil;
	switch key {
	case "$id":
		token = NewIdToken(value.(string));
		return token;
	case "type":
		token = NewTypeToken(value.(string));
		return token;
	case "properties":
		token = NewPropertiesToken(value.(map[string]any));
	}
	return token;
}