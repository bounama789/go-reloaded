package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cont = new(string)
var results = new(string)

func main() {

	args := os.Args

	if len(args) == 3 {
		filename := args[1]
		outfilename := args[2]
		content, _ := os.ReadFile(filename)
		*cont = string(content)
		var output string

		file, _ := os.Create(outfilename)
		erro := os.Truncate(outfilename, 0)

		if erro != nil {
			panic(erro)
		}
		output = process(*cont)
		_, err := file.WriteString(output)
		if err != nil {
			panic(err)
		}

	} else {
		fmt.Println("USAGE: go run . <input file> <output file>")
	}
}

func getOption(s string, begin int) (string, string, int, int) {
	options := []string{"cap", "up", "hex", "bin", "low"}

	var a, b, nword int
	var err error
	var text, option, temp string
	for i := begin; i < len(s); i++ {
		nword = 0
		if s[i] == '(' {
			a = i
			if a > begin {
				text = s[begin : a-1]
				if len(text) == 0 {
					continue
				}

			}
		}
		if s[i] == ')' {
			b = i
			temp = s[a+1 : b]
			if strings.Contains(temp, ",") {
				commaIndex := strings.Index(temp, ",")
				nword, err = strconv.Atoi(temp[commaIndex+2:])
				if err != nil {
					continue
				}
				option = strings.ToLower(temp[:commaIndex])
				if option == options[2] || option == options[3] {
					option = ""
					continue
				}
				if !isIn(options, option) {
					continue
				}

				if nword <= 0 {
					fmt.Printf("Error at (%s): number of words must be greater than 0 \n", option)
					os.Exit(1)
				}

				*results += text

			} else {
				option = strings.ToLower(s[a+1 : b])
				if !isIn(options, option) {
					text = s[begin : b+1]
					continue
				}
				*results += text
			}
			break
		}
	}
	if !isIn(options, option) {
		return *results, "", b, nword

	}
	return *results, option, b, nword
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
		for i := len(s) - 1; i >= 0; i-- {
			if nword == 0 {
				beg = i + 1
				break
			}
			if s[i] == ' ' {
				nword--
			}
		}
		//		w = words[len(words)-(nword)]
		//beg = strings.LastIndex(s[:(len(s)-1)-len(s[x:])], w)
	}
	end = len(s) - 1
	idx = []int{beg, end}
	return idx
}

