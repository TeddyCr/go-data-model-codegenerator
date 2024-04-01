package lexer

import (
	"os"
	"path"
	"testing"

	"github.com/TeddyCr/go-data-model-codegenerator/pkg/lexer/tokens"
	"github.com/stretchr/testify/assert"
)

func TestInvalidLexerTokenizer(t *testing.T) {
	dir, _ := os.Getwd()
	dirPath := path.Join(dir, "testdata", "invalids")
	files, _ := os.ReadDir(dirPath)

	for _, file := range files {
		filePath := path.Join(dirPath, file.Name())
		lexer := NewLexer(filePath)
		_, err := lexer.Tokenize()
		assert.NotNil(t, err)
	}
}

func TestValidLexerTokenizer(t *testing.T) {

	expected := map[string]tokens.IToken{
		"id":   tokens.NewIdToken("https://open-metadata.org/schema/tests/testSuite.json"),
		"type": tokens.NewTypeToken("object"),
		"properties": tokens.NewPropertiesToken(
			map[string]interface{}{
				"id": map[string]interface{}{
					"description": "Unique identifier of this test suite instance",
					"type":        "string",
				},
				"name": map[string]interface{}{
					"description": "Name that identifies this test suite",
					"type":        "string",
				},
			},
		),
	}

	dir, _ := os.Getwd()
	dirPath := path.Join(dir, "testdata", "valids")
	files, _ := os.ReadDir(dirPath)

	for _, file := range files {
		filePath := path.Join(dirPath, file.Name())
		lexer := NewLexer(filePath)
		tokens, _ := lexer.Tokenize()
		assert.Equal(t, 3, len(tokens))
		validateTokens(t, tokens, expected)
	}
}

func validateTokens(t *testing.T, actual []tokens.IToken, expected map[string]tokens.IToken) {
	for _, element := range actual {
		// revive:disable
		type_ := element.GetType()
		expectedToken := expected[type_]
		assert.NotNil(t, expectedToken)

		if element.GetStringValue() != "" {
			assert.Equal(t, expectedToken.GetStringValue(), element.GetStringValue())
		} else if element.GetMapValue() != nil {
			// TODO: Implement a more robust way to compare maps
			assert.Equal(t, 2, len(element.GetMapValue()))
		} else {
			assert.Fail(t, "Neither string nor map value found in token")
		}
	}
}
