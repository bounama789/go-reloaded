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
			if a > 0 {
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

			} else {
				option = strings.ToLower(s[a+1 : b])
				if !isIn(options, option) {
					continue
				}

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
		return output + content[start:]
	}

	idx := getIndexRange(text, nword)
	output += correct(text, option, idx)

	return convert(content, output, last+1)
}

func correct(text string, option string, idx []int) string {
	beg := idx[0]
	end := idx[1]
	word := text[beg : end+1]

	switch option {
	case "low":
		// for i, v := range word {
		// 	if isAlpha(string(v)) {
		text = text[:beg] + strings.ToLower(word) + text[end+1:]
		// }
		// }
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
		if s[i] < 'a' && s[i] > 'z' || s[i] < 'A' && s[i] > 'Z' || s[i] < '0' && s[i] > '9' || s[i] == ' ' {
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
	inQuote := false
	var last,next rune
	for i := 0; i < len(str); i++ {
		if i > 0 && i < len(str)-1 {
			last = rune(str[i-1])
			next = rune(str[i+1])
		}
			if isQuote(rune(str[i])) {
				if i >0 && isQuote(rune(str[i])) && !inQuote &&  (!isAlpha(string(last)) || !isAlpha(string(next))) {
					inQuote = true
					if str[i-1] != ' ' {
						str = addSpaceBetweenString(str, i-1)
						i++
					}
					if str[i+1] == ' ' {
						str = RemoveStringElem(str, i+1)
					}
				} else if i < len(str) && (!isAlpha(string(last)) || !isAlpha(string(next))){
					inQuote = true
					if i>0 && str[i-1] == ' ' {
						str = RemoveStringElem(str, i-1)
						i--

					}
					if i<len(str)-1 && str[i+1] != ' ' {
						str = addSpaceBetweenString(str, i)
						i++
					}
				}
				if i == len(str)-1 && str[i-1] == ' '{
					str = RemoveStringElem(str, i-1)
				}
			}
		
	}
	return str
}

func isQuote(r rune) bool {
	return r == '"' || r == '\''
}

// (max(max(cap, 2)))

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
				if i < len(str)-1 && str[i+1] != ' ' && !IsBracket(rune(str[i+1])) {
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

	output := format(s)
	output = puncCheck(output)
	output = fixQuotes(output)

	out += convert(output, out, 0)
	out = puncCheck(out)
	out = fixQuotes(out)

	out = grammarCheck(out)

	if len(out) > 0 {
		if out[len(out)-1] == ' ' {
			out = RemoveStringElem(out, len(out)-1)
		}
		if out[0] == ' ' {
			out = RemoveStringElem(out, 0)
		}
	}

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
