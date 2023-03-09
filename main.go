package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	if len(args) == 3 {
		filename := args[1]
		outfilename := args[2]
		content, _ := os.ReadFile(filename)
		text := string(content)
		var output string

		file, _ := os.Create(outfilename)

		output += process(text, output, 0)
		output = puncCheck(output)
		output = quoteCheck(output)
		output = grammarCheck(output)

		_, err := file.WriteString(output)
		if err != nil {
			panic(err)
		}
		println(output)

	}

}

func getOption(s string, begin int) (string, string, int) {
	var a, b int
	var text, option string
	for i := begin; i < len(s); i++ {
		if s[i] == '(' {
			a = i
			text = s[begin : a-1]
		}
		if s[i] == ')' {
			b = i
			option = s[a+1 : b]
			break
		}
	}
	return text, option, b
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
		w = words[len(words)-(nword+1)]
		beg = strings.LastIndex(s, w)
	}
	end = len(s) - 1
	idx = []int{beg, end}
	return idx
}

func process(content string, output string, start int) string {
	text, option, last := getOption(content, start)
	var nword = 0

	if option == "" {
		return output + content[start:]
	}

	opt := strings.Split(option, ", ")
	option = opt[0]

	if len(opt) > 1 {
		nword, _ = strconv.Atoi(opt[1])
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
		if i > 0 && i < len(content)-1 {
			next := content[i+1]
			last := content[i-1]

			if isPunctuation(v) && isAlpha(string(next)) && next != ' ' && !isPunctuation(rune(next)) {
				return false, i + 1, 1
			}
			if isPunctuation(v) && last == ' ' {
				return false, i - 1, 0
			}

		}
	}
	return true, -1, -1
}

func isPunctuation(val rune) bool {
	punctuation := []rune{'.', ',', '!', '?', ':', ';', '\'', '"'}

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
		if v == "a" && isIn(chars, string(words[i+1][0])) {
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
