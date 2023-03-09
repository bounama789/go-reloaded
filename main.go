package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	args := os.Args

	if len(args) != 3 {
		filename := args[1]
		content, _ := os.ReadFile(filename)
		text := string(content)
		var output string

		output += process(text,output,0)
		println(output)
		
	}

}

func getOption(s string, begin int) (string,string,int) {
	var a,b int
	var text,option string
	for i := begin; i < len(s); i++ {
		if s[i] == '(' {
			a = i
			text = s[begin:a-1]
		}
		if s[i] == ')' {
			b=i
			option = s[a+1:b]
			break
		}
	}
	return text,option,b
}

// func convertHex(s string, idx int) string {
// 	base := "0123456789ABCDEF"

// 	var hex string
// 	for i := idx - 2; i > 0; i-- {
// 		if s[i] == ' ' {
// 			hex = s[i : idx-1]
// 			s = s[:i] + strconv.Itoa(piscine.AtoiBase(hex, base)) + s[:idx-1]
// 		}
// 	}
// 	return s
// }

func getIndexRange(s string, nword int) []int {
	words := strings.Split(s," ")
	var beg,end int
	var w string
	var idx []int

	if nword <= 1 {
		w = words[len(words)-1]
		beg = strings.LastIndex(s,w)
	} else {
		w = words[len(words)-(nword+1)]
		beg = strings.LastIndex(s,w)
	}
	end = len(s)-1
	idx = []int{beg,end}
	return idx
}

func process(content string, output string, start int) string {
	text,option,last := getOption(content,start)
	var nword = 0

	if option == "" {
		return output+ content[start:]
	}

	opt := strings.Split(option,", ")
	option = opt[0]

	if len(opt) > 1 {
		nword,_ = strconv.Atoi(opt[1])
	}

	idx := getIndexRange(text,nword)
	output += correct(text,option,idx)

	return process(content,output,last+1)
}

func correct(text string, option string, idx []int) string {
	beg := idx[0]
	end := idx[1]
	word := text[beg:end+1]

	switch option {
	case "low":
		text = text[:beg]+ strings.ToLower(word) + text[end+1:]
	case "up":
		text = text[:beg]+ strings.ToUpper(word) + text[end+1:]
	case "cap":
		text = text[:beg]+ strings.Title(word) + text[end+1:]
	case "bin":
		w,_:= strconv.ParseInt(word,2,8)
		s:= strconv.FormatInt(w,10)
		text = text[:beg]+ s + text[end+1:]
	case "hex":
		w,_:= strconv.ParseInt(word,16,8)
		s:= strconv.FormatInt(w,10)
		text = text[:beg]+ s + text[end+1:]
	}
	return text
}

func validPunct(content string)  {
	
}
