package main

import (
	"os"
	"piscine"
	"strconv"
)

func main()  {

	const  (
		hex="0123456789ABCDEF"
		bin="01"
	)

	args := os.Args
	options := []string{"(low)","(up)","(bin)","(hex)","(cap)"}
	if len(args) > 1 {
		filename := args[1]
		content, _ := os.ReadFile(filename)
		lines := piscine.Split(string(content),"\n")
		var n string
		var words []string
		for _, v := range lines {
			words = append(words, piscine.SplitWhiteSpaces(v)...)

		}
		for i, v := range words {
			if v ==  options[3]{
				n = words[i-1]
				words[i-1] = strconv.Itoa(piscine.AtoiBase(n,hex))
				words = piscine.RemoveStringElem(words,i)
			}
			if v == options[2] {
				n = words[i-1]
				words[i-1] = strconv.Itoa(piscine.AtoiBase(n,bin))
				words = piscine.RemoveStringElem(words,i)
			}
			if v == options[1] {
				n = words[i-1]
				words[i-1] = piscine.ToUpper(n)
				words = piscine.RemoveStringElem(words,i)
			}
			if v== options[0] {
				n = words[i-1]
				words[i-1] = piscine.ToLower(n)
				words = piscine.RemoveStringElem(words,i)
			}
			if v == options[4] {
				n = words[i-1]
				words[i-1] = piscine.Capitalize(string(n[0]))+string(words[i-1][1:])
				words = piscine.RemoveStringElem(words,i)
			}

		}
		
		for _, v := range words {
			println(v)
		}

	}
		
}

func getPuncIndexes(content string) [][]int {
	options := []string{"(low)","(up)","(bin)","(hex)","(cap)"}
	punctuations := []rune{'.', ',', '!', '?', ':', ';'}
	runes := []rune(content)
	var dots,comma,excl,kst,col,scol,ln []int
	var n = 0

	for i := 0; i < len(runes); i++ {
		if string(runes[i:i+len(options[0])]) == options[0] || 
		string(runes[i:i+len(options[0])]) == options[2] ||
		string(runes[i:i+len(options[0])]) == options[3] || 
		string(runes[i:i+len(options[0])]) == options[4]{
			n += 5
		} else if string(runes[i:i+len(options[0])]) == options[1] {
			n+=4
		}

		if (runes[i] == punctuations[0] ){
			dots = append(dots, i-n)
		}
		if (runes[i] == punctuations[1] ){
			comma = append(comma, i-n)
		}
		if (runes[i] == punctuations[2] ){
			excl = append(excl, i-n)
		}
		if (runes[i] == punctuations[3] ){
			kst = append(kst, i-n)
		}
		if (runes[i] == punctuations[4] ){
			col = append(col, i-n)
		}
		if (runes[i] == punctuations[5] ){
			scol = append(scol, i-n)
		}
		if (runes[i] == punctuations[5] ){
			scol = append(scol, i-n)
		}
	}
}

