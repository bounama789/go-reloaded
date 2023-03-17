package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cont = new(string)

func main() {

	args := os.Args

	if len(args) == 3 {
		filename := args[1]
		outfilename := args[2]
		content, _ := os.ReadFile(filename)
		*cont = string(content)
		var output string
		fmt.Println(*cont)

		file, _ := os.Create(outfilename)
		erro := os.Truncate(outfilename, 0)

		if erro != nil {
			panic(erro)
		}

		output += process(*cont, output, 0)
		output = puncCheck(output)
		output = fixQuotes(output)
		output = grammarCheck(output)
		if output[len(output)-1] == ' ' {
			output = RemoveStringElem(output, len(output)-1)
		}

		_, err := file.WriteString(output)
		if err != nil {
			panic(err)
		}

	}

}

func getOption(s string, begin int) (string, string, int, int) {
	options := []string{"cap", "up", "hex", "bin", "low"}

	var a, b, nword int
	var err error
	var text, option string
	for i := begin; i < len(s); i++ {
		nword = 0
		if s[i] == '(' {
			a = i
			text = s[begin : a-1]
		}
		if s[i] == ')' {
			b = i
			option = s[a+1 : b]
			if strings.Contains(option, ",") {
				commaIndex := strings.Index(option, ",")
				nword, err = strconv.Atoi(option[commaIndex+2:])
				if err != nil {
					continue
				}
				if nword <= 0 {
					fmt.Printf("Error at (%s): number of words must be greater than 0 \n", option)
					os.Exit(1)
				}

				option = option[:commaIndex]

			}
			if !isIn(options, option) {
				continue
			}
			break
		}
	}
	return text, option, b, nword
}

func getIndexRange(s string, nword int) []int {
	words := strings.Split(s, " ")
	var beg, end int
	var w string
	var idx []int

	if nword <= 1 {
		w = words[len(words)-1]
		beg = strings.LastIndex(s, w)
	} else {
		x := strings.LastIndex(s, words[nword-1])
		print(x)
		w = words[len(words)-(nword)]
		beg = strings.LastIndex(s[:(len(s)-1)-len(s[x:])], w)
	}
	end = len(s) - 1
	idx = []int{beg, end}
	return idx
}

func process(content string, output string, start int) string {
	text, option, last, nword := getOption(content, start)

	if option == "" {
		return output + content[start:]
	}

	idx := getIndexRange(text, nword)
	output += correct(text, option, idx)

	return process(content, output, last+1)
}

func correct(text string, option string, idx []int) string {
	beg := idx[0]
	end := idx[1]
	word := text[beg : end+1]

	switch option {
	case "low":
		text = text[:beg] + strings.ToLower(word) + text[end+1:]
	case "up":
		text = text[:beg] + strings.ToUpper(word) + text[end+1:]
	case "cap":
		text = text[:beg] + strings.Title(word) + text[end+1:]
	case "bin":
		w, _ := strconv.ParseInt(word, 2, 8)
		s := strconv.FormatInt(w, 10)
		text = text[:beg] + s + text[end+1:]
	case "hex":
		w, _ := strconv.ParseInt(word, 16, 8)
		s := strconv.FormatInt(w, 10)
		text = text[:beg] + s + text[end+1:]
	}
	return text
}

func isValidPunct(content string) (bool, int, int) {
	for i, v := range content {
		var next, last byte
		if i > 0 && i < len(content)-1 {

			next = content[i+1]
			last = content[i-1]

		}

		if isPunctuation(v) && last == ' ' {
			return false, i - 1, 0
		}

		if isPunctuation(v) && next != ' ' && !isPunctuation(rune(next)) {
			return false, i, 1
		}
	}
	return true, -1, -1
}

func isPunctuation(val rune) bool {
	punctuation := []rune{'.', ',', '!', '?', ':', ';'}

	for _, v := range punctuation {
		if v == val {
			return true
		}
	}
	return false
}

func isAlpha(s string) bool {
	for i := range s {
		if s[i] < 'a' && s[i] > 'z' || s[i] < 'A' && s[i] > 'Z' || s[i] < '0' && s[i] > '9' {
			return false
		}
	}
	return true
}

func puncCheck(content string) string {
	valid, idxError, act := isValidPunct(content)

	if valid {
		return content
	}

	switch act {
	case 0:
		content = RemoveStringElem(content, idxError)
	case 1:
		content = addSpaceBetweenString(content, idxError)
	}
	return puncCheck(content)
}

func RemoveStringElem(s string, index int) string {
	s = s[:index] + s[index+1:]
	return s
}

func addSpaceBetweenString(s string, index int) string {
	s = s[:index+1] + " " + s[index+1:]
	return s
}

func quoteCheck(content string) string {
	isOpenQuote := true
	for i, v := range content {
		if i > 0 && i < len(content)-1 {
			last := content[i-1]
			next := content[i+1]
			isQuote := v == '\'' || v == '"'

			if isQuote && isAlpha(string(next)) && isAlpha(string(last)) {
				continue
			}
			if isQuote {
				if last != ' ' && isPunctuation(rune(content[i-1])) && isOpenQuote {
					content = addSpaceBetweenString(content, i-1)
					i++
				}
				if isOpenQuote && next == ' ' {
					content = RemoveStringElem(content, i+1)
				}
				isOpenQuote = !isOpenQuote
			}
		}

	}
	return content
}

func grammarCheck(content string) string {
	words := strings.Split(content, " ")
	chars := []string{"a", "e", "i", "o", "u", "a", "h"}

	for i, v := range words {
		v = strings.ToLower(v)
		if v == "a" && isIn(chars, strings.ToLower(string(words[i+1][0]))) {
			words[i] += "n"
		}
	}
	return strings.Join(words, " ")
}

func isIn(arr []string, val string) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func fixQuotes(str string) string {
	// Define variables to keep track of whether we're inside quotes and which kind of quote we're using
	inQuotes := false
	quoteType := ""

	// Create a new string builder to store the corrected string
	var fixedStr strings.Builder

	// Iterate over each character in the input string
	for _, char := range str {
		// Check if the character is a quote
		if char == '\'' || char == '"' {
			// If we're not currently inside quotes, start a new quote
			if !inQuotes {
				inQuotes = true
				quoteType = string(char)
				fixedStr.WriteRune(char)
			} else {
				// If we're already inside quotes, end the current quote
				if quoteType == string(char) {
					inQuotes = false
					quoteType = ""
					fixedStr.WriteRune(char)
				} else {
					// If we encounter a different type of quote inside the current quote, treat it as a regular character
					fixedStr.WriteRune(char)
				}
			}
		} else {
			// If the character is not a quote, check if we're currently inside quotes
			if inQuotes {
				// If we're inside quotes, remove any spaces before or after the current character
				if char != ' ' {
					fixedStr.WriteRune(char)
				}
			} else {
				// If we're not inside quotes, just add the current character to the output string
				fixedStr.WriteRune(char)
			}
		}
	}

	// Return the corrected string
	return fixedStr.String()
}
