package main

import (
	"fmt"
	"os"
	"unicode"
)

func parseConfigFile(path string) map[string]string {

	configContent := readFileToString(path)

	//lexical analysis
	tokens := []configToken{}
	x := 0
	for x < len(configContent) {
		if configContent[x] == ' ' {
			//skip whitespace
			x += 1
			continue
		} else if unicode.IsLetter(rune(configContent[x])) {
			//identifier
			id := ""
			for x < len(configContent) && (unicode.IsLetter(rune(configContent[x])) || unicode.IsDigit(rune(configContent[x])) || configContent[x] == '-' || configContent[x] == '_') {
				id += string(configContent[x])
				x += 1
			}
			tokens = append(tokens, configToken{kind: "IDENTIFIER", value: id})
			continue
		} else if configContent[x] == '=' {
			//equals sign
			tokens = append(tokens, configToken{kind: "EQUALS", value: "="})
			x += 1
			continue
		} else if configContent[x] == '"' {
			//string literal

			end_quote := false

			for i := x + 1; i < len(configContent); i++ {
				if configContent[i] == '"' {
					end_quote = true
					break
				}
			}

			if !end_quote {
				fmt.Println("Error: unterminated string literal in n.config file")
				os.Exit(1)
			}

			x += 1 //skip opening quote
			strLit := ""
			for x < len(configContent) && configContent[x] != '"' {
				strLit += string(configContent[x])
				x += 1
			}
			x += 1 //skip closing quote
			tokens = append(tokens, configToken{kind: "STRING_LITERAL", value: strLit})
			continue
		} else if configContent[x] == '\n' {
			//newline
			tokens = append(tokens, configToken{kind: "NEWLINE", value: "\n"})
			x += 1
			continue
		} else {
			//unknown character
			fmt.Printf("Unknown character in n.config file: %c\n", configContent[x])
			os.Exit(1)
		}
		x += 1
	}

	//parsing
	x = 0
	configMap := make(map[string]string)

	for x < len(tokens) {
		if tokens[x].kind == "IDENTIFIER" {
			key := tokens[x].value
			x += 1
			if x < len(tokens) && tokens[x].kind == "EQUALS" {
				x += 1
				if x < len(tokens) && tokens[x].kind == "STRING_LITERAL" {
					value := tokens[x].value
					configMap[key] = value
					x += 1
					if x < len(tokens) && tokens[x].kind == "NEWLINE" {
						x += 1
						// Handle multiple newlines or empty lines
						for x < len(tokens) && (tokens[x].kind == "NEWLINE" || tokens[x].kind == "EMPTY_LINE") {
							x += 1
						}
						continue
					} else if x == len(tokens) {
						break
					} else {
						fmt.Println("Error: expected newline after string literal in n.config file")
						os.Exit(1)
					}
				} else {
					fmt.Println("Error: expected string literal after equals sign in n.config file")
					os.Exit(1)
				}
			} else {
				fmt.Println("Error: expected equals sign after identifier in n.config file")
				os.Exit(1)
			}
		} else {
			fmt.Println("Error: expected identifier at start of line in n.config file")
			os.Exit(1)
		}
	}

	//return config map
	return configMap
}

type configToken struct {
	kind  string
	value string
}

// readFileToString
func readFileToString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", path)
		os.Exit(1)
	}

	content := string(data)
	return content
}
