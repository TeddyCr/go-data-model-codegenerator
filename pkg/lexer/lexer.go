// Package lexer provides the lexer for the json-schema file
// It reads the json file and returns a list of tokens
package lexer

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/TeddyCr/go-data-model-codegenerator/pkg/lexer/tokens"
)

// A Lexer reads a json file and returns a list of tokens
type Lexer struct {
	File *os.File
}

// NewLexer is a creational method instantiating a `Lexer` object
// It takes a file path as an argument and returns a pointer to a `Lexer` object
func NewLexer(filePath string) *Lexer {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Lexer{
		File: jsonFile,
	}
}

func (lexer *Lexer) serialize() (map[string]interface{}, error) {
	var data map[string]interface{}
	jsonBytes, err := io.ReadAll(lexer.File)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Tokenize is a method of the `Lexer` object that reads the json file and returns a list of tokens
// It returns a list of tokens and an error if any occurs (i.e. file not found, invalid json)
func (lexer *Lexer) Tokenize() ([]tokens.IToken, error) {
	// lexer.validate()
	serializedJSON, err := lexer.serialize()
	if err != nil {
		return nil, err
	}
	var tokenList []tokens.IToken
	for key, value := range serializedJSON {
		token := tokens.TokenParser(key, value)
		if token != nil {
			tokenList = append(tokenList, token)
		}
	}
	lexer.close()
	return tokenList, nil
}

func (lexer *Lexer) close() {
	lexer.File.Close()
}