func convert(content string, output string, start int) string {
	text, option, last, nword := getOption(content, start)

	if option == "" {
		result := *results + content[start:]
		*results = ""
		return result
	}

	idx := getIndexRange(text, nword)
	*results = correct(text, option, idx)

	return convert(content, output, last+1)
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
		text = text[:beg] + strings.Title(strings.ToLower(word)) + text[end+1:]
	case "bin":
		w, err := strconv.ParseInt(word, 2, 64)
		if err != nil {
			fmt.Printf("Cannot convert %s (not a binary number) to decimal \n", word)
			os.Exit(1)
		}
		s := strconv.FormatInt(w, 10)
		text = text[:beg] + s + text[end+1:]
	case "hex":
		w, err := strconv.ParseInt(word, 16, 64)
		if err != nil {
			fmt.Printf("Cannot convert %s (not an hexadecimal number) to decimal \n", word)
			os.Exit(1)
		}
		
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

		if isPunctuation(v) && last == ' ' && last != 0 {
			return false, i - 1, 0
		}

		if isPunctuation(v) && next != ' ' && !isPunctuation(rune(next)) && !isQuote(rune(next)) && next != 0 {
			return false, i, 1
		}
	}
	if isPunctuation(rune(content[len(content)-1])) && content[len(content)-2] == ' ' {
		return false, len(content) - 2, 0

	}
	if isPunctuation(rune(content[0])) && content[1] != ' ' && !isPunctuation(rune(content[1])) {
		return false, 1, 1

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
		if (s[i] < 'a' && s[i] > 'z') || (s[i] < 'A' && s[i] > 'Z') || (s[i] < '0' && s[i] > '9') || s[i] == ' ' || isQuote(rune(s[i])) {
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
	//content = fixQuotes(content)
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

func grammarCheck(content string) string {
	words := strings.Split(content, " ")
	chars := []string{"a", "e", "i", "o", "u", "a", "h"}

	for i, v := range words {
		valid := true
		v = strings.ToLower(v)

		if v == "a" && i < len(words)-1 {
			if len(words[i:]) > 0 {
				w := words[i : i+2]
				_, option, _, _ := getOption(strings.Join(w, " "), 0)
				if option != "" && len(words[i+1:]) > 1 {
					for _, c := range words[i+2] {
						if isAlpha(string(c)) {
							if isIn(chars, strings.ToLower(string(c))) {
								valid = false
								break
							} else {
								break
							}
						}
						continue
					}
				} else {
					for _, c := range words[i+1] {
						if isAlpha(string(c)) {
							if isIn(chars, strings.ToLower(string(c))) {
								valid = false
								break
							} else {
								break
							}
						}
						continue
					}
				}
				if !valid {
					words[i] += "n"

				}
			}
		}
	}
	*results = ""
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

func isQuote(r rune) bool {
	return r == '"' || r == '\''
}

func format(str string) string {
	for i := 0; i < len(str)-1; i++ {
		if i >= 0 && i < len(str)-1 {
			if IsBracket(rune(str[i])) {
				if i > 0 && str[i-1] != ' ' && !IsBracket(rune(str[i-1])) {
					str = addSpaceBetweenString(str, i-1)
					i++
				}
				if i < len(str)-1 && str[i+1] == ' ' {
					str = RemoveStringElem(str, i+1)
				}
			}
			if str[i] == ')' {
				if i > 0 && str[i-1] == ' ' {
					str = RemoveStringElem(str, i-1)
					i--
				}
				if i < len(str)-1 && str[i+1] != ' ' && str[i+1] != ')' {
					str = addSpaceBetweenString(str, i)
					i++
				}

			}
		}
	}
	return str
}

func process(s string) string {
	var out string

	if len(s) == 0 {
		fmt.Println("Error: Empty file")
		os.Exit(1)
	}

	output := RemSpace(s)
	output = format(output)
	// output = fixQuotes(output)
	output = puncCheck(output)
	output = grammarCheck(output)

	out += convert(output, out, 0)
	// out = puncCheck(out)
	out = quoteCheck(out)
	out = RemSpace(out)

	if len(out) > 0 && out[len(out)-1] == ' ' {
		out = RemoveStringElem(out, len(out)-1)
	}
	if len(out) > 0 && out[0] == ' ' {
		out = RemoveStringElem(out, 0)
	}

	*results = ""

	return out
}

func IsBracket(r rune) bool {
	brackets := []rune{'(', ')', '{', '}', '[', ']'}
	for _, v := range brackets {
		if r == v {
			return true
		}
	}
	return false
}

func RemSpace(s string) string {
	valid := true
	var idx int
	for i, v := range s {
		if i < len(s)-1 && v == ' ' && s[i+1] == ' ' {
			valid = false
			idx = i + 1
			break
		}
	}
	if valid {
		if (isPunctuation(rune(s[len(s)-1])) || isQuote(rune(s[len(s)-1])) || IsBracket(rune(s[len(s)-1]))) && rune(s[len(s)-2]) == ' ' {
			s = RemoveStringElem(s, len(s)-2)
		}
		return s
	}

	s = RemoveStringElem(s, idx)
	return RemSpace(s)
}

func quoteCheck(s string) string {
	var quoteType rune
	squote := '\''
	dquote := '"'
	var qa, qb = true, true
	if s[0] == byte(squote) || s[0] == byte(dquote) {
		if s[1] == ' ' {
			s = RemoveStringElem(s, 1)
		}
		qa = s[0] != byte(squote)
		qb = s[0] != byte(dquote)
	}
	for i := 0; i < len(s); i++ {
		if i > 0 && i < len(s)-1 {
			if isQuote(rune(s[i])) {
				if isAlpha(string(s[i-1])) && isAlpha(string(s[i+1])) {
					if isPunctuation(rune(s[i+1])) || isPunctuation(rune(s[i-1])) {
						qa = byte(quoteType) != byte(squote)
						qb = byte(quoteType) != byte(dquote)
					}
					continue
				}
				quoteType = rune(s[i])
				switch quoteType {
				case squote:
					if qa {
						if !isQuote(rune(s[i-1])) && s[i-1] != ' ' {
							s = addSpaceBetweenString(s, i-1)
							i++
						}
						if s[i+1] == ' ' {
							s = RemoveStringElem(s, i+1)

						}
						qa = false
					} else {
						if s[i-1] == ' ' {

							s = RemoveStringElem(s, i-1)
							i--
						}
						if s[i+1] != ' ' && !isQuote(rune(s[i+1])) {
							s = addSpaceBetweenString(s, i+1)

						}
						qa = true
					}
					quoteType = 0
				case dquote:
					if qb {
						if !isQuote(rune(s[i-1])) && s[i-1] != ' ' {
							s = addSpaceBetweenString(s, i-1)

						}
						if s[i+1] == ' ' {
							s = RemoveStringElem(s, i+1)

						}
						qb = false
					} else {
						if s[i-1] == ' ' {
							s = RemoveStringElem(s, i-1)
							i--
						}
						if s[i+1] != ' ' && !isQuote(rune(s[i+1])) {
							s = addSpaceBetweenString(s, i+1)

						}
						qb = true
					}
					quoteType = 0
				}
			}
		}
	}
	return s
}
